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
	"math"
)

// RSI momentum types
const (
	RSI_MOM_Linear RSI_MOM = iota
	RSI_MOM_LinearRatio
	RSI_MOM_LogRatio
)

const RSIPeriod uint = 14

type RSI_MOM uint                  // RSI momentum
type rsi_mom func(float64) float64 // RSI momentum calculators

type RSI struct {
	cnt, periodSmooth  uint
	val                float64
	mom                rsi_mom
	avgU, avgD, smooth *run.EWMAvg
}

func NewRSI(mom RSI_MOM, period, periodSmooth uint) *RSI {
	alpha := run.EWMCom(float64(period - 1))
	var smooth *run.EWMAvg
	if periodSmooth > 0 {
		smooth = run.NewEWMAvg(run.EWMHalflife(float64(periodSmooth)), true)
	}
	rsi := &RSI{
		avgU:         run.NewEWMAvg(alpha, true),
		avgD:         run.NewEWMAvg(alpha, true),
		periodSmooth: periodSmooth,
		smooth:       smooth,
	}
	moms := []rsi_mom{rsi.momLinear, rsi.momLinearRatio, rsi.momLogRatio}
	rsi.mom = moms[mom]
	return rsi
}

func (rsi *RSI) Reset(val float64) {
	rsi.val = val
}

func (rsi *RSI) App(val float64) float64 {
	mom, u, d := rsi.mom(val), 0., 0.
	if mom > +1e-13 {
		u = +mom
	}
	if mom < -1e-13 {
		d = -mom
	}
	u, d = rsi.avgU.App(u), rsi.avgD.App(d)
	tot, idc := u+d, 50.
	if rsi.cnt > rsi.periodSmooth && tot > 1e-13 {
		idc = u / tot * 100
	}
	if rsi.smooth != nil {
		idc = rsi.smooth.App(idc)
	}

	rsi.cnt++
	rsi.val = val
	return idc
}

func (rsi *RSI) momLinear(val float64) float64 {
	return val - rsi.val
}

func (rsi *RSI) momLinearRatio(val float64) float64 {
	return (val - rsi.val) / rsi.val
}

func (rsi *RSI) momLogRatio(val float64) float64 {
	return math.Log(val / rsi.val)
}
