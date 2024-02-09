package pattern

/*
Паттерн Состояние

поведенческий паттерн уровня объекта, позволяющий объекты менять свое
поведение в зависимости от внутреннего состояния

применяется когда:
- поведение объекта зависит от его состояния
- поведение объекта должно изменяться во время выполнения программы
- состояний достаточно много и использовать для этого условные операторы,
разбросанные по коду, достаточно затруднительно

плюсы:
- избавляет от множества условных операторов машины состояний
- концентрирует в одном месте код, связанный с определённым состоянием
- упрощает код контекста

минусы:
- может усложнить код если состояний мало
*/

// MobileActionStater общий интерфейс для разных состояний
type MobileActionStater interface {
	Sound() string
}

// MobileSound реализует поведение в зависимости от состояния
type MobileSound struct {
	state MobileActionStater
}

// Sound возвращает звук
func (a *MobileSound) Sound() string {
	return a.state.Sound()
}

// SetState меняет состояние телефона
func (a *MobileSound) SetState(state MobileActionStater) {
	a.state = state
}

// NewMobileSound конструктор MobileSound
func NewMobileSound() *MobileSound {
	return &MobileSound{state: &MobileAlarmClock{}}
}

// MobileAlarmClock реализует поведение будильника
type MobileAlarmClock struct {
}

func (a *MobileAlarmClock) Sound() string {
	return "Ауууу, на работу! как на праздник! шагом марш на работу! буги вуги!"
}

// MobileMessage реализует поведение во время приема сообщения
type MobileMessage struct {
}

func (a *MobileMessage) Sound() string {
	return "бдддддзззззз"
}

// MobileCall реализует поведение во время звонка
type MobileCall struct {
}

func (a *MobileCall) Sound() string {
	return "Почему так в России березы шумят, почему белоствольные все понимают"
}
