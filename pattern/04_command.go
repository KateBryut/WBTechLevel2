package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern

Команда (Command) — преобразует запросы в объекты, что позволяет хранить их и делать отмену операций;
Паттерн Команда преобразует все параметры операции или события в объект-команду. Впоследствии можно выполнить эту операцию,
вызвав соответствующий метод объекта. Объект-команда заключает в себе всё необходимое для проведения операции, поэтому её легко выполнять,
логировать и отменять.

Паттерн Команда применяется, когда:
- нужно преобразовать операции в объекты, которые можно обрабатывать и хранить: использование объектов вместо операций позволяет
создавать очереди, передавать команды дальше или выполнять их в нужный момент;
- требуется реализовать операцию отмены выполненных действий.

Плюсы:
- Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
- Позволяет реализовать простую отмену и повтор операций.
- Позволяет реализовать отложенный запуск операций.
- Позволяет собирать сложные команды из простых.
- Реализует принцип открытости/закрытости.
Минусы:
- Усложняет код программы из-за введения множества дополнительных классов.

Примеры:
- Шаблон Команда применяется в работе с базами данных. В стандартной библиотеке Go есть пример SQL-инструкции sql.Stmt.
Такую заранее подготовленную инструкцию можно многократно выполнять методом Stmt.Exec, не задумываясь о её внутренней структуре.
sql.Stmt, выполненную в рамках транзакции Tx.Stmt(), легко откатить с помощью Tx.Rollback().
- При редактировании в фотошопе для реализации команд отмены действий.

Мой пример:
Есть телевизор, его можно включить при помощи пульта управления.
*/

import "fmt"

// recieverTV - телевизор
type receiverTV interface {
	action()
	turnOn()
	turnOff()
}

// метод включения телевизора
func (r *rcvrTV) turnOn() {
	fmt.Println("Телевизор включен")
}

// метод выключения телевизор
func (r *rcvrTV) turnOff() {
	fmt.Println("Телевизор выключен")
}

// интерфейс команды
type command interface {
	execute()
}

// конкретная реализация command для включения телевизора
type onTVCommand struct {
	receiver receiverTV
}

func (c *onTVCommand) execute() {
	c.receiver.turnOn()
}

// конкретная реализация command для выключения телевизора
type offTVCommand struct {
	receiver receiverTV
}

func (c *offTVCommand) execute() {
	c.receiver.turnOff()
}

// invoker - пульт
type invoker struct {
	commands map[string]command
}

func newInvoker() *invoker {
	i := new(invoker)
	i.commands = make(map[string]command)
	return i
}

func (i *invoker) do(c string) {
	i.commands[c].execute()
}

// реализация receiver
type rcvrTV struct {
	name string
}

func (r *rcvrTV) action() {
	fmt.Println(r.name)
}

func mainCommand() {
	tv := rcvrTV{"TV"}
	var cmd1 = onTVCommand{&tv}
	var cmd2 = offTVCommand{&tv}

	remoteController := newInvoker()
	remoteController.commands["on_TV"] = &cmd1
	remoteController.commands["off_TV"] = &cmd2
	// применение команд
	remoteController.do("on_TV")
	remoteController.do("off_TV")
}
