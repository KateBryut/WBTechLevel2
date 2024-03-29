package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern

Посетитель (Visitor) — даёт возможность добавлять новый функционал к объектам, не внося в них изменения;

Шаблон Посетитель позволяет отвязать функциональность от объекта. Новые методы добавляются не для каждого типа из семейства,
а для промежуточного объекта visitor, аккумулирующего функциональность. Типам семейства добавляется только один метод accept(visitor).
Так проще добавлять операции к существующей базе кода без особых изменений и страха всё сломать. Этот паттерн чаще всего используется,
когда нужно добавить функционал к объектам разного типа.

Паттерн Посетитель используется, когда:
- нужно применить одну и ту же операцию к объектам разных типов;
- часто добавляются новые операции для объектов;
- требуется добавить новый функционал, но избежать усложнения кода объекта.

Плюсы:
- Упрощает добавление операций, работающих со сложными структурами объектов.
- Объединяет родственные операции в одном классе.
- Посетитель может накапливать состояние при обходе структуры элементов.
Минусы:
- Паттерн не оправдан, если иерархия элементов часто меняется.
- Может привести к нарушению инкапсуляции элементов.

Примеры:
1) Готова реализация какой-то структуры и нужно добавить новую логику, не ломая старую.
Есть структура Заказ, нам необходимо добавить возможность подсчета суммы одного заказа и его запись в файл,
При добавлении данного метода в интерфейс, который имплементируют Заказ, обратная совместимость может быть сломана.

Мой пример:
Предположим, есть конструктор, который собирает автомобиль из колёс и двигателя. В какой-то момент нужно добавить
тестирование компонентов при сборке.

*/

// CarPart — семейство типов, которым хотим добавить
// функциональность детали автомобиля.
type CarPart interface {
	Accept(CarPartVisitor)
}

// CarPartVisitor — интерфейс visitor,
// в его коде и содержится новая функциональность.
type CarPartVisitor interface {
	testWheel(wheel *Wheel)
	testEngine(engine *Engine)
}

// Wheel — реализация деталей.
type Wheel struct {
	Name string
}

// Accept — единственный метод, который нужно добавить типам семейства,
// ссылка на метод visitor.
func (w *Wheel) Accept(visitor CarPartVisitor) {
	visitor.testWheel(w)
}

type Engine struct{}

func (e *Engine) Accept(visitor CarPartVisitor) {
	visitor.testEngine(e)
}

type Car struct {
	parts []CarPart
}

// NewCar — конструктор автомобиля.
func NewCar() *Car {
	this := new(Car)
	this.parts = []CarPart{
		&Wheel{"front left"},
		&Wheel{"front right"},
		&Wheel{"rear right"},
		&Wheel{"rear left"},
		&Engine{}}
	return this
}

func (c *Car) Accept(visitor CarPartVisitor) {
	for _, part := range c.parts {
		part.Accept(visitor)
	}
}

// TestVisitor — конкретная реализация visitor,
// которая может проверять колёса и двигатель.
type TestVisitor struct {
}

func (v *TestVisitor) testWheel(wheel *Wheel) {
	fmt.Printf("Testing the %v wheel\n", wheel.Name)
}

func (v *TestVisitor) testEngine(engine *Engine) {
	fmt.Println("Testing engine")
}

func mainVisitor() {
	// клиентский код
	car := NewCar()
	visitor := new(TestVisitor)
	car.Accept(visitor)
}
