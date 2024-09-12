package spi

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"math/rand/v2"
	"sync/atomic"

	cfg "github.com/slotopol/server/config"
	"github.com/slotopol/server/config/links"
	"github.com/slotopol/server/util"

	"github.com/gin-gonic/gin"
	keno "github.com/slotopol/server/game/keno"
	slot "github.com/slotopol/server/game/slot"
)

var (
	SpinBuf util.SqlBuf[Spinlog]
	MultBuf util.SqlBuf[Multlog]
	BankBat = map[uint64]*SqlBank{}
	JoinBuf = SqlStory{}
)

// Joins to game and creates new instance of game.
func SpiGameJoin(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"arg"`
		CID     uint64   `json:"cid" yaml:"cid" xml:"cid,attr" form:"cid"`
		UID     uint64   `json:"uid" yaml:"uid" xml:"uid,attr" form:"uid"`
		Alias   string   `json:"alias" yaml:"alias" xml:"alias" form:"alias" binding:"required"`
	}
	var ret struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"ret"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr"`
		Game    any      `json:"game" yaml:"game" xml:"game"`
		Scrn    any      `json:"scrn" yaml:"scrn" xml:"scrn"`
		Wallet  float64  `json:"wallet" yaml:"wallet" xml:"wallet"`
	}

	if err = c.ShouldBind(&arg); err != nil {
		Ret400(c, SEC_game_join_nobind, err)
		return
	}
	if arg.CID == 0 {
		Ret400(c, SEC_game_join_norid, ErrNoCID)
		return
	}
	if arg.UID == 0 {
		Ret400(c, SEC_game_join_nouid, ErrNoUID)
		return
	}

	var club *Club
	if club, ok = Clubs.Get(arg.CID); !ok {
		Ret404(c, SEC_game_join_noclub, ErrNoClub)
		return
	}

	var user *User
	if user, ok = Users.Get(arg.UID); !ok {
		Ret404(c, SEC_game_join_nouser, ErrNoUser)
		return
	}

	var admin, al = MustAdmin(c, arg.CID)
	if (al&ALmem == 0) || (admin != user && al&ALgame == 0) {
		Ret403(c, SEC_game_join_noaccess, ErrNoAccess)
		return
	}

	var alias = util.ToID(arg.Alias)
	var maker, has = links.GameFactory[alias]
	if !has {
		Ret400(c, SEC_game_join_noalias, ErrNoAliase)
		return
	}

	var anygame = maker()
	var gid = atomic.AddUint64(&StoryCounter, 1)
	var scene = &Scene{
		Story: Story{
			GID:   gid,
			Alias: alias,
			CID:   arg.CID,
			UID:   arg.UID,
			Flow:  true,
		},
		Game: anygame,
	}

	// insert new story entry
	if Cfg.ClubInsertBuffer > 1 {
		go JoinBuf.Join(cfg.XormStorage, &scene.Story)
	} else if err = JoinBuf.Join(cfg.XormStorage, &scene.Story); err != nil {
		Ret500(c, SEC_game_join_sql, err)
		return
	}

	Scenes.Set(scene.GID, scene)
	user.games.Set(scene.GID, scene)

	type kwins = keno.Wins
	type kscrn = keno.KenoScreen

	ret.GID = gid
	ret.Game = anygame
	var scrn slot.Screen
	switch game := anygame.(type) {
	case slot.SlotGame:
		scrn = game.NewScreen()

		// make game screen object
		club.mux.RLock()
		var rtp = GetRTP(user, club)
		club.mux.RUnlock()
		game.Spin(scrn, rtp)
		ret.Scrn = scrn

	case keno.KenoGame:
		var ks kscrn
		var kw kwins
		game.Spin(&ks, kw.Hits[:])
		ret.Scrn = &ks
	}
	ret.Wallet = user.GetWallet(arg.CID)

	RetOk(c, ret)
}

