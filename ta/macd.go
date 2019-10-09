// Copyright 2019 yuanqj <yuanqj8191@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ta

import (
	"github.com/yuanqj/sfts/ts/run"
)

const (
	MACDPeriodFast   uint = 12
	MACDPeriodSlow   uint = 26
	MACDPeriodSignal uint = 9
)

type MACD struct {
	emaFast, emaSlow, emaSignal *run.EWMAvg
}

func NewMACD(periodFast, periodSlow, periodSignal uint) *MACD {
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
