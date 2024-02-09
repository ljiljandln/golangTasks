package pattern

import "log"

/*
Фабричный метод
порождающий паттерн уровня класса, сфокусирован только на отношениях между классами
является основой для всех порождающих паттернов, поэтому может быть легко модифицирован
под другие порождающие паттерны в процессе

применяется когда:
- система должна легко расширяться при добавлении новых типов
- нет четких требований к системе продукта
- не ясен способ организации взаимодействия между продуктами

плюсы:
- избавляет от привязки к конкретным структурам
- код производства продуктов оказывается в одном месте, проще поддерживать
- проще добавлять элементы в программу
- реализует [[принцип открытости закрытости]]

минусы:
- может порождать большие иерархии структур, так как для каждой структуры
нужно будет создать свою структуру создателя
*/

type action string

type Creator interface {
	CreateProduct(a action) Product // фабричный метод
}

type Product interface {
	Use() string
}

type ConcreteCreator struct{}

func NewCreator() Creator {
	return &ConcreteCreator{}
}

func (p *ConcreteCreator) CreateProduct(action action) Product {
	var product Product
	switch action {
	case "Fly":
		product = &Plain{string(action)}
	case "Sail":
		product = &Boat{string(action)}
	default:
		log.Fatalln("Unknown Action")
	}
	return product
}

type Plain struct {
	action string
}

func (p *Plain) Use() string {
	return p.action
}

type Boat struct {
	action string
}

func (p *Boat) Use() string {
	return p.action
}
