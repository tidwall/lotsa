# lotsaa

`Lotsaa` = `Lotsa` + `a` fixed duration feature

This package was created by forking Josh Baker's [Lotsa](https://github.com/tidwall/lotsa) framework.

To run the operations spread over 4 threads for a fixed duration, use `lotsa.Time`

```go
var total int64
lotsa.Output = os.Stdout
lotsa.Time(23 * time.Millisecond, 4,
    func(_ *rand.Rand, thread int) {
        atomic.AddInt64(&total, 1)
    },
)
```

Prints:

```
654,330 ops over 4 threads in 24ms, 27,207,775/sec, 36 ns/op
```

# lotsa

Lotsa is a simple Go library for executing lots of operations spread over any number of threads.

## Install

```
go get -u github.com/tidwall/lotsa
```

## Example

Here we load 1,000,000 operations spread over 4 threads.

```go
var total int64
lotsa.Ops(1000000, 4,
    func(i, thread int) {
        atomic.AddInt64(&total, 1)
    },
)
println(total)
```

Prints `1000000`

To output some benchmarking results, set the `lotsa.Output` prior to calling `lotsa.Ops`

```go
var total int64
lotsa.Output = os.Stdout
lotsa.Ops(1000000, 4,
    func(i, thread int) {
        atomic.AddInt64(&total, 1)
    },
)
```

Prints: 

```
1,000,000 ops over 4 threads in 23ms, 43,580,037/sec, 22 ns/op
```

## Contact

Josh Baker [@tidwall](http://twitter.com/tidwall)

## License

Source code is available under the MIT [License](/LICENSE).