// Removes instance of opened game.
func SpiGamePart(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"arg"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr" form:"gid"`
	}

	if err = c.ShouldBind(&arg); err != nil {
		Ret400(c, SEC_game_part_nobind, err)
		return
	}
	if arg.GID == 0 {
		Ret400(c, SEC_game_part_nogid, ErrNoGID)
		return
	}

	var scene *Scene
	if scene, ok = Scenes.Get(arg.GID); !ok {
		Ret404(c, SEC_game_part_notopened, ErrNotOpened)
		return
	}

	var user *User
	if user, ok = Users.Get(scene.UID); !ok {
		Ret500(c, SEC_game_part_nouser, ErrNoUser)
		return
	}

	var admin, al = MustAdmin(c, scene.CID)
	if admin != user && al&ALgame == 0 {
		Ret403(c, SEC_game_part_noaccess, ErrNoAccess)
		return
	}

	// update story entry
	if Cfg.ClubUpdateBuffer > 1 {
		go JoinBuf.Flow(cfg.XormStorage, arg.GID, false)
	} else if err = JoinBuf.Flow(cfg.XormStorage, arg.GID, false); err != nil {
		Ret500(c, SEC_game_part_sql, err)
		return
	}

	scene.Flow = false
	Scenes.Delete(arg.GID)
	user.games.Delete(arg.GID)

	Ret204(c)
}

// Returns full info of game scene with given GID, and balance on wallet.
func SpiGameInfo(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"arg"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr" form:"gid"`
	}
	var ret struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"ret"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr"`
		Alias   string   `json:"alias" yaml:"alias" xml:"alias"`
		CID     uint64   `json:"cid" yaml:"cid" xml:"cid,attr"`
		UID     uint64   `json:"uid" yaml:"uid" xml:"uid,attr"`
		SID     uint64   `json:"sid" yaml:"sid" xml:"sid,attr"`
		Game    any      `json:"game" yaml:"game" xml:"game"`
		Wallet  float64  `json:"wallet" yaml:"wallet" xml:"wallet"`
	}

	if err = c.ShouldBind(&arg); err != nil {
		Ret400(c, SEC_game_info_nobind, err)
		return
	}
	if arg.GID == 0 {
		Ret400(c, SEC_game_info_nogid, ErrNoGID)
		return
	}

	var scene *Scene
	if scene, ok = Scenes.Get(arg.GID); !ok {
		Ret404(c, SEC_game_info_notopened, ErrNotOpened)
		return
	}

	var user *User
	if user, ok = Users.Get(scene.UID); !ok {
		Ret500(c, SEC_game_info_nouser, ErrNoUser)
		return
	}

	var admin, al = MustAdmin(c, scene.CID)
	if admin != user && al&ALgame == 0 {
		Ret403(c, SEC_game_info_noaccess, ErrNoAccess)
		return
	}

	var props *Props
	if props, ok = user.props.Get(scene.CID); !ok {
		Ret500(c, SEC_game_info_noprops, ErrNoProps)
		return
	}

	ret.GID = arg.GID
	ret.Alias = scene.Alias
	ret.CID = scene.CID
	ret.UID = scene.UID
	ret.SID = scene.SID
	ret.Game = scene.Game
	ret.Wallet = props.Wallet

	RetOk(c, ret)
}

