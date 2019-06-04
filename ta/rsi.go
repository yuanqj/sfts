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
