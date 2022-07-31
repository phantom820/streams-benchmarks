package main

import (
	"os"
	"streams-benchmarks/benchmarks"
	"streams-benchmarks/streams"
)

const (
	trials     = 1
	count      = 1
	operations = 30000
	size       = 60000
)

func Benchmark(sizesTable, concurrencyTable []int) {

	timer := benchmarks.Timer{}
	setup := benchmarks.NewSetup(trials, count, operations, size)
	benchmark := benchmarks.NewBenchmark(setup, &timer)

	results, _ := os.Create("results.csv")
	streams.CountFizzBuzz(results, benchmark, sizesTable, concurrencyTable)
	streams.CountPrimes(results, benchmark, sizesTable, concurrencyTable)
	streams.Sum(results, benchmark, sizesTable, concurrencyTable)
	streams.Transformation(results, benchmark, sizesTable, concurrencyTable)

}

func main() {

	sizesTable := []int{1e5, 2 * 1e5, 5 * 1e5, 1e6}
	concurrencyTable := []int{2, 4, 8, 16}
	Benchmark(sizesTable, concurrencyTable)
}
