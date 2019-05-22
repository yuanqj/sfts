package ta

import (
	"github.com/yuanqj/sfts/ts/run"
)

const (
	MACDPeriodFast   uint = 0
	MACDPeriodSlow   uint = 0
	MACDPeriodSignal uint = 0
)

type MACD struct {
	emaFast, emaSlow, emaSignal *run.EWMAvg
}

func NewMACD(periodFast, periodSlow, periodSignal uint) *MACD {
	if periodFast == MACDPeriodFast {
		periodFast = 12
	}
	if periodSlow == MACDPeriodSlow {
		periodSlow = 26
	}
	if periodSignal == MACDPeriodSignal {
		periodSignal = 9
	}
	return &MACD{
		emaFast:   run.NewEWMAvg(run.EWMSpan(float64(periodFast)), true),
		emaSlow:   run.NewEWMAvg(run.EWMSpan(float64(periodSlow)), true),
		emaSignal: run.NewEWMAvg(run.EWMSpan(float64(periodSignal)), true),
	}
}

func (macd *MACD) App(val float64) (dif, dem float64) {
	dif = macd.emaFast.App(val) - macd.emaSlow.App(val)
	dem = macd.emaSignal.App(dif)
	return
}
