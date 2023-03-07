# typeconv

Easy and safe type converting from one pointer to value & value to pointer in Go

Jump To:

- [Getting Started](#Getting-Started)
- [Quick Examples](#Quick-Examples)
- [Getting Help](#Getting-Help)
- [Contributing](#Contributing)
- [More Resources](#Resources)

## Getting Started

### Installing

Use `go get` to retrieve the pkg to add it to your `GOPATH` workspace, or
project's Go module dependencies.

```sh
go get github.com/mdsohelmia/typeconv
```

To update the pkg use `go get -u` to retrieve the latest version of the pkg.

```sh
   go get -u github.com/mdsohelmia/typeconv
```

## Quick Examples

```go
 package main

import (
	"fmt"

	"github.com/mdsohelmia/typeconv"
)

func main() {
	boolPtr := typeconv.Bool(true)
	fmt.Println(boolPtr)
	boolValue := typeconv.BoolValue(boolPtr)
	fmt.Println(boolValue)
}
```