// Returns bet value.
func SpiSlotBetGet(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"arg"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr" form:"gid"`
	}
	var ret struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"ret"`
		Bet     float64  `json:"bet" yaml:"bet" xml:"bet"`
	}

	if err = c.ShouldBind(&arg); err != nil {
		Ret400(c, SEC_slot_betget_nobind, err)
		return
	}
	if arg.GID == 0 {
		Ret400(c, SEC_slot_betget_nogid, ErrNoGID)
		return
	}

	var scene *Scene
	if scene, ok = Scenes.Get(arg.GID); !ok {
		Ret404(c, SEC_slot_betget_notopened, ErrNotOpened)
		return
	}
	var game slot.SlotGame
	if game, ok = scene.Game.(slot.SlotGame); !ok {
		Ret403(c, SEC_slot_betget_notslot, ErrNotSlot)
		return
	}

	var admin, al = MustAdmin(c, scene.CID)
	if admin.UID != scene.UID && al&ALgame == 0 {
		Ret403(c, SEC_slot_betget_noaccess, ErrNoAccess)
		return
	}

	ret.Bet = game.GetBet()

	RetOk(c, ret)
}

// Set bet value.
func SpiSlotBetSet(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"arg"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr"`
		Bet     float64  `json:"bet" yaml:"bet" xml:"bet" binding:"required"`
	}

	if err = c.ShouldBind(&arg); err != nil {
		Ret400(c, SEC_slot_betset_nobind, err)
		return
	}
	if arg.GID == 0 {
		Ret400(c, SEC_slot_betset_nogid, ErrNoGID)
		return
	}

	var scene *Scene
	if scene, ok = Scenes.Get(arg.GID); !ok {
		Ret404(c, SEC_slot_betset_notopened, ErrNotOpened)
		return
	}
	var game slot.SlotGame
	if game, ok = scene.Game.(slot.SlotGame); !ok {
		Ret403(c, SEC_slot_betset_notslot, ErrNotSlot)
		return
	}

	var admin, al = MustAdmin(c, scene.CID)
	if admin.UID != scene.UID && al&ALgame == 0 {
		Ret403(c, SEC_slot_betset_noaccess, ErrNoAccess)
		return
	}

	if err = game.SetBet(arg.Bet); err != nil {
		Ret403(c, SEC_slot_betset_badbet, err)
		return
	}

	Ret204(c)
}

// Returns selected bet lines bitset.
func SpiSlotSblGet(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"arg"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr" form:"gid"`
	}
	var ret struct {
		XMLName xml.Name    `json:"-" yaml:"-" xml:"ret"`
		SBL     slot.Bitset `json:"sbl" yaml:"sbl" xml:"sbl"`
	}

	if err = c.ShouldBind(&arg); err != nil {
		Ret400(c, SEC_slot_sblget_nobind, err)
		return
	}
	if arg.GID == 0 {
		Ret400(c, SEC_slot_sblget_nogid, ErrNoGID)
		return
	}

	var scene *Scene
	if scene, ok = Scenes.Get(arg.GID); !ok {
		Ret404(c, SEC_slot_sblget_notopened, ErrNotOpened)
		return
	}
	var game slot.SlotGame
	if game, ok = scene.Game.(slot.SlotGame); !ok {
		Ret403(c, SEC_slot_sblget_notslot, ErrNotSlot)
		return
	}

	var admin, al = MustAdmin(c, scene.CID)
	if admin.UID != scene.UID && al&ALgame == 0 {
		Ret403(c, SEC_slot_sblget_noaccess, ErrNoAccess)
		return
	}

	ret.SBL = game.GetLines()

	RetOk(c, ret)
}

// Set selected bet lines bitset.
func SpiSlotSblSet(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name    `json:"-" yaml:"-" xml:"arg"`
		GID     uint64      `json:"gid" yaml:"gid" xml:"gid,attr"`
		SBL     slot.Bitset `json:"sbl" yaml:"sbl" xml:"sbl" binding:"required"`
	}

	if err = c.ShouldBind(&arg); err != nil {
		Ret400(c, SEC_slot_sblset_nobind, err)
		return
	}
	if arg.GID == 0 {
		Ret400(c, SEC_slot_sblset_nogid, ErrNoGID)
		return
	}

	var scene *Scene
	if scene, ok = Scenes.Get(arg.GID); !ok {
		Ret404(c, SEC_slot_sblset_notopened, ErrNotOpened)
		return
	}
	var game slot.SlotGame
	if game, ok = scene.Game.(slot.SlotGame); !ok {
		Ret403(c, SEC_slot_sblset_notslot, ErrNotSlot)
		return
	}

	var admin, al = MustAdmin(c, scene.CID)
	if admin.UID != scene.UID && al&ALgame == 0 {
		Ret403(c, SEC_slot_sblset_noaccess, ErrNoAccess)
		return
	}

	if err = game.SetLines(arg.SBL); err != nil {
		Ret403(c, SEC_slot_sblset_badlines, err)
		return
	}

	Ret204(c)
}

