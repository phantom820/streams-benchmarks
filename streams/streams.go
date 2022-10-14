package streams

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"streams-benchmarks/benchmarks"
	"streams-benchmarks/data"
	"strings"
	"time"

	"github.com/phantom820/streams"
)

// generateSlice generate slice of integer from 1-> n
func generateSlice(n int) []int {

	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i + 1
	}
	return slice
}

func randomSlice(max, n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = rand.Intn(max)
	}
	return slice
}

// CountSuccesfulLogins how many numbers are divisible by 3 and 5 in the range [1,n].
func CountSuccesfulLogins(f *os.File, benchmark *benchmarks.Benchmark, sizesTable []int, concurrencyTable []int) {

	type config struct {
		concurrencyTable []int
	}

	configMap := map[string]config{
		"sequentialStream": {concurrencyTable: []int{1}},
		"concurrentStream": {concurrencyTable: concurrencyTable},
	}

	for key, config := range configMap {
		for _, size := range sizesTable {
			slice := data.LoadWebLogData()
			for _, concurrency := range config.concurrencyTable {
				r := benchmark.Benchmark(key, func() {
					stream := streams.FromSlice(func() []data.WebLog { return slice }, concurrency)
					stream.Filter(func(x data.WebLog) bool {
						return x.Status == "200" && strings.Contains(x.URL, "login")
					}).Count()
				})
				fmt.Fprintf(f, "%s,%s,%v,%v,%v\n", r.Name, "CountSuccesfulLogins", size, concurrency, r.Duration)
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
					stream := streams.FromSlice(func() []int { return slice }, concurrency)
					stream.Filter(func(x int) bool { return isPrime(x) }).Count()
				})
				fmt.Fprintf(f, "%s,%s,%v,%v,%v\n", r.Name, "CountPrimes", size, concurrency, r.Duration)
			}
		}
	}
}

// FrequencyCount Counts how many times a given value occurs in list of random integers.
func FrequencyCount(f *os.File, benchmark *benchmarks.Benchmark, sizesTable []int, concurrencyTable []int) {

	type config struct {
		concurrencyTable []int
	}

	configMap := map[string]config{
		"sequentialStream": {concurrencyTable: []int{1}},
		"concurrentStream": {concurrencyTable: concurrencyTable},
	}
	rand.Seed(time.Now().Unix())
	max := 1000 // interval for drawing values.
	for key, config := range configMap {
		for _, size := range sizesTable {
			slice := randomSlice(max, size)
			for _, concurrency := range config.concurrencyTable {
				r := benchmark.Benchmark(key, func() {
					key := rand.Intn(max)
					_ = streams.FromSlice(func() []int { return slice }, concurrency).
						Filter(func(x int) bool { return x == key }).Count()
				})
				fmt.Fprintf(f, "%s,%s,%v,%v,%v\n", r.Name, "FrequencyCount", size, concurrency, r.Duration)
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
					_, _ = streams.FromSlice(func() []int { return slice }, concurrency).Reduce(func(x, y int) int { return x + y })
				})
				fmt.Fprintf(f, "%s,%s,%v,%v,%v\n", r.Name, "Sum", size, concurrency, r.Duration)
			}
		}
	}
}

// Product compute product of the values in the range [1,n].
func Product(f *os.File, benchmark *benchmarks.Benchmark, sizesTable []int, concurrencyTable []int) {

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
					_, _ = streams.FromSlice(func() []int { return slice }, concurrency).Reduce(func(x, y int) int { return x * y })
				})
				fmt.Fprintf(f, "%s,%s,%v,%v,%v\n", r.Name, "Product", size, concurrency, r.Duration)
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

// Transformation on a custom data type.
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
					stream := streams.FromSlice(func() []Employee { return slice }, concurrency)
					stream.Filter(func(x Employee) bool { return x.age > 20 && x.rating > 0.5 }).Map(func(x Employee) Employee {
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
