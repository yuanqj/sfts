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

type WinSum struct {
	win, idx, cnt uint
	vals          []float64
	sum           float64
}

type WinMax struct {
	win, idx, cnt uint
	vals          []float64
}

type WinMin struct {
	win, idx, cnt uint
	vals          []float64
}

type WinAvg struct {
	win, idx, cnt uint
	vals          []float64
	sum           float64
}

type WinFst struct {
	win, idx, cnt uint
	vals          []float64
}

// Simple Linear Regression over timeline
type WinSlr struct {
	win, idx, cnt uint
	vals, ws      []float64
	varX, sumY    float64
}

func NewWinSum(win uint) *WinSum {
	return &WinSum{win: win, vals: make([]float64, win)}
}

func (run *WinSum) App(val float64) (sum float64) {
	run.sum += val - run.vals[run.idx]
	run.vals[run.idx] = val
	if run.cnt < run.win {
		run.cnt++
	}
	run.idx = next(run.win, run.idx+1)
	return run.sum
}

func NewWinMax(win uint) *WinMax {
	return &WinMax{win: win, vals: make([]float64, win)}
}

func (run *WinMax) App(val float64) (max float64) {
	run.vals[run.idx] = val
	if run.cnt < run.win {
		run.cnt++
	}
	max = run.vals[0]
	for i := uint(1); i < run.cnt; i++ {
		if run.vals[i] > max {
			max = run.vals[i]
		}
	}
	run.idx = next(run.win, run.idx+1)
	return
}

func NewWinMin(win uint) *WinMin {
	return &WinMin{win: win, vals: make([]float64, win)}
}

func (run *WinMin) App(val float64) (min float64) {
	run.vals[run.idx] = val
	if run.cnt < run.win {
		run.cnt++
	}
	min = run.vals[0]
	for i := uint(1); i < run.cnt; i++ {
		if run.vals[i] < min {
			min = run.vals[i]
		}
	}
	run.idx = next(run.win, run.idx+1)
	return
}

func NewWinAvg(win uint) *WinAvg {
	return &WinAvg{win: win, vals: make([]float64, win)}
}

func (run *WinAvg) App(val float64) (avg float64) {
	run.sum += val - run.vals[run.idx]
	if run.cnt < run.win {
		run.cnt++
	}
	run.vals[run.idx] = val
	run.idx = next(run.win, run.idx+1)
	return run.sum / float64(run.cnt)
}

func NewWinFst(win uint) *WinFst {
	return &WinFst{win: win, vals: make([]float64, win)}
}

func (run *WinFst) App(val float64) (fst float64) {
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

func NewWinSlr(win uint) *WinSlr {
	run := &WinSlr{win: win, vals: make([]float64, win), ws: make([]float64, win)}
	avg := float64(win-1) / float64(2)
	for i := uint(0); i < win; i++ {
		run.ws[i] = float64(i) - avg
		run.varX += run.ws[i] * run.ws[i]
	}
	return run
}

func (run *WinSlr) App(val float64) (slope float64) {
	run.sumY += val - run.vals[run.idx]
	run.vals[run.idx] = val
	run.idx = next(run.win, run.idx+1)
	if run.cnt < run.win-1 {
		run.cnt++
		return
	}
	cov, yAvg := float64(0), run.sumY/float64(run.win)
	for i := uint(0); i < run.win; i++ {
		cov += run.ws[i] * (run.vals[next(run.win, run.idx+i)] - yAvg)
	}
	return cov / run.varX
}

func next(win, cnt uint) uint {
	return cnt % win
}
