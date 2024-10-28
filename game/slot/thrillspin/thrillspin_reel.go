package thrillspin

import (
	"github.com/slotopol/server/game/slot"
)

// reels lengths [35, 35, 35, 35, 35], total reshuffles 52521875
// symbols: 52.799(lined) + 14.71(scatter) = 67.509138%
// free spins 4345245, q = 0.082732, sq = 1/(1-q) = 1.090194
// free games frequency: 1/181.31
// RTP = 67.509(sym) + 0.082732*220.79(fg) = 85.775904%
var Reels857 = slot.Reels5x{
	{5, 4, 6, 11, 1, 12, 11, 8, 7, 4, 8, 10, 13, 3, 12, 13, 10, 9, 5, 8, 13, 9, 10, 13, 12, 6, 7, 9, 3, 5, 6, 11, 7, 2, 12},
	{9, 11, 6, 7, 2, 12, 8, 13, 9, 11, 6, 13, 10, 7, 13, 12, 5, 8, 6, 9, 4, 11, 5, 10, 12, 3, 5, 13, 10, 8, 1, 3, 12, 4, 7},
	{13, 8, 9, 7, 12, 3, 11, 10, 6, 4, 7, 13, 12, 4, 7, 9, 1, 5, 11, 6, 13, 8, 5, 3, 2, 8, 11, 9, 12, 6, 10, 5, 13, 12, 10},
	{12, 13, 2, 12, 5, 8, 13, 7, 10, 12, 11, 13, 3, 9, 12, 7, 6, 4, 3, 11, 6, 13, 10, 7, 11, 1, 8, 5, 9, 10, 8, 9, 5, 4, 6},
	{7, 4, 9, 2, 11, 12, 3, 10, 4, 13, 8, 5, 11, 7, 9, 8, 12, 9, 3, 6, 13, 12, 7, 6, 13, 12, 8, 5, 1, 10, 11, 5, 6, 10, 13},
}

// reels lengths [35, 35, 35, 35, 35], total reshuffles 52521875
// symbols: 53.836(lined) + 14.71(scatter) = 68.546610%
// free spins 4345245, q = 0.082732, sq = 1/(1-q) = 1.090194
// free games frequency: 1/181.31
// RTP = 68.547(sym) + 0.082732*224.19(fg) = 87.094098%
var Reels870 = slot.Reels5x{
	{13, 11, 6, 8, 7, 4, 10, 12, 9, 13, 7, 8, 10, 1, 5, 6, 3, 13, 8, 7, 10, 11, 9, 4, 12, 9, 5, 13, 3, 2, 12, 4, 11, 6, 5},
	{9, 11, 6, 7, 2, 12, 8, 13, 9, 11, 6, 13, 10, 7, 13, 12, 5, 8, 6, 9, 4, 11, 5, 10, 12, 3, 5, 13, 10, 8, 1, 3, 12, 4, 7},
	{13, 8, 9, 7, 12, 3, 11, 10, 6, 4, 7, 13, 12, 4, 7, 9, 1, 5, 11, 6, 13, 8, 5, 3, 2, 8, 11, 9, 12, 6, 10, 5, 13, 12, 10},
	{12, 13, 2, 12, 5, 8, 13, 7, 10, 12, 11, 13, 3, 9, 12, 7, 6, 4, 3, 11, 6, 13, 10, 7, 11, 1, 8, 5, 9, 10, 8, 9, 5, 4, 6},
	{7, 4, 9, 2, 11, 12, 3, 10, 4, 13, 8, 5, 11, 7, 9, 8, 12, 9, 3, 6, 13, 12, 7, 6, 13, 12, 8, 5, 1, 10, 11, 5, 6, 10, 13},
}

