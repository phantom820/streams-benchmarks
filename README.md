# streams-benchmarks

This provides some idea on the performance of streams that have been provided [streams](https://github.com/phantom820/streams) . Generally parallel streams work better than sequential streams when the cost of processing a single element is high while sequential streams are better when the cost of processing a single element is low.

| Name      | Description |
| ----------- | ----------- |
| CountPrimes | Counts the number of primes in the range [1, source size]|
| WordCount| Counts the number of times a word occurs in tweets dataset. |
| Sum | Computes the sum of the values in the range [1,source size]. |
| NormalizeVector | Goes through a number of vectors and normalizes each vector.|



![streams](https://github.com/phantom820/streams-benchmarks/blob/main/visuals/streams.png)
