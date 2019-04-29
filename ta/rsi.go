package ta

import (
	"github.com/yuanqj8191/sfts/ts/run"
)

const RSIPeriod uint = 14

type RSI struct {
	ini, val   float64
	avgU, avgD *run.EWMAvg
}

func NewRSI(period uint) *RSI {
	alpha := run.EWMCom(float64(period - 1))
	return &RSI{
		avgU: run.NewEWMAvg(alpha, false),
		avgD: run.NewEWMAvg(alpha, false),
	}
}

func (rsi *RSI) App(val float64) float64 {
	dv, du, dd := (val-rsi.val)*rsi.ini, 0., 0.
	rsi.val, rsi.ini = val, 1
	if dv > +1e-13 {
		du = +dv
	}
	if dv < -1e-13 {
		dd = -dv
	}
	u, d := rsi.avgU.App(du), rsi.avgD.App(dd)
	tot := u + d
	if tot < 1e-13 {
		return 50
	}
	return u / tot * 100
}