// reels lengths [35, 35, 35, 35, 35], total reshuffles 52521875
// symbols: 54.71(lined) + 14.71(scatter) = 69.420532%
// free spins 4345245, q = 0.082732, sq = 1/(1-q) = 1.090194
// free games frequency: 1/181.31
// RTP = 69.421(sym) + 0.082732*227.05(fg) = 88.204487%
var Reels882 = slot.Reels5x{
	{5, 4, 6, 11, 1, 12, 11, 8, 7, 4, 8, 10, 13, 3, 12, 13, 10, 9, 5, 8, 13, 9, 10, 13, 12, 6, 7, 9, 3, 5, 6, 11, 7, 2, 12},
	{7, 10, 8, 4, 13, 10, 12, 11, 9, 12, 13, 8, 10, 7, 13, 9, 5, 6, 9, 13, 11, 6, 5, 11, 4, 8, 6, 7, 1, 4, 3, 5, 12, 3, 2},
	{5, 3, 6, 11, 8, 9, 4, 7, 9, 2, 13, 10, 6, 1, 11, 9, 7, 8, 13, 7, 12, 8, 5, 13, 3, 4, 5, 10, 11, 13, 12, 10, 4, 6, 12},
	{12, 13, 2, 12, 5, 8, 13, 7, 10, 12, 11, 13, 3, 9, 12, 7, 6, 4, 3, 11, 6, 13, 10, 7, 11, 1, 8, 5, 9, 10, 8, 9, 5, 4, 6},
	{7, 4, 9, 2, 11, 12, 3, 10, 4, 13, 8, 5, 11, 7, 9, 8, 12, 9, 3, 6, 13, 12, 7, 6, 13, 12, 8, 5, 1, 10, 11, 5, 6, 10, 13},
}

// reels lengths [35, 35, 35, 35, 35], total reshuffles 52521875
// symbols: 55.531(lined) + 14.71(scatter) = 70.241523%
// free spins 4345245, q = 0.082732, sq = 1/(1-q) = 1.090194
// free games frequency: 1/181.31
// RTP = 70.242(sym) + 0.082732*229.73(fg) = 89.247624%
var Reels892 = slot.Reels5x{
	{13, 11, 6, 8, 7, 4, 10, 12, 9, 13, 7, 8, 10, 1, 5, 6, 3, 13, 8, 7, 10, 11, 9, 4, 12, 9, 5, 13, 3, 2, 12, 4, 11, 6, 5},
	{7, 10, 8, 4, 13, 10, 12, 11, 9, 12, 13, 8, 10, 7, 13, 9, 5, 6, 9, 13, 11, 6, 5, 11, 4, 8, 6, 7, 1, 4, 3, 5, 12, 3, 2},
	{13, 8, 9, 7, 12, 3, 11, 10, 6, 4, 7, 13, 12, 4, 7, 9, 1, 5, 11, 6, 13, 8, 5, 3, 2, 8, 11, 9, 12, 6, 10, 5, 13, 12, 10},
	{12, 13, 2, 12, 5, 8, 13, 7, 10, 12, 11, 13, 3, 9, 12, 7, 6, 4, 3, 11, 6, 13, 10, 7, 11, 1, 8, 5, 9, 10, 8, 9, 5, 4, 6},
	{7, 4, 9, 2, 11, 12, 3, 10, 4, 13, 8, 5, 11, 7, 9, 8, 12, 9, 3, 6, 13, 12, 7, 6, 13, 12, 8, 5, 1, 10, 11, 5, 6, 10, 13},
}

// reels lengths [35, 35, 35, 35, 35], total reshuffles 52521875
// symbols: 56.3(lined) + 14.71(scatter) = 71.010631%
// free spins 4345245, q = 0.082732, sq = 1/(1-q) = 1.090194
// free games frequency: 1/181.31
// RTP = 71.011(sym) + 0.082732*232.25(fg) = 90.224839%
var Reels902 = slot.Reels5x{
	{13, 11, 6, 8, 7, 4, 10, 12, 9, 13, 7, 8, 10, 1, 5, 6, 3, 13, 8, 7, 10, 11, 9, 4, 12, 9, 5, 13, 3, 2, 12, 4, 11, 6, 5},
	{7, 10, 8, 4, 13, 10, 12, 11, 9, 12, 13, 8, 10, 7, 13, 9, 5, 6, 9, 13, 11, 6, 5, 11, 4, 8, 6, 7, 1, 4, 3, 5, 12, 3, 2},
	{13, 8, 9, 7, 12, 3, 11, 10, 6, 4, 7, 13, 12, 4, 7, 9, 1, 5, 11, 6, 13, 8, 5, 3, 2, 8, 11, 9, 12, 6, 10, 5, 13, 12, 10},
	{13, 9, 8, 11, 12, 6, 9, 4, 5, 3, 12, 13, 7, 6, 10, 5, 7, 8, 13, 7, 11, 5, 10, 13, 9, 6, 4, 1, 10, 4, 8, 12, 2, 11, 3},
	{13, 6, 12, 4, 10, 8, 2, 7, 13, 11, 4, 9, 6, 12, 11, 8, 13, 5, 6, 4, 7, 1, 10, 9, 3, 7, 10, 3, 8, 12, 5, 13, 9, 11, 5},
}

