# slicefunk

Providing some collection functions in go

Example usage:
```
package main

import (
    "fmt"

    sf "github.com/sa-/slicefunk"
)

func main() {
    original := []int{1, 2, 3, 4, 5}
    newArray := sf.Map(original, func(item int) int { return item + 1 })
    newArray = sf.Map(newArray, func(item int) int { return item * 3 })
    newArray = sf.Filter(newArray, func(item int) bool { return item%2 == 0 })
    fmt.Println(newArray)
}
```