package pattern

/*
Цепочка вызовов
поведенческий паттерн проектирования уровня объекта, позволяет передавать
запросы последовательно по цепочке обработчиков

применяется когда:
- неизвестно какой запрос придет и какой обработчик понадобится
- важно чтобы обработчики выполнялись в строгом порядке

Плюсы:
- уменьшает зависимость между клиентом и обработчиком
- реализует принцип единственной обязанности
- реализует принцип открытости закрытости

Минусы:
- запрос может остаться без обработки
*/

type Handler interface {
	SendRequest(msg int) string
}

type FirstHandler struct {
	next Handler
}

func (h *FirstHandler) SendRequest(msg int) (res string) {
	if msg == 1 {
		res = "FirstHandler response"
	} else if h.next != nil {
		res = h.next.SendRequest(msg)
	}
	return
}

type SecondHandler struct {
	next Handler
}

func (h *SecondHandler) SendRequest(msg int) (res string) {
	if msg == 2 {
		res = "SecondHandler response"
	} else if h.next != nil {
		res = h.next.SendRequest(msg)
	}
	return
}

type ThirdHandler struct {
	next Handler
}

func (h *ThirdHandler) SendRequest(msg int) (res string) {
	if msg == 3 {
		res = "ThirdHandler response"
	} else if h.next != nil {
		res = h.next.SendRequest(msg)
	}
	return
}
