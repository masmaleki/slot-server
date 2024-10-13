//go:build !prod || full || playtech

package goldentour

import (
	"context"

	game "github.com/slotopol/server/game"
	"github.com/spf13/pflag"
)

func init() {
	var gi = game.GameInfo{
		Aliases: []game.GameAlias{
			{ID: "goldentour", Name: "Golden Tour"},
		},
		Provider: "Playtech",
		SX:       5,
		SY:       3,
		GP: game.GPsel |
			game.GPfgno |
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
				CalcStat(ctx, mrtp)
			}
		})
		game.GameFactory[ga.ID] = func() any {
			return NewGame()
		}
	}
}
