package pattern

import(
	"fmt"
)

type Color interface {
	Transition() Color
}

 
type RedColor struct{}

func (s *RedColor) Transition() Color {
	fmt.Println("Переключение на желтый свет")
	return &YellowColor{}
}

type YellowColor struct{}

func (s *YellowColor) Transition() Color {
	fmt.Println("Переключение на зеленый свет")
	return &GreenColor{}
}

 
type GreenColor struct{}

func (s *GreenColor) Transition() Color {
	fmt.Println("Переключение на красный свет")
	return &RedColor{}
}

 
type TrafficLight struct {
	Color Color
}

func NewTrafficLight() *TrafficLight {
	return &TrafficLight{Color: &RedColor{}}
}

func (tl *TrafficLight) ChangeColor() {
	tl.Color = tl.Color.Transition()
}


/* пример вывода 
func main() {
	trafficLight := NewTrafficLight()
	for i := 0; i < 3; i++ {
		trafficLight.ChangeColor()
	}
}

*/ 

/*
	Паттерн "Состояние" (Color) является поведенческим паттерном проектирования и используется для управления состоянием объекта и переходами между состояниями. Он позволяет объекту изменять свое поведение при изменении его внутреннего состояния, не меняя структуру самого объекта.


*/