// Returns master RTP for given GID.
func SpiSlotRtpGet(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"arg"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr" form:"gid"`
	}
	var ret struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"ret"`
		RTP     float64  `json:"rtp" yaml:"rtp" xml:"rtp"`
	}

	if err = c.ShouldBind(&arg); err != nil {
		Ret400(c, SEC_slot_rdget_nobind, err)
		return
	}
	if arg.GID == 0 {
		Ret400(c, SEC_slot_rdget_nogid, ErrNoGID)
		return
	}

	var scene *Scene
	if scene, ok = Scenes.Get(arg.GID); !ok {
		Ret404(c, SEC_slot_rdget_notopened, ErrNotOpened)
		return
	}

	var club *Club
	if club, ok = Clubs.Get(scene.CID); !ok {
		Ret500(c, SEC_slot_rdget_noclub, ErrNoClub)
		return
	}

	var user *User
	if user, ok = Users.Get(scene.UID); !ok {
		Ret500(c, SEC_slot_rdget_nouser, ErrNoUser)
		return
	}

	var admin, al = MustAdmin(c, scene.CID)
	if admin.UID != scene.UID && al&ALgame == 0 {
		Ret403(c, SEC_slot_rdget_noaccess, ErrNoAccess)
		return
	}

	club.mux.RLock()
	ret.RTP = GetRTP(user, club)
	club.mux.RUnlock()

	RetOk(c, ret)
}

// Make a spin.
func SpiSlotSpin(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"arg"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr" form:"gid"`
	}
	var ret struct {
		XMLName xml.Name      `json:"-" yaml:"-" xml:"ret"`
		SID     uint64        `json:"sid" yaml:"sid" xml:"sid,attr"`
		Game    slot.SlotGame `json:"game" yaml:"game" xml:"game"`
		Scrn    slot.Screen   `json:"scrn" yaml:"scrn" xml:"scrn"`
		Wins    slot.Wins     `json:"wins,omitempty" yaml:"wins,omitempty" xml:"wins,omitempty"`
		Wallet  float64       `json:"wallet" yaml:"wallet" xml:"wallet"`
	}

	if err = c.ShouldBind(&arg); err != nil {
		Ret400(c, SEC_slot_spin_nobind, err)
		return
	}
	if arg.GID == 0 {
		Ret400(c, SEC_slot_spin_nogid, ErrNoGID)
		return
	}

	var scene *Scene
	if scene, ok = Scenes.Get(arg.GID); !ok {
		Ret404(c, SEC_slot_spin_notopened, ErrNotOpened)
		return
	}
	var game slot.SlotGame
	if game, ok = scene.Game.(slot.SlotGame); !ok {
		Ret403(c, SEC_slot_spin_notslot, ErrNotSlot)
		return
	}

	var club *Club
	if club, ok = Clubs.Get(scene.CID); !ok {
		Ret500(c, SEC_slot_spin_noclub, ErrNoClub)
		return
	}

	var user *User
	if user, ok = Users.Get(scene.UID); !ok {
		Ret500(c, SEC_slot_spin_nouser, ErrNoUser)
		return
	}

	var admin, al = MustAdmin(c, scene.CID)
	if admin.UID != scene.UID && al&ALgame == 0 {
		Ret403(c, SEC_slot_spin_noaccess, ErrNoAccess)
		return
	}

	var (
		fs       = game.FreeSpins()
		bet      = game.GetBet()
		sbl      = game.GetLines()
		totalbet float64
		banksum  float64
	)
	if fs == 0 {
		totalbet = bet * float64(sbl.Num())
	}

	var props *Props
	if props, ok = user.props.Get(scene.CID); !ok {
		Ret500(c, SEC_slot_spin_noprops, ErrNoProps)
		return
	}
	if props.Wallet < totalbet {
		Ret403(c, SEC_slot_spin_nomoney, ErrNoMoney)
		return
	}

	club.mux.RLock()
	var bank = club.Bank
	var mrtp = GetRTP(user, club)
	club.mux.RUnlock()

	// spin until gain less than bank value
	var wins slot.Wins
	defer wins.Reset()
	var scrn = game.NewScreen()
	defer scrn.Free()
	var n = 0
	for {
		game.Spin(scrn, mrtp)
		game.Scanner(scrn, &wins)
		game.Spawn(scrn, wins)
		banksum = totalbet - wins.Gain()
		if bank+banksum >= 0 || (bank < 0 && banksum > 0) {
			break
		}
		wins.Reset()
		if n >= cfg.Cfg.MaxSpinAttempts {
			Ret500(c, SEC_slot_spin_badbank, ErrBadBank)
			return
		}
		n++
	}

	// write gain and total bet as transaction
	if Cfg.ClubUpdateBuffer > 1 {
		go BankBat[scene.CID].Put(cfg.XormStorage, scene.UID, banksum)
	} else if err = BankBat[scene.CID].Put(cfg.XormStorage, scene.UID, banksum); err != nil {
		Ret500(c, SEC_slot_spin_sqlbank, err)
		return
	}

	// make changes to memory data
	club.mux.Lock()
	club.Bank += banksum
	club.mux.Unlock()

	props.Wallet -= banksum
	game.Apply(scrn, wins)

	// write spin result to log and get spin ID
	var sid = atomic.AddUint64(&SpinCounter, 1)
	scene.SID = sid
	var rec = Spinlog{
		SID:    sid,
		GID:    arg.GID,
		MRTP:   mrtp,
		Gain:   game.GetGain(),
		Wallet: props.Wallet,
	}
	var b []byte

	if b, err = json.Marshal(scene.Game); err != nil {
		return
	}
	rec.Game = util.B2S(b)

	if b, err = json.Marshal(scrn); err != nil {
		return
	}
	rec.Screen = util.B2S(b)

	if len(wins) > 0 {
		if b, err = json.Marshal(wins); err != nil {
			return
		}
		rec.Wins = util.B2S(b)
	}

	go func() {
		if err = SpinBuf.Put(cfg.XormSpinlog, rec); err != nil {
			log.Printf("can not write to spin log: %s", err.Error())
		}
	}()

	// prepare result
	ret.SID = sid
	ret.Game = game
	ret.Scrn = scrn
	ret.Wins = wins
	ret.Wallet = props.Wallet

	RetOk(c, ret)
}

