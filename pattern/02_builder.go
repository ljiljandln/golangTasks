package pattern

/*
Паттерн Builder
порождающий паттерн проектирования уровня объекта, определяет этапы построения объекта

применение:
- когда нужно создавать сложные составные объекты или разные сборки объектов

плюсы:
- пошаговая сборка
- переиспользование кода для других продуктов
- изоляция кода-сборки от бизнес-логики

минусы:
- усложняет код
- привязывает к конкретному билдеру
*/

const (
	SamsungBuilder = "samsung"
	IphoneBuilder  = "iphone"
)

type Builder interface {
	SetBrand()
	SetModel()
	SetPrice()
	GetPhone() Phone
}

func getBuilder(builderType string) Builder {
	switch builderType {
	default:
		return nil
	case SamsungBuilder:
		return &SBuilder{}
	case IphoneBuilder:
		return &IBuilder{}
	}
}

type Phone struct {
	Brand, Model string
	Price        float64
}

type SBuilder struct {
	Brand, Model string
	Price        float64
}

func (b SBuilder) SetBrand() {
	b.Brand = "Samsung"
}

func (b SBuilder) SetModel() {
	b.Model = "Galaxy"
}

func (b SBuilder) SetPrice() {
	b.Price = 1999.99
}

func (b SBuilder) GetPhone() Phone {
	return Phone{
		Brand: b.Brand,
		Model: b.Model,
		Price: b.Price,
	}
}

type IBuilder struct {
	Brand, Model string
	Price        float64
}

func (b IBuilder) SetBrand() {
	b.Brand = "Iphone"
}

func (b IBuilder) SetModel() {
	b.Model = "14 s"
}

func (b IBuilder) SetPrice() {
	b.Price = 2999.99
}

func (b IBuilder) GetPhone() Phone {
	return Phone{
		Brand: b.Brand,
		Model: b.Model,
		Price: b.Price,
	}
}
