package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern

Паттерн Состояние позволяет конструировать объект, способный иметь набор дискретных состояний и ведущий себя по-разному в
зависимости от состояния, в котором находится.

Паттерн Состояние используется, когда:
- нужно менять поведение объекта в зависимости от его состояния;
- нужно реализовать конечный автомат, основанный на таблице переходов.

Плюсы:
- Избавление от больших условных конструкций (или свитчей)
- Весь код, от относящийся к одному состоянию объекта находится в одном месте -> упрощение читабельности кода

Минусы:
- Может неоправданно усложнить код, если состояний мало и они редко меняются.

Примеры:
1) Состояние Документа при его модификации:в состоянии черновика, под редакцией или уже опубликован. В каждом из этих состояний должны
быть осуществлены свои действия
2) Состояние банкомата в зависимости от наличия в нем средств или же от отдельных факторов: к примеру, если средства, есть состояние выдачи, если нет ->
средства выдать невозможно, но, к примеру, можно произвести перевод средств и т.д.

Пример: реализация структуры Human, изменение поведения в зависимости от разных состояний (sick, healthy)
*/

type Condition interface {
	Work()
	Sleep()
}

type Human struct {
	fio string
	Condition
}

type healthyCondition struct{}

type sickCondition struct{}

func (t healthyCondition) Work() {
	println("works hard")
}

func (t healthyCondition) Sleep() {
	println("doesn't want to sleep")
}

func (r sickCondition) Work() {
	println("doesn't want to work")
}

func (r sickCondition) Sleep() {
	println("sleeps a lot")
}

func (h Human) SwitchState(c Condition) {
	h.Condition = c
}

func mainState() {
	healthy := &healthyCondition{}
	human := &Human{
		fio:       "FIO",
		Condition: healthy,
	}
	human.Sleep()
	human.Work()
	sick := &sickCondition{}
	human.SwitchState(sick)
	human.Sleep()
	human.Work()
}
