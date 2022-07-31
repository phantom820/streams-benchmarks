package streams

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"streams-benchmarks/benchmarks"
	"strings"

	"github.com/phantom820/streams"
)

func generateSlice(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i + 1
	}
	return slice
}

// CountFizzBuzz counts how many numbers are divisible by 3 and 5 in the range [1,n].
func CountFizzBuzz(f *os.File, benchmark *benchmarks.Benchmark, sizesTable []int, concurrencyTable []int) {

	type config struct {
		concurrencyTable []int
	}

	configMap := map[string]config{
		"sequentialStream": {concurrencyTable: []int{1}},
		"concurrentStream": {concurrencyTable: concurrencyTable},
	}

	for key, config := range configMap {
		for _, size := range sizesTable {
			slice := generateSlice(size)
			for _, concurrency := range config.concurrencyTable {
				r := benchmark.Benchmark(key, func() {
					stream := streams.NewFromSlice(func() []int { return slice }, concurrency)
					stream.Filter(func(x int) bool { return x%3 == 0 && x%5 == 0 }).Count()
				})
				fmt.Fprintf(f, "%s,%s,%v,%v,%v\n", r.Name, "CountFizzBuzz", size, concurrency, r.Duration)
			}
		}
	}

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
func CountPrimes(f *os.File, benchmark *benchmarks.Benchmark, sizesTable []int, concurrencyTable []int) {

	type config struct {
		concurrencyTable []int
	}

	configMap := map[string]config{
		"sequentialStream": {concurrencyTable: []int{1}},
		"concurrentStream": {concurrencyTable: concurrencyTable},
	}

	for key, config := range configMap {
		for _, size := range sizesTable {
			slice := generateSlice(size)
			for _, concurrency := range config.concurrencyTable {
				r := benchmark.Benchmark(key, func() {
					stream := streams.NewFromSlice(func() []int { return slice }, concurrency)
					stream.Filter(func(x int) bool { return isPrime(x) }).Count()
				})
				fmt.Fprintf(f, "%s,%s,%v,%v,%v\n", r.Name, "CountPrimes", size, concurrency, r.Duration)
			}
		}
	}
}

// Sum sums the values in the range [1,n].
func Sum(f *os.File, benchmark *benchmarks.Benchmark, sizesTable []int, concurrencyTable []int) {

	type config struct {
		concurrencyTable []int
	}

	configMap := map[string]config{
		"sequentialStream": {concurrencyTable: []int{1}},
		"concurrentStream": {concurrencyTable: concurrencyTable},
	}

	for key, config := range configMap {
		for _, size := range sizesTable {
			slice := generateSlice(size)
			for _, concurrency := range config.concurrencyTable {
				r := benchmark.Benchmark(key, func() {
					stream := streams.NewFromSlice(func() []int { return slice }, concurrency)
					stream.Map(func(x int) interface{} { return x * x }).Reduce(func(x, y interface{}) interface{} { return x.(int) + y.(int) })
				})
				fmt.Fprintf(f, "%s,%s,%v,%v,%v\n", r.Name, "Sum", size, concurrency, r.Duration)
			}
		}
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type Employee struct {
	id      int
	name    string
	surname string
	age     int
	rating  float32
}

func generateEmployee() Employee {
	return Employee{
		id:      rand.Intn(1e6),
		name:    randString(100 + rand.Intn(100)),
		surname: randString(100 + rand.Intn(100)),
		age:     16 + rand.Intn(16),
		rating:  rand.Float32(),
	}
}

func generateEmployees(n int) []Employee {
	employees := make([]Employee, n)
	for i := 0; i < n; i++ {
		employees[i] = generateEmployee()
	}

	return employees
}

// Transformation
func Transformation(f *os.File, benchmark *benchmarks.Benchmark, sizesTable []int, concurrencyTable []int) {

	type config struct {
		concurrencyTable []int
	}

	configMap := map[string]config{
		"sequentialStream": {concurrencyTable: []int{1}},
		"concurrentStream": {concurrencyTable: concurrencyTable},
	}

	for key, config := range configMap {
		for _, size := range sizesTable {
			slice := generateEmployees(size)
			for _, concurrency := range config.concurrencyTable {
				r := benchmark.Benchmark(key, func() {
					stream := streams.NewFromSlice(func() []Employee { return slice }, concurrency)
					stream.Filter(func(x Employee) bool { return x.age > 20 && x.rating > 0.5 }).Map(func(x Employee) interface{} {
						newEmployee := Employee{
							id:      x.id,
							name:    strings.ToUpper(x.name),
							surname: strings.Repeat(x.surname, 2),
							age:     x.age,
							rating:  x.rating * 100,
						}
						return newEmployee
					}).Skip(10).Collect()
				})
				fmt.Fprintf(f, "%s,%s,%v,%v,%v\n", r.Name, "Transformation", size, concurrency, r.Duration)
			}
		}
	}
}
