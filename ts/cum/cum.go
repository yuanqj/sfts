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

package cum

import "math"

type Cum interface {
	App(val float64) float64
}

type Sum struct {
	sum float64
}

type Prd struct {
	prd float64
}

type Max struct {
	max float64
}

type Min struct {
	min float64
}

type Cnt struct {
	cnt float64
}

type Avg struct {
	sum, cnt float64
}

func NewSum() *Sum {
	return &Sum{}
}

func (cum *Sum) App(val float64) (sum float64) {
	cum.sum += val
	return cum.sum
}

func NewPrd() *Prd {
	return &Prd{prd: 1}
}

func (cum *Prd) App(val float64) (prd float64) {
	cum.prd *= val
	return cum.prd
}

func NewMax() *Max {
	return &Max{max: math.Inf(-1)}
}

func (cum *Max) App(val float64) (max float64) {
	if cum.max < val {
		cum.max = val
	}
	return cum.max
}

func NewMin() *Min {
	return &Min{min: math.Inf(1)}
}

func (cum *Min) App(val float64) (min float64) {
	if cum.min > val {
		cum.min = val
	}
	return cum.min
}

func NewCnt() *Cnt {
	return &Cnt{}
}

func (cum *Cnt) App(val float64) (cnt float64) {
	cum.cnt++
	return cum.cnt
}

func NewAvg() *Avg {
	return &Avg{}
}

func (cum *Avg) App(val float64) (avg float64) {
	cum.cnt++
	cum.sum += val
	return cum.sum / cum.cnt
}
