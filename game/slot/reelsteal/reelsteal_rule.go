package reelsteal

// See: https://www.youtube.com/watch?v=5wEFg65Maa0

import (
	"github.com/slotopol/server/game/slot"
)

// Lined payment.
var LinePay = [12][5]float64{
	{},                    //  1 wild
	{},                    //  2 scatter
	{0, 0, 25, 150, 1500}, //  3 killer
	{0, 0, 20, 100, 1000}, //  4 baby
	{0, 0, 15, 75, 750},   //  5 boss
	{0, 0, 12, 60, 400},   //  6 driver
	{0, 0, 10, 50, 200},   //  7 thug
	{0, 0, 10, 20, 100},   //  8 safe
	{0, 0, 5, 15, 75},     //  9 case
	{0, 0, 4, 12, 60},     // 10 bag
	{0, 0, 2, 10, 50},     // 11 plan
	{0, 0, 2, 10, 40},     // 12 gun
}

// Scatters payment.
var ScatPay = [5]float64{0, 2, 4, 15, 100} // 2 scatter

// Scatter freespins table
var ScatFreespin = [5]int{0, 0, 15, 20, 25} // 2 scatter

// Bet lines
var BetLines = slot.BetLinesNvm9

type Game struct {
	slot.Slot5x3 `yaml:",inline"`
	// free spin number
	FS int `json:"fs,omitempty" yaml:"fs,omitempty" xml:"fs,omitempty"`
}

func NewGame() *Game {
	return &Game{
		Slot5x3: slot.Slot5x3{
			Sel: len(BetLines),
			Bet: 1,
		},
		FS: 0,
	}
}

const wild, scat = 1, 2

func (g *Game) Scanner(screen slot.Screen, wins *slot.Wins) {
	g.ScanLined(screen, wins)
	g.ScanScatters(screen, wins)
}

// Lined symbols calculation.
func (g *Game) ScanLined(screen slot.Screen, wins *slot.Wins) {
	for li := 1; li <= g.Sel; li++ {
		var line = BetLines[li-1]

		var mw float64 = 1 // mult wild
		var numl slot.Pos = 5
		var syml slot.Sym
		var x slot.Pos
		for x = 1; x <= 5; x++ {
			var sx = screen.Pos(x, line)
			if sx == wild {
				mw = 5
			} else if syml == 0 && sx != scat {
				syml = sx
			} else if sx != syml {
				numl = x - 1
				break
			}
		}

		if numl >= 3 && syml > 0 {
			var mm float64 = 1 // mult mode
			if g.FS > 0 {
				mm = 5
			}
			var pay = LinePay[syml-1][numl-1]
			*wins = append(*wins, slot.WinItem{
				Pay:  g.Bet * pay,
				Mult: mw * mm,
				Sym:  syml,
				Num:  numl,
				Line: li,
				XY:   line.CopyL(numl),
			})
		}
	}
}

// Scatters calculation.
func (g *Game) ScanScatters(screen slot.Screen, wins *slot.Wins) {
	var count = screen.ScatNum(scat)
	if g.FS > 0 {
		*wins = append(*wins, slot.WinItem{
			Pay:  0,
			Mult: 1,
			Sym:  scat,
			Num:  count,
			XY:   screen.ScatPos(scat),
			Free: int(count),
		})
	} else if count >= 2 {
		var pay, fs = ScatPay[count-1], ScatFreespin[count-1]
		*wins = append(*wins, slot.WinItem{
			Pay:  g.Bet * float64(g.Sel) * pay,
			Mult: 1,
			Sym:  scat,
			Num:  count,
			XY:   screen.ScatPos(scat),
			Free: fs,
		})
	}
}

func (g *Game) Spin(screen slot.Screen, mrtp float64) {
	var reels, _ = slot.FindReels(ReelsMap, mrtp)
	screen.Spin(reels)
}

func (g *Game) Apply(screen slot.Screen, wins slot.Wins) {
	if g.FS > 0 {
		g.Gain += wins.Gain()
	} else {
		g.Gain = wins.Gain()
	}

	if g.FS > 0 {
		g.FS--
	}
	for _, wi := range wins {
		if wi.Free > 0 {
			g.FS += wi.Free
		}
	}
}

func (g *Game) FreeSpins() int {
	return g.FS
}

func (g *Game) SetSel(sel int) error {
	return g.SetSelNum(sel, len(BetLines))
}
