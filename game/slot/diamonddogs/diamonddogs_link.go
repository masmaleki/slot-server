//go:build !prod || full || netent

package diamonddogs

import (
	"context"

	game "github.com/slotopol/server/game"
	"github.com/spf13/pflag"
)

func init() {
	var gi = game.GameInfo{
		Aliases: []game.GameAlias{
			{ID: "diamonddogs", Name: "Diamond Dogs"},
			{ID: "voodoovibes", Name: "Voodoo Vibes"},
		},
		Provider: "NetEnt",
		SX:       5,
		SY:       3,
		GP: game.GPsel |
			game.GPretrig |
			game.GPfgmult |
			game.GPfgreel |
			game.GPscat |
			game.GPwild,
		SN:  len(LinePay),
		LN:  len(BetLines),
		BN:  1,
		RTP: game.MakeRtpList(ReelsMap),
	}
	game.GameList = append(game.GameList, gi)

	for _, ga := range gi.Aliases {
		game.ScanIters = append(game.ScanIters, func(flags *pflag.FlagSet, ctx context.Context) {
			if is, _ := flags.GetBool(ga.ID); is {
				var mrtp, _ = flags.GetFloat64("reels")
				CalcStatReg(ctx, mrtp)
			}
		})
		game.GameFactory[ga.ID] = func() any {
			return NewGame()
		}
	}
}
