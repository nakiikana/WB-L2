package pattern

/*
 Строитель - порождающий паттерн проектирования, диктующий универсальные правила создания нового сложного объекта.
Каждый шаг данной абстракции осуществляет отдельный builder.

Составляющие паттерна:
- построенный сложный объект
- builder, который предоставляет интерфейс для создания частей итогового объекта
- director, управляющий builder'ами и регулирующий процессом создания объекта

Процесс создания объекта:
- создается структура самого сложного объекта
- создается интерфейс со всеми необходимыми методами для создания объекта
- создается Director c методом для конструирования объекта с собственной конфигурацией.
*/

// Lorry - тот сложный объект? который мы будем конструировать.
type Lorry struct {
	color      string
	engineType string
	maxSpeed   int
	capacity   int
}

// LorryBuilder - тот самый интерфейс, который нам нужно реализовать для конструирования частей итогового объекта
type LorryBuilder interface {
	SetColor(color string) LorryBuilder
	SetEngineType(engine string) LorryBuilder
	SetMaxSpeed(speed int) LorryBuilder
	SetCapacity(capacity int) LorryBuilder
	Build() *Lorry
}

// Создает новый LorryBuilder
func NewLorryBuilder() LorryBuilder {
	//Так как receiver у нас именно указатель, к нему мы и обращаемся
	return &lorryBuilder{lorry: &Lorry{}}
}

// Удовлетворяет интерфейсу LorryBuilder:
type lorryBuilder struct {
	lorry *Lorry
}

// все необходимые методы для этого:
func (lb *lorryBuilder) SetColor(color string) LorryBuilder {
	lb.lorry.color = color
	return lb
}

func (lb *lorryBuilder) SetEngineType(engine string) LorryBuilder {
	lb.lorry.engineType = engine
	return lb
}

func (lb *lorryBuilder) SetMaxSpeed(speed int) LorryBuilder {
	lb.lorry.maxSpeed = speed
	return lb
}

func (lb *lorryBuilder) SetCapacity(capacity int) LorryBuilder {
	lb.lorry.capacity = capacity
	return lb
}

// Возвращаем существующий объект.
func (lb *lorryBuilder) Build() *Lorry {
	return lb.lorry
}

// Director предоставляет интерфейс, с помощью которого новый объект будет собираться.
type Director struct {
	builder LorryBuilder
}

func (d *Director) ConstructLorry(color, engine string, speed, capacity int) *Lorry {
	//Цепочка вызовов методов для установки свойств объекта.
	d.builder.SetColor(color).SetEngineType(engine).SetMaxSpeed(speed).SetCapacity(capacity)
	return d.builder.Build()
}
