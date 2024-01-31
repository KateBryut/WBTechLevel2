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
Вывод: 
error

В переменной err хранится информация о типе (в tab), поэтому err!=nil.
Для функций, возвращающих ошибки, рекомендуется всегда использовать тип error.

```