// Double up gamble on last gain.
func SpiSlotDoubleup(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"arg"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr" form:"gid"`
		Mult    int      `json:"mult" yaml:"mult" xml:"mult" form:"mult"`
	}
	var ret struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"ret"`
		ID      uint64   `json:"id" yaml:"id" xml:"id,attr"`
		Gain    float64  `json:"gain" yaml:"gain" xml:"gain"`
		Wallet  float64  `json:"wallet" yaml:"wallet" xml:"wallet"`
	}

	if err = c.ShouldBind(&arg); err != nil {
		Ret400(c, SEC_slot_doubleup_nobind, err)
		return
	}
	if arg.GID == 0 {
		Ret400(c, SEC_slot_doubleup_nogid, ErrNoGID)
		return
	}
	if arg.Mult < 2 {
		Ret400(c, SEC_slot_doubleup_nomult, ErrNoMult)
		return
	}
	if arg.Mult > 10 {
		Ret400(c, SEC_slot_doubleup_bigmult, ErrBigMult)
		return
	}

	var scene *Scene
	if scene, ok = Scenes.Get(arg.GID); !ok {
		Ret404(c, SEC_slot_doubleup_notopened, ErrNotOpened)
		return
	}
	var game slot.SlotGame
	if game, ok = scene.Game.(slot.SlotGame); !ok {
		Ret403(c, SEC_slot_doubleup_notslot, ErrNotSlot)
		return
	}

	var club *Club
	if club, ok = Clubs.Get(scene.CID); !ok {
		Ret500(c, SEC_slot_doubleup_noclub, ErrNoClub)
		return
	}

	var user *User
	if user, ok = Users.Get(scene.UID); !ok {
		Ret500(c, SEC_slot_doubleup_nouser, ErrNoUser)
		return
	}

	var admin, al = MustAdmin(c, scene.CID)
	if admin.UID != scene.UID && al&ALgame == 0 {
		Ret403(c, SEC_slot_doubleup_noaccess, ErrNoAccess)
		return
	}

	var props *Props
	if props, ok = user.props.Get(scene.CID); !ok {
		Ret500(c, SEC_slot_doubleup_noprops, ErrNoProps)
		return
	}

	var risk = game.GetGain()
	if risk == 0 {
		Ret403(c, SEC_slot_doubleup_nomoney, ErrNoMoney)
		return
	}

	club.mux.RLock()
	var bank = club.Bank
	var mrtp = GetRTP(user, club)
	club.mux.RUnlock()

	var multgain float64 // new multiplied gain
	if bank >= risk*float64(arg.Mult) {
		var r = rand.Float64()
		var side = 1 / float64(arg.Mult) * mrtp / 100
		if r < side {
			multgain = risk * float64(arg.Mult)
		}
	}
	var banksum = risk - multgain

	// write gain and total bet as transaction
	if Cfg.ClubUpdateBuffer > 1 {
		go BankBat[scene.CID].Put(cfg.XormStorage, scene.UID, banksum)
	} else if err = BankBat[scene.CID].Put(cfg.XormStorage, scene.UID, banksum); err != nil {
		Ret500(c, SEC_slot_doubleup_sqlbank, err)
		return
	}

	// make changes to memory data
	club.mux.Lock()
	club.Bank += banksum
	club.mux.Unlock()

	props.Wallet -= banksum

	game.SetGain(multgain)

	// write doubleup result to log and get spin ID
	var id = atomic.AddUint64(&MultCounter, 1)
	go func() {
		var rec = Multlog{
			ID:     id,
			GID:    arg.GID,
			MRTP:   mrtp,
			Mult:   arg.Mult,
			Risk:   risk,
			Gain:   multgain,
			Wallet: props.Wallet,
		}
		if err = MultBuf.Put(cfg.XormSpinlog, rec); err != nil {
			log.Printf("can not write to mult log: %s", err.Error())
		}
	}()

	// prepare result
	ret.ID = id
	ret.Gain = multgain
	ret.Wallet = props.Wallet

	RetOk(c, ret)
}

func SpiSlotCollect(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"arg"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr" form:"gid"`
	}

	if err = c.ShouldBind(&arg); err != nil {
		Ret400(c, SEC_slot_collect_nobind, err)
		return
	}
	if arg.GID == 0 {
		Ret400(c, SEC_slot_collect_nogid, ErrNoGID)
		return
	}

	var scene *Scene
	if scene, ok = Scenes.Get(arg.GID); !ok {
		Ret404(c, SEC_slot_collect_notopened, ErrNotOpened)
		return
	}
	var game slot.SlotGame
	if game, ok = scene.Game.(slot.SlotGame); !ok {
		Ret403(c, SEC_slot_collect_notslot, ErrNotSlot)
		return
	}

	var admin, al = MustAdmin(c, scene.CID)
	if admin.UID != scene.UID && al&ALgame == 0 {
		Ret403(c, SEC_slot_collect_noaccess, ErrNoAccess)
		return
	}

	if err = game.SetGain(0); err != nil {
		Ret403(c, SEC_slot_collect_denied, err)
		return
	}

	RetOk(c, nil)
}