// reels lengths [35, 35, 35, 35, 35], total reshuffles 52521875
// symbols: 56.963(lined) + 14.71(scatter) = 71.672927%
// free spins 4345245, q = 0.082732, sq = 1/(1-q) = 1.090194
// free games frequency: 1/181.31
// RTP = 71.673(sym) + 0.082732*234.41(fg) = 91.066340%
var Reels910 = slot.Reels5x{
	{13, 11, 6, 8, 7, 4, 10, 12, 9, 13, 7, 8, 10, 1, 5, 6, 3, 13, 8, 7, 10, 11, 9, 4, 12, 9, 5, 13, 3, 2, 12, 4, 11, 6, 5},
	{7, 10, 8, 4, 13, 10, 12, 11, 9, 12, 13, 8, 10, 7, 13, 9, 5, 6, 9, 13, 11, 6, 5, 11, 4, 8, 6, 7, 1, 4, 3, 5, 12, 3, 2},
	{5, 3, 6, 11, 8, 9, 4, 7, 9, 2, 13, 10, 6, 1, 11, 9, 7, 8, 13, 7, 12, 8, 5, 13, 3, 4, 5, 10, 11, 13, 12, 10, 4, 6, 12},
	{12, 13, 2, 12, 5, 8, 13, 7, 10, 12, 11, 13, 3, 9, 12, 7, 6, 4, 3, 11, 6, 13, 10, 7, 11, 1, 8, 5, 9, 10, 8, 9, 5, 4, 6},
	{7, 4, 9, 2, 11, 12, 3, 10, 4, 13, 8, 5, 11, 7, 9, 8, 12, 9, 3, 6, 13, 12, 7, 6, 13, 12, 8, 5, 1, 10, 11, 5, 6, 10, 13},
}

// reels lengths [35, 35, 35, 35, 35], total reshuffles 52521875
// symbols: 57.663(lined) + 14.71(scatter) = 72.373206%
// free spins 4345245, q = 0.082732, sq = 1/(1-q) = 1.090194
// free games frequency: 1/181.31
// RTP = 72.373(sym) + 0.082732*236.7(fg) = 91.956103%
var Reels919 = slot.Reels5x{
	{13, 11, 6, 8, 7, 4, 10, 12, 9, 13, 7, 8, 10, 1, 5, 6, 3, 13, 8, 7, 10, 11, 9, 4, 12, 9, 5, 13, 3, 2, 12, 4, 11, 6, 5},
	{7, 10, 8, 4, 13, 10, 12, 11, 9, 12, 13, 8, 10, 7, 13, 9, 5, 6, 9, 13, 11, 6, 5, 11, 4, 8, 6, 7, 1, 4, 3, 5, 12, 3, 2},
	{5, 3, 6, 11, 8, 9, 4, 7, 9, 2, 13, 10, 6, 1, 11, 9, 7, 8, 13, 7, 12, 8, 5, 13, 3, 4, 5, 10, 11, 13, 12, 10, 4, 6, 12},
	{13, 9, 8, 11, 12, 6, 9, 4, 5, 3, 12, 13, 7, 6, 10, 5, 7, 8, 13, 7, 11, 5, 10, 13, 9, 6, 4, 1, 10, 4, 8, 12, 2, 11, 3},
	{7, 4, 9, 2, 11, 12, 3, 10, 4, 13, 8, 5, 11, 7, 9, 8, 12, 9, 3, 6, 13, 12, 7, 6, 13, 12, 8, 5, 1, 10, 11, 5, 6, 10, 13},
}

