# Unwrap() for Go

# Usage:
```
package main

import (
	. "app/unwrap"
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("Result = %v\n", Unwrap(strconv.Atoi("2"))*Unwrap(strconv.Atoi("3")))
	fmt.Printf("Result = %v\n", strconv.Itoa(Unwrap(strconv.Atoi("42"))))
}```



