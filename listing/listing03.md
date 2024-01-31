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
```
Вывод:
<nil>
false

При вызове функции foo() возвращается значение типа error. В поле tab находится указатель на тип os.PathError, а data будет nil.
Поэтому при выводе переменной err получаем значение nil. 

Но при сравнении err == nil мы получаем false, т.к. для того чтобы значение интерфейса равнялось nil, поля tab и data должны быть равны nil.
В нашем примере в поле tab находится указатель на тип os.PathError. 

```
