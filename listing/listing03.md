Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.
```go
package main

import (
    "fmt"
    "os"
)

func Foo() error {
    var err *os.PathError = nil
    return err
}

func main() {
    err := Foo()
    fmt.Println(err)
    fmt.Println(err == nil)
}
```

программа выведет:  
`<nil>`  
`false`

функция Foo() возвращает ошибку *os.PathError равную nil, поэтому в первой строке выводится nil

во второй строке выводится результат сравнения err (*os.PathError) и nil, и они не равны, так как err - это указатель типа *os.PathErro, в отличии от просто nil


Интерфейс в Go является контрактом, и каждый тип, который реализовывает методы,
которые есть в интерфейсе – является реализацией интерфейса
Задается структурой iface
```go
type iface struct {  
    tab  *itab  
    data unsafe.Pointer  
}
```

Пустой же интерфейс не содержит методов, и согласно этому – под пустой интерфейс
подходит любой тип.
задается структурой eface
```go
type eface struct {
	_type *_type
	data  unsafe.Pointer
}
```