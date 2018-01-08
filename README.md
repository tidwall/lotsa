# lotsa

Lotsa is a simple Go library for executing lots of operations spread over any number of threads.

## Install

```
go get -u github.com/tidwall/lotsa
```

## Example

```
var total int64
lotsa.Ops(1000000, 4, os.Stdout, 
    func(i, thread int) {
        atomic.AddInt64(&total, 1)
    },
)
println(total)
```

Prints `1000000`

## Contact

Josh Baker [@tidwall](http://twitter.com/tidwall)

## License

Source code is available under the MIT [License](/LICENSE).