# signature

[![pkg.go.dev][gopkg-badge]][gopkg]

`signature` finds low readability functions.

```go
package a

// func decls
func _(_, _, _, _, _ int)    {}         // OK
func _(_, _, _, _, _, _ int) {}         // want "too many params"
func _() (_, _, _ int)       { return } // OK
func _() (_, _, _, _ int)    { return } // want "too many results"

// func literals
func _() {
	_ = func(_, _, _, _, _ int) {}         // OK
	_ = func(_, _, _, _, _, _ int) {}      // want "too many params"
	_ = func() (_, _, _ int) { return }    // OK
	_ = func() (_, _, _, _ int) { return } // want "too many results"
}
```

## Install

You can get `signature` by `go install` command (Go 1.16 and higher).

```bash
$ go install github.com/gostaticanalysis/signature/cmd/signature@latest
```

## How to use

`signature` run with `go vet` as below when Go is 1.12 and higher.

```bash
$ go vet -vettool=$(which signature) ./...
```

`signature` accepts two options `-params` and `-results`. `-params` specifies max number of parameters and `-results` specifies max number of results.
The default value of `-params` is 5 and `-results` is 3.

```bash
$ go vet -vettool=$(which signature) -signature.params 3 -signature.results 4 ./...
```

## Analyze with golang.org/x/tools/go/analysis

You can use [signature.Analyzer](https://pkg.go.dev/github.com/gostaticanalysis/signature/#Analyzer) with [unitchecker](https://golang.org/x/tools/go/analysis/unitchecker).

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gostaticanalysis/signature
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gostaticanalysis/signature?status.svg
