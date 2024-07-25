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
```
выводом будет error , потому что мы создаем переменную error, после чего вовзвращаеми указатель на структуру которая имлементирует интерфейс error, который уже равен nil, в итоге типом будет являться структура а не nil, поэтому сравнение != nil будет положительным и будет выводиться error.

```