// reels lengths [35, 35, 35, 35, 35], total reshuffles 52521875
// symbols: 58.133(lined) + 14.71(scatter) = 72.843439%
// free spins 4345245, q = 0.082732, sq = 1/(1-q) = 1.090194
// free games frequency: 1/181.31
// RTP = 72.843(sym) + 0.082732*238.24(fg) = 92.553572%
var Reels925 = slot.Reels5x{
	{13, 11, 6, 8, 7, 4, 10, 12, 9, 13, 7, 8, 10, 1, 5, 6, 3, 13, 8, 7, 10, 11, 9, 4, 12, 9, 5, 13, 3, 2, 12, 4, 11, 6, 5},
	{7, 10, 8, 4, 13, 10, 12, 11, 9, 12, 13, 8, 10, 7, 13, 9, 5, 6, 9, 13, 11, 6, 5, 11, 4, 8, 6, 7, 1, 4, 3, 5, 12, 3, 2},
	{5, 3, 6, 11, 8, 9, 4, 7, 9, 2, 13, 10, 6, 1, 11, 9, 7, 8, 13, 7, 12, 8, 5, 13, 3, 4, 5, 10, 11, 13, 12, 10, 4, 6, 12},
	{13, 9, 8, 11, 12, 6, 9, 4, 5, 3, 12, 13, 7, 6, 10, 5, 7, 8, 13, 7, 11, 5, 10, 13, 9, 6, 4, 1, 10, 4, 8, 12, 2, 11, 3},
	{13, 6, 12, 4, 10, 8, 2, 7, 13, 11, 4, 9, 6, 12, 11, 8, 13, 5, 6, 4, 7, 1, 10, 9, 3, 7, 10, 3, 8, 12, 5, 13, 9, 11, 5},
}

// reels lengths [35, 35, 35, 35, 35], total reshuffles 52521875
// symbols: 58.699(lined) + 14.71(scatter) = 73.409299%
// free spins 4345245, q = 0.082732, sq = 1/(1-q) = 1.090194
// free games frequency: 1/181.31
// RTP = 73.409(sym) + 0.082732*240.09(fg) = 93.272543%
var Reels932 = slot.Reels5x{
	{13, 4, 5, 10, 12, 1, 11, 6, 7, 8, 11, 3, 4, 9, 8, 5, 9, 10, 12, 8, 6, 5, 11, 4, 7, 12, 10, 13, 6, 3, 13, 9, 7, 3, 2},
	{7, 10, 8, 4, 13, 10, 12, 11, 9, 12, 13, 8, 10, 7, 13, 9, 5, 6, 9, 13, 11, 6, 5, 11, 4, 8, 6, 7, 1, 4, 3, 5, 12, 3, 2},
	{5, 3, 6, 11, 8, 9, 4, 7, 9, 2, 13, 10, 6, 1, 11, 9, 7, 8, 13, 7, 12, 8, 5, 13, 3, 4, 5, 10, 11, 13, 12, 10, 4, 6, 12},
	{13, 9, 8, 11, 12, 6, 9, 4, 5, 3, 12, 13, 7, 6, 10, 5, 7, 8, 13, 7, 11, 5, 10, 13, 9, 6, 4, 1, 10, 4, 8, 12, 2, 11, 3},
	{13, 6, 12, 4, 10, 8, 2, 7, 13, 11, 4, 9, 6, 12, 11, 8, 13, 5, 6, 4, 7, 1, 10, 9, 3, 7, 10, 3, 8, 12, 5, 13, 9, 11, 5},
}

// reels lengths [35, 35, 35, 35, 35], total reshuffles 52521875
// symbols: 59.339(lined) + 14.71(scatter) = 74.049603%
// free spins 4345245, q = 0.082732, sq = 1/(1-q) = 1.090194
// free games frequency: 1/181.31
// RTP = 74.05(sym) + 0.082732*242.19(fg) = 94.086102%
var Reels940 = slot.Reels5x{
	{13, 11, 6, 8, 7, 4, 10, 12, 9, 13, 7, 8, 10, 1, 5, 6, 3, 13, 8, 7, 10, 11, 9, 4, 12, 9, 5, 13, 3, 2, 12, 4, 11, 6, 5},
	{5, 8, 13, 12, 7, 13, 10, 2, 4, 3, 1, 6, 7, 3, 6, 11, 5, 9, 12, 8, 11, 9, 5, 6, 13, 9, 7, 4, 12, 10, 8, 4, 11, 3, 10},
	{5, 3, 6, 11, 8, 9, 4, 7, 9, 2, 13, 10, 6, 1, 11, 9, 7, 8, 13, 7, 12, 8, 5, 13, 3, 4, 5, 10, 11, 13, 12, 10, 4, 6, 12},
	{11, 9, 4, 5, 8, 4, 6, 13, 7, 5, 1, 10, 8, 12, 5, 10, 13, 6, 2, 13, 12, 11, 10, 3, 11, 7, 9, 3, 8, 4, 7, 12, 9, 3, 6},
	{13, 10, 2, 5, 9, 12, 8, 11, 12, 8, 3, 10, 4, 7, 9, 4, 7, 10, 4, 12, 1, 9, 6, 13, 3, 6, 5, 7, 8, 11, 13, 6, 11, 5, 3},
}

