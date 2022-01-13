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

Ответ:
```go
package main
import (
"fmt"
)
// Т.к. отложенные функции могут считывать и присваивать именованные возвращаемые значения
// возвращаемой функции, функция test() увеличит х на единицу и вернет увеличенное значение, а
// функция anotherTest() увеличит х на единицу, но вернет значение х до вхождения в отложенную функцию
func test() (x int) {
	defer func() {
		x++
	}()
	x=1
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
	fmt.Println(test())         // 2

	fmt.Println(anotherTest())  // 1

}

```