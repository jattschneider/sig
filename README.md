# sig


# Getting Started

Just a quick example how to use the sig library:

#### main.go
```
package main

import (
	"github.com/jattschneider/sig"
)

func main() {
    sig.Observe()
    // Some process
    // Eg: http.ListenAndServe(...)
}
```
