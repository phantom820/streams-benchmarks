package streams

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"regexp"
	"streams-benchmarks/data"
	"strings"
	"time"

	"github.com/phantom820/streams"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// sequence generates slice of integer from 1-> n
func sequence(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i + 1
	}
	return slice
}

// randomInts generate a slice of random integers in the range [0,n).
func randomInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = rand.Intn(n)
	}
	return slice
}

// randomString generates a random string of the given length.
func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// randomStrings generates a slice of N random strings of length n.
func randomStrings(n int, N int) []string {
	slice := make([]string, N)
	for i := 0; i < n; i++ {
		slice[i] = randomString(n)
	}
	return slice
}

// randomFloats generate a slice of random floats in the range [0,1).
func randomFloats(n int) []float64 {
	slice := make([]float64, n)
	for i := 0; i < n; i++ {
		slice[i] = rand.Float64()
	}
	return slice
}

// randomVectors generates N random vectors in dimension n.
func randomVectors(N, n int) [][]float64 {
	slice := make([][]float64, N)
	for i := 0; i < N; i++ {
		slice[i] = randomFloats(n)
	}
	return slice
}

type benchmark[T any] struct {
	name        string
	parallelism int
	run         func(supplier func() []T)
}

func runBenchmark[T any](numberOfTrials int, data func() []T, benchmark benchmark[T]) float64 {
	totalDuration := 0.0
	for i := 0; i < numberOfTrials; i++ {
		startTime := time.Now()
		benchmark.run(data)
		totalDuration = totalDuration + float64(time.Since(startTime))
	}
	return totalDuration / float64(numberOfTrials)
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	} else if n == 2 || n == 3 || n == 5 {
		return true
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}
	sqrt := math.Ceil(math.Sqrt(float64(n)))
	for i := 4; i <= int(sqrt); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// CountsPrimes counts how many numbers are primes range [1,n].
func CountPrimes(f *os.File, sizesTable []int, numberOfTrials int) {

	benchmarks := []benchmark[int]{
		{name: "CountPrimes",
			parallelism: 1,
			run:         func(supplier func() []int) { streams.New(supplier).Filter(isPrime).Count() }},
		{name: "CountPrimes",
			parallelism: 2,
			run:         func(supplier func() []int) { streams.New(supplier).Parallelize(2).Filter(isPrime).Count() }},
		{name: "CountPrimes",
			parallelism: 4,
			run:         func(supplier func() []int) { streams.New(supplier).Parallelize(4).Filter(isPrime).Count() }},
		{name: "CountPrimes",
			parallelism: 8,
			run:         func(supplier func() []int) { streams.New(supplier).Parallelize(8).Filter(isPrime).Count() }},
		{name: "CountPrimes",
			parallelism: 16,
			run:         func(supplier func() []int) { streams.New(supplier).Parallelize(16).Filter(isPrime).Count() }},
	}

	for _, size := range sizesTable {
		supplier := func() []int {
			return sequence(size)
		}
		for _, benchmark := range benchmarks {
			duration := runBenchmark(numberOfTrials, supplier, benchmark)
			fmt.Fprintf(f, "%s,%v,%v,%v\n", "CountPrimes", size, benchmark.parallelism, duration)
		}
	}
}

// Sum sums the values in the range [1,n].
func Sum(f *os.File, sizesTable []int, numberOfTrials int) {

	sum := func(x, y int) int {
		return x + y
	}

	benchmarks := []benchmark[int]{
		{name: "Sum",
			parallelism: 1,
			run:         func(supplier func() []int) { streams.New(supplier).Reduce(sum) }},
		{name: "Sum",
			parallelism: 2,
			run:         func(supplier func() []int) { streams.New(supplier).Parallelize(2).Reduce(sum) }},
		{name: "Sum",
			parallelism: 4,
			run:         func(supplier func() []int) { streams.New(supplier).Parallelize(4).Reduce(sum) }},
		{name: "Sum",
			parallelism: 8,
			run:         func(supplier func() []int) { streams.New(supplier).Parallelize(8).Reduce(sum) }},
		{name: "Sum",
			parallelism: 16,
			run:         func(supplier func() []int) { streams.New(supplier).Parallelize(16).Reduce(sum) }},
	}

	for _, size := range sizesTable {
		supplier := func() []int {
			return sequence(size)
		}
		for _, benchmark := range benchmarks {
			duration := runBenchmark(numberOfTrials, supplier, benchmark)
			fmt.Fprintf(f, "%s,%v,%v,%v\n", "Sum", size, benchmark.parallelism, duration)
		}
	}
}

