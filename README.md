# streams-benchmarks

This provides some idea on the performance of streams that have been provided [streams](https://github.com/phantom820/streams) . Generally concurrent streams work better than sequential streams when thecost of processing a single element is high while sequential streams are better when the cost of processing a single element is low.

| Name      | Description |
| ----------- | ----------- |
| CountSuccesFul logins| Counts the number of succesful logins from given web log data.|
| CountPrimes | Counts the number of primes in the range [1, source size]|
| FrequencyCount| Counts how many times a given value occurs in list of random integers. |
| Sum | Computes the sum of the values in the range [1,source size]. |
| Product | Computes the product of the values in the range [1,source size].|
| Transformation | Performs a sequence of operation in deriving new instance of a type from given instance. |



![streams](https://github.com/phantom820/streams-benchmarks/blob/main/visuals/streams.png)
