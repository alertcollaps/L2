Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)
func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) *
				time.Millisecond)
		}
		close(c)
	}()
	return c
}
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	k := 0
	m:=0
	go func() {
		for {
			select {
			case v,ok:= <-a:
				if ok{c <- v}else{k+=1}
			case v,ok := <-b:
				if ok{c <- v}else{m+=1}
			}
			if k>=1&&m>=1{
				close(c) // при отсутствии данных из обоих каналов - закрываем общий! Если не закроем, будем бесконечно выводить в stdout значение int 
				break}   // по умолчанию, т.е. 0
		}
	}()
	return c
}
func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("The end!!!")
}

```