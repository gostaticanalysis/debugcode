# debugcode

[![pkg.go.dev][gopkg-badge]][gopkg]

`debugcode` finds debug codes.

* [builtinprint](https://github.com/gostaticanalysis/builtinprint): finds calling builtin `print` or `println`.
* [commentout](./passes/commentout): finds a commented out debug code without reason

## Install

You can get `debugcode` by `go install` command (Go 1.16 and higher).

```bash
$ go install github.com/gostaticanalysis/debugcode@latest
```

## How to use

`debugcode` run with `go vet` as below when Go is 1.12 and higher.

```bash
$ go vet -vettool=$(which debugcode) ./...
```

## Analyzers

### builtinprint

See: https://github.com/gostaticanalysis/builtinprint

### commentout

[commentout](./passes/commentout): finds a commented out debug code without reason.

```go
package a

//func f() { // want "do not leave a commented out debug code without reason"
//	panic("not implement")
//}

// with reason // OK
//func f() {
//	panic("not implement")
//}
```

## Analyze with golang.org/x/tools/go/analysis

You can get analyzers of debugcode from [debugcode.Analyzers](https://pkg.go.dev/github.com/gostaticanalysis/debugcode/debugcode/#Analyzers).
And you can use them with [unitchecker](https://golang.org/x/tools/go/analysis/unitchecker).

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gostaticanalysis/debugcode
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gostaticanalysis/debugcode?status.svg
