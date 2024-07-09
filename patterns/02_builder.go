package patterns

/*
 Строитель - порождающий паттерн проектирования, диктующий универсальные правила создания нового сложного объекта.
Каждый шаг данной абстракции осуществляет отдельный builder.

Составляющие паттерна:
- построенный сложный объект
- builder, который предоставляет интерфейс для создания частей итогового объекта
- director, управляющий builder'ами и регулирующий процессом создания объекта
*/

type Lorry struct {
	color      string
	engineType string
	maxSpeed   int
	capacity   int
}

type LorryBuilder interface {
	SetColor(color string) LorryBuilder
	SetEngineType(engine string) LorryBuilder
	SetMaxSpeed(speed int) LorryBuilder
	SetCapacity(capacity int) LorryBuilder
	Build() *Lorry
}

func NewLorryBuilder() lorryBuilder {
	return lorryBuilder{lorry: &Lorry{}}
}

type lorryBuilder struct {
	lorry *Lorry
}

func (lb *lorryBuilder) SetColor(color string) {
	lb.lorry.color = color
}

func (lb *lorryBuilder) SetEngineType(engine string) {
	lb.lorry.engineType = engine
}

func (lb *lorryBuilder) SetMaxSpeed(speed int) {
	lb.lorry.maxSpeed = speed
}

func (lb *lorryBuilder) SetCapacity(capacity int) {
	lb.lorry.capacity = capacity
}

func (lb *lorryBuilder) Build() *Lorry {
	return lb.lorry
}

type Director struct {
	builder LorryBuilder
}

func (d *Director) ConstructLorry(color, engine string, speed, capacity int) *Lorry {
	d.builder.SetColor(color).SetEngineType(engine).SetMaxSpeed(speed).SetCapacity(capacity)
	return d.builder.Build()
}