// WordCount counts how many times a word occurs in tweets dataset.
func WordCount(f *os.File, filePath string, sizesTable []int, numberOfTrials int) {

	reg, _ := regexp.Compile("[^A-Za-z0-9]+")
	cleanString := func(x string) string {
		return strings.ToLower(reg.ReplaceAllString(x, ""))
	}

	split := func(x string) []string {
		return strings.Split(x, " ")
	}

	group := func(x string) string {
		return x
	}

	tweets := data.ReadTweets(filePath)
	benchmarks := []benchmark[string]{
		{name: "WordCount",
			parallelism: 1,
			run: func(supplier func() []string) {
				streams.New(supplier).
					Partition(split).
					Map(cleanString).
					FlatMap().
					GroupBy(group).
					Count()
			}},
		{name: "WordCount",
			parallelism: 2,
			run: func(supplier func() []string) {
				streams.New(supplier).Partition(split).
					Parallelize(2).
					Map(cleanString).
					FlatMap().
					GroupBy(group).
					Count()
			}},
		{name: "WordCount",
			parallelism: 4,
			run: func(supplier func() []string) {
				streams.New(supplier).
					Partition(split).
					Parallelize(4).
					Map(cleanString).
					FlatMap().
					GroupBy(group).
					Count()
			}},
		{name: "WordCount",
			parallelism: 8,
			run: func(supplier func() []string) {
				streams.New(supplier).
					Partition(split).
					Parallelize(8).
					Map(cleanString).
					FlatMap().
					GroupBy(group).
					Count()
			}},
		{name: "WordCount",
			parallelism: 16,
			run: func(supplier func() []string) {
				streams.New(supplier).
					Partition(split).
					Parallelize(16).
					Map(cleanString).
					FlatMap().
					GroupBy(group).
					Count()
			}},
	}

	for _, size := range sizesTable {
		supplier := func() []string {
			return tweets[:size]
		}
		for _, benchmark := range benchmarks {
			duration := runBenchmark(numberOfTrials, supplier, benchmark)
			fmt.Fprintf(f, "%s,%v,%v,%v\n", "WordCount", size, benchmark.parallelism, duration)
		}
	}
}

// NormalizeVector normalizes a bunch of vectors.
func VectorSum(f *os.File, sizesTable []int, numberOfTrials int) {

	norm := func(x []float64) float64 {
		norm := 0.0
		for _, val := range x {
			norm = norm + math.Pow(val, 2)
		}
		return math.Sqrt(norm)
	}

	normalize := func(x []float64) []float64 {
		y := make([]float64, len(x))
		norm := norm(x)
		for i := range x {
			y[i] = (x[i]) / norm
		}
		return y
	}

	benchmarks := []benchmark[[]float64]{
		{name: "NormalizeVector",
			parallelism: 1,
			run:         func(supplier func() [][]float64) { streams.New(supplier).Map(normalize).Collect() }},
		{name: "NormalizeVector",
			parallelism: 2,
			run:         func(supplier func() [][]float64) { streams.New(supplier).Parallelize(2).Map(normalize).Collect() }},
		{name: "NormalizeVector",
			parallelism: 4,
			run:         func(supplier func() [][]float64) { streams.New(supplier).Parallelize(4).Map(normalize).Collect() }},
		{name: "NormalizeVector",
			parallelism: 8,
			run:         func(supplier func() [][]float64) { streams.New(supplier).Parallelize(8).Map(normalize).Collect() }},
		{name: "NormalizeVector",
			parallelism: 16,
			run:         func(supplier func() [][]float64) { streams.New(supplier).Parallelize(16).Map(normalize).Collect() }},
	}

	for _, size := range sizesTable {
		supplier := func() [][]float64 {
			return randomVectors(size, 48)
		}
		for _, benchmark := range benchmarks {
			duration := runBenchmark(numberOfTrials, supplier, benchmark)
			fmt.Fprintf(f, "%s,%v,%v,%v\n", "NormalizeVector", size, benchmark.parallelism, duration)
		}
	}
}