// reels lengths [35, 35, 35, 35, 35], total reshuffles 52521875
// symbols: 60.159(lined) + 14.71(scatter) = 74.869071%
// free spins 4345245, q = 0.082732, sq = 1/(1-q) = 1.090194
// free games frequency: 1/181.31
// RTP = 74.869(sym) + 0.082732*244.87(fg) = 95.127304%
var Reels951 = slot.Reels5x{
	{13, 4, 5, 10, 12, 1, 11, 6, 7, 8, 11, 3, 4, 9, 8, 5, 9, 10, 12, 8, 6, 5, 11, 4, 7, 12, 10, 13, 6, 3, 13, 9, 7, 3, 2},
	{5, 8, 13, 12, 7, 13, 10, 2, 4, 3, 1, 6, 7, 3, 6, 11, 5, 9, 12, 8, 11, 9, 5, 6, 13, 9, 7, 4, 12, 10, 8, 4, 11, 3, 10},
	{5, 3, 6, 11, 8, 9, 4, 7, 9, 2, 13, 10, 6, 1, 11, 9, 7, 8, 13, 7, 12, 8, 5, 13, 3, 4, 5, 10, 11, 13, 12, 10, 4, 6, 12},
	{13, 9, 8, 11, 12, 6, 9, 4, 5, 3, 12, 13, 7, 6, 10, 5, 7, 8, 13, 7, 11, 5, 10, 13, 9, 6, 4, 1, 10, 4, 8, 12, 2, 11, 3},
	{13, 6, 12, 4, 10, 8, 2, 7, 13, 11, 4, 9, 6, 12, 11, 8, 13, 5, 6, 4, 7, 1, 10, 9, 3, 7, 10, 3, 8, 12, 5, 13, 9, 11, 5},
}

// reels lengths [35, 35, 35, 35, 35], total reshuffles 52521875
// symbols: 61.759(lined) + 14.71(scatter) = 76.469404%
// free spins 4345245, q = 0.082732, sq = 1/(1-q) = 1.090194
// free games frequency: 1/181.31
// RTP = 76.469(sym) + 0.082732*250.1(fg) = 97.160658%
var Reels971 = slot.Reels5x{
	{13, 4, 5, 10, 12, 1, 11, 6, 7, 8, 11, 3, 4, 9, 8, 5, 9, 10, 12, 8, 6, 5, 11, 4, 7, 12, 10, 13, 6, 3, 13, 9, 7, 3, 2},
	{5, 8, 13, 12, 7, 13, 10, 2, 4, 3, 1, 6, 7, 3, 6, 11, 5, 9, 12, 8, 11, 9, 5, 6, 13, 9, 7, 4, 12, 10, 8, 4, 11, 3, 10},
	{13, 11, 8, 12, 2, 5, 11, 7, 9, 1, 10, 6, 13, 7, 8, 9, 4, 5, 12, 3, 6, 5, 10, 4, 3, 12, 11, 7, 9, 13, 3, 6, 4, 10, 8},
	{13, 9, 8, 11, 12, 6, 9, 4, 5, 3, 12, 13, 7, 6, 10, 5, 7, 8, 13, 7, 11, 5, 10, 13, 9, 6, 4, 1, 10, 4, 8, 12, 2, 11, 3},
	{13, 6, 12, 4, 10, 8, 2, 7, 13, 11, 4, 9, 6, 12, 11, 8, 13, 5, 6, 4, 7, 1, 10, 9, 3, 7, 10, 3, 8, 12, 5, 13, 9, 11, 5},
}

