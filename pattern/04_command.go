package pattern

import "fmt"

/*
Паттерн команда
поведенческий паттерн уровня объекта, позволяет представить запрос
в виде объекта. Из этого следует, что команда - это объект. Такие запросы,
например, можно ставить в очередь, отменять или возобновлять.

применяется когда нужно:
- составить очередь из запросов
- выполнять запросы по расписанию
- отменять запросы
- передавать по сети

плюсы:
- нет прямой зависимости между объектами
- легко добавлять/удалять запросы
- легко выполнить отложенный запуск запросов
- реализует [[принцип открытости закрытости]]

минусы:
- усложняет код
*/

type Command interface {
	Execute()
}

type SwitchOnCMD struct {
	receiver *Receiver
}

func (c *SwitchOnCMD) Execute() {
	fmt.Println(c.receiver.SwitchOn())
}

type SwitchOffCMD struct {
	receiver *Receiver
}

func (c *SwitchOffCMD) Execute() {
	fmt.Println(c.receiver.SwitchOff())
}

type Receiver struct {
	name string
}

func (r *Receiver) SwitchOn() string {
	return fmt.Sprintf("%s switchOn", r.name)
}

func (r *Receiver) SwitchOff() string {
	return fmt.Sprintf("%s switchOff", r.name)
}

type Invoker struct { //вызыватель
	commands []Command
}

func (i *Invoker) StoreAndExecute(cmd Command) {
	i.commands = append(i.commands, cmd)
	cmd.Execute()
}
