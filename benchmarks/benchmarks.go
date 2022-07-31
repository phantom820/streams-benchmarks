package benchmarks

import (
	"time"
)

type Timer struct {
	start time.Time
	end   time.Time
}

func (timer *Timer) Start() {
	timer.start = time.Now()
}

func (timer *Timer) Stop() {
	timer.end = time.Now()
}

func (timer *Timer) Duration() float32 {
	return float32(timer.end.Sub(timer.start).Nanoseconds())
}

type Setup struct {
	count      int // number of measurements to aggregate.
	trials     int // number of samples needed to produce a single measurement.
	operations int // number of operations to do for number of operations based benchmarks.
	size       int // size of collections to be used. (size>=count * trial* operations)
}

func (setup *Setup) Count() int {
	return setup.count
}

func (setup *Setup) Trials() int {
	return setup.trials
}

func (setup *Setup) Operations() int {
	return setup.operations
}

func (setup *Setup) Size() int {
	return setup.size
}

func NewSetup(count, trials, operations, size int) *Setup {
	return &Setup{count: count, trials: trials, operations: operations, size: size}
}

type Benchmark struct {
	setup *Setup
	timer *Timer
}

func (benchmark *Benchmark) Setup() *Setup {
	return benchmark.setup
}

func NewBenchmark(setup *Setup, timer *Timer) *Benchmark {
	return &Benchmark{setup: setup, timer: timer}
}

func (benchmark *Benchmark) Benchmark(name string, f func()) Result {

	trial := func() float32 {
		sample := float32(0.0)
		for j := 0; j < benchmark.setup.trials; j++ {
			benchmark.timer.Start()
			f()
			benchmark.timer.Stop()
			sample += (benchmark.timer.Duration())
		}
		return sample / float32(benchmark.setup.trials)
	}

	duration := float32(0.0)
	for i := 0; i < benchmark.setup.count; i++ {
		duration += trial()
	}
	return Result{Name: name, Duration: duration / float32(benchmark.setup.count)}
}

type Result struct {
	Name     string
	Duration float32
}
