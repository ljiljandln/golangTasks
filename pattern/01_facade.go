package pattern

import (
	"errors"
	"fmt"
)

/*
Структурный паттерн проектирования, позволяющий скрыть сложную реализацию системы
с помощью простого объекта, с которым будет взаимодействовать с пользователем

Плюсы - прост в реализации
Минусы - может стать "божественным объектом"
*/

type Person struct {
	name string
	card Card
	car  Car
}

func (p Person) fillTheCar(l uint, petrol Petrol) error {
	price := petrol.getPrice(l)
	err := p.card.buySomething(price)
	if err != nil {
		return err
	}
	p.car.addPetrol(l)
	fmt.Printf("%s заправил машину на %d литров, с карты списано %f рублей\n", p.name, l, price)
	return nil
}

type Card struct {
	balance float64
}

func (c Card) buySomething(price float64) error {
	if price > c.balance {
		return errors.New("not enough money")
	}
	c.balance -= price
	return nil
}

type Petrol struct {
	year  uint
	price float64
}

func (p Petrol) getPrice(l uint) float64 {
	return float64(l) * p.price
}

type Car struct {
	model  string
	petrol uint
}

func (c Car) addPetrol(l uint) {
	c.petrol += l
}
