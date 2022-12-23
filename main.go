package main

import (
	"os"
	"streams-benchmarks/streams"
)

func main() {

	sizesTable := []int{1e2, 1e3, 1e4, 5e4, 1e5, 1e6}
	numberOfTrials := 5
	results, _ := os.Create("results.csv")

	streams.CountPrimes(results, sizesTable, numberOfTrials)
	streams.Sum(results, sizesTable, numberOfTrials)
	streams.WordCount(results, "data/tweets.csv", sizesTable, numberOfTrials)
	streams.VectorSum(results, sizesTable, numberOfTrials)
}
