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

Ответ:
```go
package main

import (
	"fmt"
	"os"
)

/*Интерфейс в golang представлен структурой:
type iface struct {
	tab  *itab
	data unsafe.Pointer
}
и содержит данные о типах (интерфейса и типа переменной внутри него) 
и указатель на, собственно, саму переменную со статическим (конкретным) типом (поле data в iface). 
В нашем случае в err содержится информация о типе переменной (*os.PathError). 
Таким образом iface.data != nil (true), соответственно и весь интерфейс не равен nil.*/



func Foo() error {
	var err *os.PathError = nil //fmt.Println(err == nil) true. Здесь err-это нулевое значение указателя
	return err
}
func main() {
	err := Foo()
	fmt.Println(err)            // <nil>
	fmt.Println(err == nil)     // false
}

```