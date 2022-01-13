Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```go
//Структура customError реализует метод Error(), а значит реализует интерфейс error.
package main

type customError struct {
	msg string
}
func (e *customError) Error() string {
	return e.msg
}
func test() *customError {
	{
		// do something
	}
	return nil
}
func main() {
	var err error // err сейчас является нулевым интерфейсом err==nil (true)
	err = test() // здесь err содержит информацию о типе переменной.
	// Тип переменной - это  *customError. Соответственно err уже не будет nil.
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

```