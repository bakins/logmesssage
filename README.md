# logmessage

Ensure that [zap](https://godoc.org/go.uber.org/zap) log messages are not capitalized.

This tool analyzes Go code, finds calls to zap Logger(s), and reports any messages that begin
with a capital letter.

In normal usage, zap errors are JSON and are parsed by machines, so complete sentences are not needed.
See also https://github.com/golang/go/wiki/CodeReviewComments#error-strings

## Install

```bash
go get github.com/bakins/logmessage
```

## Usage

By default, `logmessage` prints the output of the analyzer to stdout. You can pass
a file, directory or a Go package:

```sh
$ logmessage foo.go # pass a file
$ logmessage ./...  # recursively analyze all files
$ logmessage github.com/fatih/gomodifytags # or pass a package
```

When called it displays the error with the line and column:

```
gomodifytags@v1.0.1/main.go:200:16: log messages should not be capitalized
gomodifytags@v1.0.1/main.go:641:17: call could wrap the error with error-wrapping directive %w
```

## Acknowledgements

* Thanks to [Fatih Arslan](https://github.com/fatih) for https://arslan.io/2019/06/13/using-go-analysis-to-write-a-custom-linter/
* Portions of this linter are based on https://github.com/golang/lint/blob/master/lint.go

## LICENSE

See [LICENSE](./LICENSE)

