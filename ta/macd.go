package ta

import (
	"github.com/yuanqj8191/sfts/ts/run"
)

type MACD struct {
	emaFast, emaSlow, emaSignal *run.EWMAvg
}

func NewMACD(spanFast, spanSlow, spanSignal uint64) *MACD {
	return &MACD{
		emaFast:   run.NewEWMAvg(run.EWMSpan(float64(spanFast)), true),
		emaSlow:   run.NewEWMAvg(run.EWMSpan(float64(spanSlow)), true),
		emaSignal: run.NewEWMAvg(run.EWMSpan(float64(spanSignal)), true),
	}
}

func (macd *MACD) App(val float64) (dif, dem float64) {
	dif = macd.emaFast.App(val) - macd.emaSlow.App(val)
	dem = macd.emaSignal.App(dif)
	return
}
