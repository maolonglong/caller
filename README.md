# Caller

Get caller info with LRU cache.

## Install

```bash
go get go.chensl.me/caller
```

## Usage

```go
package main

import (
	"fmt"

	"go.chensl.me/caller"
)

func main() {
	_, file, line, _ := caller.Get(1)
	fmt.Printf("%s:%d\n", file, line)
}
```

## Benchmark

```text
BenchmarkRuntimeCaller-4         1066604              1317  ns/op
BenchmarkGet-4                   2754568              415.5 ns/op
```
