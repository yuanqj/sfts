package run

type RunSum struct {
	win, idx, cnt uint
	vals          []float64
	sum           float64
}

type RunMax struct {
	win, idx, cnt uint
	vals          []float64
}

type RunMin struct {
	win, idx, cnt uint
	vals          []float64
}

type RunAvg struct {
	win, idx, cnt uint
	vals          []float64
	sum           float64
}

type RunFst struct {
	win, idx, cnt uint
	vals          []float64
}

func NewRunSum(win uint) *RunSum {
	return &RunSum{win: win, vals: make([]float64, win)}
}

func (run *RunSum) Set(val float64) (sum float64) {
	run.sum += val - run.vals[run.idx]
	run.vals[run.idx] = val
	if run.cnt < run.win {
		run.cnt++
	}
	run.idx = cnt2idx(run.win, run.idx+1)
	return run.sum
}

func NewRunMax(win uint) *RunMax {
	return &RunMax{win: win, vals: make([]float64, win)}
}

func (run *RunMax) Set(val float64) (max float64) {
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
	run.idx = cnt2idx(run.win, run.idx+1)
	return
}

func NewRunMin(win uint) *RunMin {
	return &RunMin{win: win, vals: make([]float64, win)}
}

func (run *RunMin) Set(val float64) (min float64) {
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
	run.idx = cnt2idx(run.win, run.idx+1)
	return
}

func NewRunAvg(win uint) *RunAvg {
	return &RunAvg{win: win, vals: make([]float64, win)}
}

func (run *RunAvg) Set(val float64) (avg float64) {
	run.sum += val - run.vals[run.idx]
	if run.cnt < run.win {
		run.cnt++
	}
	run.vals[run.idx] = val
	run.idx = cnt2idx(run.win, run.idx+1)
	return run.sum / float64(run.cnt)
}

func NewRunFst(win uint) *RunFst {
	return &RunFst{win: win, vals: make([]float64, win)}
}

func (run *RunFst) Set(val float64) (fst float64) {
	run.vals[run.idx] = val
	if run.cnt < run.win {
		run.cnt++
	}
	run.idx = cnt2idx(run.win, run.idx+1)
	if run.cnt == run.win {
		return run.vals[run.idx]
	} else {
		return run.vals[0]
	}
}

func cnt2idx(win, cnt uint) uint {
	return cnt % win
}
