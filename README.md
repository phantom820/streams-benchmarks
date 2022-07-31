# streams-benchmarks

This provides some idea on the performance of streams that have been provided [streams](https://github.com/phantom820/streams) . Generally concurrent streams work better than sequential streams when thecost of processing a single element is high while sequential streams are better when the cost of processing a single element is low.

#### Sequential vs Concurrent streams
| Sequential      | Concurrent |
| ----------- | ----------- |
| Processes its elements sequentially (max concurreny = 1) .    | Processes its elements concurrently using no more than a specified number of go routines (max concurrency > 1).     |
| Preserves encounter order from the source  | Does not preserve encounter order from the source.      |
| Infinite source will work if limit operation is applied. | Infinite source will not work in any case |
| Reduce operation does not require function to be commutative. | Reduce results may not make sense if given function is not commutative.|
| Performs well when cost of processing an element low | Performs well when cost of processing a single element is high.|
| Limit, Skip & Distinct operations are cheap | Limit,Skip & Distinct operations are expensive due to locks | 


![streams](https://user-images.githubusercontent.com/47748901/182045093-0361bbc5-dd19-4ea6-901e-26bf67bb043d.png)
