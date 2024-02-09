Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
    "fmt"
)

func test() (x int) {
    defer func() {
        x++
    }()
    x = 1
    return
}

func anotherTest() int {
    var x int
    defer func() {
        x++
    }()
    x = 1
    return x
}

func main() {
    fmt.Println(test())
    fmt.Println(anotherTest())
}
```
в 1ой строке программа выведет `2`, так как анонимная функция выполнится после return и x перезапишется  

в 2ой строке программа выведет `1`, так как анонимная функция была выполнена уже после того, как anotherTest() вернула значение `x`