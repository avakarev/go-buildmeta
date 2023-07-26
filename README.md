# buildmeta

> Tiny go module to hold build and runtime information

# NOTE: deprecated, merged into https://github.com/avakarev/go-util

> Scheduled for deletion after 2023-09-01


## Install

```shell
go get github.com/avakarev/go-buildmeta
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/avakarev/go-buildmeta"
)

func main() {
	fmt.Println(buildmeta.Compiler()) // => go1.18.3
}
```


## License

`go-buildmeta` is licensed under MIT license. (see [LICENSE](./LICENSE))
