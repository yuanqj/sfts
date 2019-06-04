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

package run

import (
	"math"
)

// Commonly called EMA.
type EWMAvg struct {
	alpha, decay float64
	adjust       bool
	sum, cnt     float64
	avg, ini     float64
}

func NewEWMAvg(alpha float64, adjust bool) *EWMAvg {
	if alpha < 1e-13 || alpha > 1 {
		panic("`alpha` must be between 0 and 1 inclusive")
	}
	return &EWMAvg{alpha: alpha, decay: 1 - alpha, adjust: adjust, ini: alpha}
}

func (ewm *EWMAvg) App(val float64) (avg float64) {
	if ewm.adjust {
		ewm.cnt = ewm.decay*ewm.cnt + 1
		ewm.sum = ewm.decay*ewm.sum + val
		ewm.avg = ewm.sum / ewm.cnt
	} else {
		ewm.avg = ewm.decay*ewm.avg + ewm.alpha*val/ewm.ini
		ewm.ini = 1
	}
	return ewm.avg
}

// The period of time for the exponential weight to reduce to one half.
func EWMHalflife(halflife float64) (alpha float64) {
	if halflife < 1e-13 {
		panic("`halflife` must be greater than 0")
	}
	return 1 - math.Exp(math.Log(0.5)/halflife)
}

// Commonly called an "N-day exponential weighted moving average".
func EWMSpan(span float64) (alpha float64) {
	if span < 1 {
		panic("`span` must be no smaller than 1")
	}
	return 2 / (span + 1)
}

// Center of mass.
func EWMCom(com float64) (alpha float64) {
	if com < 0 {
		panic("`span` must be no smaller than 0")
	}
	return 1 / (com + 1)
}
