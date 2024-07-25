package pattern

import "fmt"

/*
 Фасад - структурный паттерн проектирования, который предоставляет унифицированный интерфейс для взаимодействия
с подсистемами - все возможные вызовы сводятся к одному объекту, делегирующему их соответствующим объектам
системы.
*/

type ComputerFacade interface {
	TurnOn()
	TurnOff()
}

func NewComputerFacade() ComputerFacade {
	return NewComputer()
}

// Через методы данной структуры будет осуществляться доступ к остальным подсистемам (фасад).
type Computer struct {
	screen   Screen
	keyboard Keyboard
	mouse    Mouse
}

// Функция, создающая новый объект компьютера.
func NewComputer() *Computer {
	return &Computer{
		screen:   Screen{},
		keyboard: Keyboard{},
		mouse:    Mouse{},
	}
}

func (c *Computer) TurnOn() {
	c.keyboard.ConnectKeyboard()
	c.screen.ShowLoadingScreen()
	c.mouse.ConnectMouse()
}

func (c *Computer) TurnOff() {
	c.keyboard.DisconnectKeyboard()
	c.screen.CloseEverything()
	c.mouse.DisconnectMouse()
}

// Подсистема "Экран" со своим методами включения и выключения экрана.
type Screen struct {
}

func (s *Screen) ShowLoadingScreen() {
	fmt.Println("Your screen is loading")
}

func (s *Screen) CloseEverything() {
	fmt.Println("Closing all your tabs and turning off the screen")
}

// Подсистема "Клавиатура" с методами подключения к компьютеру и прекращения работы.
type Keyboard struct {
}

func (k *Keyboard) ConnectKeyboard() {
	fmt.Println("You can use your keyboard now")
}

func (k *Keyboard) DisconnectKeyboard() {
	fmt.Println("The keyboard is disconnected now")
}

// Подсистема "Мышка" со своими методами.
type Mouse struct {
}

func (m *Mouse) ConnectMouse() {
	fmt.Println("The mouse is connected to your computer")
}

func (k *Mouse) DisconnectMouse() {
	fmt.Println("The mouse is disconnected now")
}
