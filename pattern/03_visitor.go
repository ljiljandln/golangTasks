package pattern

import "fmt"

/*
Паттерн Visitor
поведенческий паттерн проектирования, позволяющий создавать новые операции,
не меняя структуры объектов

применяется когда нужно применить одну и ту же операцию к объектам разных структур

плюсы:
- упрощает добавление операций
- объединяет родственные операции в одной структуре
- может накапливать состояние

минусы:
- применение не оправдано если иерархия компонентов часто меняется
- могут начаться баги инкапсуляции компонентов
*/

type Visitor interface {
	VisitClub(c Club) string
	VisitCafe(c Cafe) string
}

type Human struct {
	name string
}

func (h Human) VisitClub(c Club) string {
	return fmt.Sprintf("%s visit club %s\n", h.name, c.name)
}

func (h Human) VisitCafe(c Cafe) string {
	return fmt.Sprintf("%s visit cafe %s\n", h.name, c.name)
}

type Cafe struct {
	name string
}

func (c Cafe) Accept(v Visitor) {
	fmt.Println(v.VisitCafe(c))
}

type Club struct {
	name string
}

func (c Club) Accept(v Visitor) {
	fmt.Println(v.VisitClub(c))
}

type Acceptor interface {
	Accept(v Visitor)
}
