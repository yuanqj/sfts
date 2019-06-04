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

type Sum struct {
	win, idx, cnt uint64
	vals          []float64
	sum           float64
}

type Max struct {
	win, idx, cnt uint64
	vals          []float64
}

type Min struct {
	win, idx, cnt uint64
	vals          []float64
}

type Avg struct {
	win, idx, cnt uint64
	vals          []float64
	sum           float64
}

type Fst struct {
	win, idx, cnt uint64
	vals          []float64
}

// Simple Linear Regression over timeline
type Slr struct {
	win, idx, cnt uint64
	vals, ws      []float64
	varX, sumY    float64
}

func NewSum(win uint64) *Sum {
	return &Sum{win: win, vals: make([]float64, win)}
}

func (run *Sum) App(val float64) (sum float64) {
	run.sum += val - run.vals[run.idx]
	run.vals[run.idx] = val
	if run.cnt < run.win {
		run.cnt++
	}
	run.idx = next(run.win, run.idx+1)
	return run.sum
}

func NewMax(win uint64) *Max {
	return &Max{win: win, vals: make([]float64, win)}
}

func (run *Max) App(val float64) (max float64) {
	run.vals[run.idx] = val
	if run.cnt < run.win {
		run.cnt++
	}
	max = run.vals[0]
	for i := uint64(1); i < run.cnt; i++ {
		if run.vals[i] > max {
			max = run.vals[i]
		}
	}
	run.idx = next(run.win, run.idx+1)
	return
}

func NewMin(win uint64) *Min {
	return &Min{win: win, vals: make([]float64, win)}
}

func (run *Min) App(val float64) (min float64) {
	run.vals[run.idx] = val
	if run.cnt < run.win {
		run.cnt++
	}
	min = run.vals[0]
	for i := uint64(1); i < run.cnt; i++ {
		if run.vals[i] < min {
			min = run.vals[i]
		}
	}
	run.idx = next(run.win, run.idx+1)
	return
}

func NewAvg(win uint64) *Avg {
	return &Avg{win: win, vals: make([]float64, win)}
}

func (run *Avg) App(val float64) (avg float64) {
	run.sum += val - run.vals[run.idx]
	if run.cnt < run.win {
		run.cnt++
	}
	run.vals[run.idx] = val
	run.idx = next(run.win, run.idx+1)
	return run.sum / float64(run.cnt)
}

func NewFst(win uint64) *Fst {
	return &Fst{win: win, vals: make([]float64, win)}
}

func (run *Fst) App(val float64) (fst float64) {
	run.vals[run.idx] = val
	if run.cnt < run.win {
		run.cnt++
	}
	run.idx = next(run.win, run.idx+1)
	if run.cnt == run.win {
		return run.vals[run.idx]
	} else {
		return run.vals[0]
	}
}

func NewSlr(win uint64) *Slr {
	run := &Slr{win: win, vals: make([]float64, win), ws: make([]float64, win)}
	avg := float64(win-1) / float64(2)
	for i := uint64(0); i < win; i++ {
		run.ws[i] = float64(i) - avg
		run.varX += run.ws[i] * run.ws[i]
	}
	return run
}

func (run *Slr) App(val float64) (slope float64) {
	run.sumY += val - run.vals[run.idx]
	run.vals[run.idx] = val
	run.idx = next(run.win, run.idx+1)
	if run.cnt < run.win-1 {
		run.cnt++
		return
	}
	cov, yAvg := float64(0), run.sumY/float64(run.win)
	for i := uint64(0); i < run.win; i++ {
		cov += run.ws[i] * (run.vals[next(run.win, run.idx+i)] - yAvg)
	}
	return cov / run.varX
}

func next(win, cnt uint64) uint64 {
	return cnt % win
}
