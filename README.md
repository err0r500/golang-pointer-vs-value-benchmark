# Impact of Pointers vs Values on CPU

The structure tested contains 20 fields ( int64 )

```
go test -bench=.
```

| Benchmarks on Exported  | # of op. | dur. per op. |
| ------------- | ------------- | ------------- |
|  BenchmarkNewValue-8  |  100000000  |  10.8 ns/op  |
|  BenchmarkNewPointer-8  |  300000000  |  5.33 ns/op  |
|  BenchmarkFuncValue-8  |  300000000  |  4.88 ns/op  |
|  BenchmarkFuncPointer-8  |  2000000000  |  0.38 ns/op  |
|  BenchmarkMethodValue-8  |  300000000  |  4.87 ns/op  |
|  BenchmarkMethodPointer-8  |  2000000000  |  0.30 ns/op  |
|  BenchmarkFuncValueNoRef-8  |  2000000000  |  0.30 ns/op  |
|  BenchmarkFuncPointerNoRef-8  |  2000000000  |  0.30 ns/op  |
|  BenchmarkMethodValueNoRef-8  |  2000000000  |  0.30 ns/op  |
|  BenchmarkMethodPointerNoRef-8  |  2000000000  |  0.30 ns/op  |

Just to be sure, let's run the same tests with unexported structs and functions :

| Benchmarks on Unexported  | # of op. | dur. per op. |
| ------------- | ------------- | ------------- |
|  BenchmarkUnexportedNewValue-8  |  100000000  |  10.7 ns/op  |
|  BenchmarkUnexportedNewPointer-8  |  300000000  |  5.32 ns/op  |
|  BenchmarkUnexportedFuncValue-8  |  300000000  |  4.77 ns/op  |
|  BenchmarkUnexportedFuncPointer-8  |  2000000000  |  0.30 ns/op  |
|  BenchmarkUnexportedMethodValue-8  |  300000000  |  4.75 ns/op  |
|  BenchmarkUnexportedMethodPointer-8  |  2000000000  |  0.30 ns/op  |
|  BenchmarkUnexportedFuncValueNoRef-8  |  2000000000  |0.30 ns/op  |
|  BenchmarkUnexportedFuncPointerNoRef-8  |  2000000000  |0.30 ns/op  |
|  BenchmarkUnexportedMethodValueNoRef-8  |  2000000000  |0.30 ns/op  |
|  BenchmarkUnexportedMethodPointerNoRef-8  |  2000000000  |0.37 ns/op  |