// reels lengths [35, 35, 35, 35, 35], total reshuffles 52521875
// symbols: 62.616(lined) + 14.71(scatter) = 77.326761%
// free spins 4345245, q = 0.082732, sq = 1/(1-q) = 1.090194
// free games frequency: 1/181.31
// RTP = 77.327(sym) + 0.082732*252.9(fg) = 98.250001%
var Reels982 = slot.Reels5x{
	{13, 4, 5, 10, 12, 1, 11, 6, 7, 8, 11, 3, 4, 9, 8, 5, 9, 10, 12, 8, 6, 5, 11, 4, 7, 12, 10, 13, 6, 3, 13, 9, 7, 3, 2},
	{5, 8, 13, 12, 7, 13, 10, 2, 4, 3, 1, 6, 7, 3, 6, 11, 5, 9, 12, 8, 11, 9, 5, 6, 13, 9, 7, 4, 12, 10, 8, 4, 11, 3, 10},
	{13, 11, 8, 12, 2, 5, 11, 7, 9, 1, 10, 6, 13, 7, 8, 9, 4, 5, 12, 3, 6, 5, 10, 4, 3, 12, 11, 7, 9, 13, 3, 6, 4, 10, 8},
	{11, 9, 4, 5, 8, 4, 6, 13, 7, 5, 1, 10, 8, 12, 5, 10, 13, 6, 2, 13, 12, 11, 10, 3, 11, 7, 9, 3, 8, 4, 7, 12, 9, 3, 6},
	{13, 6, 12, 4, 10, 8, 2, 7, 13, 11, 4, 9, 6, 12, 11, 8, 13, 5, 6, 4, 7, 1, 10, 9, 3, 7, 10, 3, 8, 12, 5, 13, 9, 11, 5},
}

// reels lengths [35, 35, 35, 35, 35], total reshuffles 52521875
// symbols: 63.291(lined) + 14.71(scatter) = 78.001195%
// free spins 4345245, q = 0.082732, sq = 1/(1-q) = 1.090194
// free games frequency: 1/181.31
// RTP = 78.001(sym) + 0.082732*255.11(fg) = 99.106924%
var Reels991 = slot.Reels5x{
	{13, 4, 5, 10, 12, 1, 11, 6, 7, 8, 11, 3, 4, 9, 8, 5, 9, 10, 12, 8, 6, 5, 11, 4, 7, 12, 10, 13, 6, 3, 13, 9, 7, 3, 2},
	{5, 8, 13, 12, 7, 13, 10, 2, 4, 3, 1, 6, 7, 3, 6, 11, 5, 9, 12, 8, 11, 9, 5, 6, 13, 9, 7, 4, 12, 10, 8, 4, 11, 3, 10},
	{13, 11, 8, 12, 2, 5, 11, 7, 9, 1, 10, 6, 13, 7, 8, 9, 4, 5, 12, 3, 6, 5, 10, 4, 3, 12, 11, 7, 9, 13, 3, 6, 4, 10, 8},
	{11, 9, 4, 5, 8, 4, 6, 13, 7, 5, 1, 10, 8, 12, 5, 10, 13, 6, 2, 13, 12, 11, 10, 3, 11, 7, 9, 3, 8, 4, 7, 12, 9, 3, 6},
	{13, 10, 2, 5, 9, 12, 8, 11, 12, 8, 3, 10, 4, 7, 9, 4, 7, 10, 4, 12, 1, 9, 6, 13, 3, 6, 5, 7, 8, 11, 13, 6, 11, 5, 3},
}

// reels lengths [35, 32, 32, 32, 32], total reshuffles 36700160
// symbols: 57.867(lined) + 17.074(scatter) = 74.940431%
// free spins 3730455, q = 0.10165, sq = 1/(1-q) = 1.113148
// free games frequency: 1/147.57
// RTP = 74.94(sym) + 0.10165*250.26(fg) = 100.378502%
var Reels100 = slot.Reels5x{
	{13, 4, 5, 10, 12, 1, 11, 6, 7, 8, 11, 3, 4, 9, 8, 5, 9, 10, 12, 8, 6, 5, 11, 4, 7, 12, 10, 13, 6, 3, 13, 9, 7, 3, 2},
	{10, 6, 3, 13, 6, 12, 9, 10, 7, 11, 13, 2, 3, 9, 7, 5, 8, 9, 6, 7, 10, 4, 8, 11, 13, 5, 12, 1, 8, 12, 11, 4},
	{9, 10, 12, 6, 13, 9, 7, 11, 13, 7, 3, 5, 8, 11, 12, 8, 5, 2, 12, 4, 9, 13, 7, 10, 6, 1, 10, 11, 3, 8, 4, 6},
	{3, 11, 7, 5, 1, 6, 8, 4, 9, 13, 6, 9, 8, 12, 11, 10, 5, 12, 10, 7, 13, 2, 11, 3, 10, 6, 12, 8, 9, 7, 13, 4},
	{4, 9, 2, 6, 9, 12, 11, 6, 7, 13, 3, 9, 8, 13, 10, 4, 1, 8, 10, 5, 13, 7, 10, 5, 12, 11, 3, 12, 8, 6, 11, 7},
}

// reels lengths [30, 32, 32, 32, 32], total reshuffles 31457280
// symbols: 57.503(lined) + 18.194(scatter) = 75.697686%
// free spins 3489480, q = 0.11093, sq = 1/(1-q) = 1.124768
// free games frequency: 1/135.22
// RTP = 75.698(sym) + 0.11093*255.43(fg) = 104.031580%
var Reels104 = slot.Reels5x{
	{12, 11, 4, 13, 12, 7, 5, 3, 1, 8, 6, 5, 10, 2, 3, 10, 11, 9, 10, 4, 11, 6, 8, 13, 9, 8, 12, 13, 9, 7},
	{10, 6, 3, 13, 6, 12, 9, 10, 7, 11, 13, 2, 3, 9, 7, 5, 8, 9, 6, 7, 10, 4, 8, 11, 13, 5, 12, 1, 8, 12, 11, 4},
	{9, 10, 12, 6, 13, 9, 7, 11, 13, 7, 3, 5, 8, 11, 12, 8, 5, 2, 12, 4, 9, 13, 7, 10, 6, 1, 10, 11, 3, 8, 4, 6},
	{3, 11, 7, 5, 1, 6, 8, 4, 9, 13, 6, 9, 8, 12, 11, 10, 5, 12, 10, 7, 13, 2, 11, 3, 10, 6, 12, 8, 9, 7, 13, 4},
	{4, 9, 2, 6, 9, 12, 11, 6, 7, 13, 3, 9, 8, 13, 10, 4, 1, 8, 10, 5, 13, 7, 10, 5, 12, 11, 3, 12, 8, 6, 11, 7},
}

// reels lengths [30, 30, 30, 32, 32], total reshuffles 27648000
// symbols: 59.797(lined) + 19.205(scatter) = 79.001421%
// free spins 3301560, q = 0.11941, sq = 1/(1-q) = 1.135608
// free games frequency: 1/125.61
// RTP = 79.001(sym) + 0.11941*269.14(fg) = 111.140979%
var Reels111 = slot.Reels5x{
	{12, 11, 4, 13, 12, 7, 5, 3, 1, 8, 6, 5, 10, 2, 3, 10, 11, 9, 10, 4, 11, 6, 8, 13, 9, 8, 12, 13, 9, 7},
	{9, 7, 1, 5, 10, 12, 8, 9, 10, 3, 11, 8, 5, 12, 6, 13, 11, 8, 7, 6, 10, 11, 2, 4, 3, 13, 12, 9, 13, 4},
	{13, 11, 10, 9, 11, 10, 7, 9, 8, 5, 9, 1, 12, 8, 10, 13, 12, 11, 4, 6, 7, 3, 5, 4, 8, 12, 6, 13, 2, 3},
	{3, 11, 7, 5, 1, 6, 8, 4, 9, 13, 6, 9, 8, 12, 11, 10, 5, 12, 10, 7, 13, 2, 11, 3, 10, 6, 12, 8, 9, 7, 13, 4},
	{4, 9, 2, 6, 9, 12, 11, 6, 7, 13, 3, 9, 8, 13, 10, 4, 1, 8, 10, 5, 13, 7, 10, 5, 12, 11, 3, 12, 8, 6, 11, 7},
}

// Map with available reels.
var ReelsMap = map[float64]*slot.Reels5x{
	85.775904:  &Reels857,
	87.094098:  &Reels870,
	88.204487:  &Reels882,
	89.247624:  &Reels892,
	90.224839:  &Reels902,
	91.066340:  &Reels910,
	91.956103:  &Reels919,
	92.553572:  &Reels925,
	93.272543:  &Reels932,
	94.086102:  &Reels940,
	95.127304:  &Reels951,
	97.160658:  &Reels971,
	98.250001:  &Reels982,
	99.106924:  &Reels991,
	100.378502: &Reels100,
	104.031580: &Reels104,
	111.140979: &Reels111,
}
