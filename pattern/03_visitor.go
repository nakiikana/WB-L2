package pattern

/* Посетитель - поведенческий паттерн проектирования: который позволяет обойти набор объектов с разными интерфейсами, а также добавить
новый метод без изменения исходного класса.

Составляющие паттерна:
- итерфейс с отдельным методом для посещения всех необходимых объектов
- интерфейс элемента, определяющий метод Accept, который принимает посетителя и делегирует ему выполнение работы - объект "приходит"
к посетителю, выполняющему над ним определенную операцию.
- конкретные элементы (granny, neighbour, friend)
- конкретные посетители
*/

type Visitor interface {
	VisitGranny(granny *Granny)
	VisitNeighbour(neighbour *Neighbour)
	VisitFriend(friend *Friend)
}

type ClosePeople interface {
	Accept(v Visitor)
}

type Granny struct {
}

// Имплементация Accept для бубушки
func (g *Granny) Accept(v Visitor) {
	v.VisitGranny(g)
}

type Neighbour struct {
}

// Имплементация Accept для соседа
func (n *Neighbour) Accept(v Visitor) {
	v.VisitNeighbour(n)
}

type Friend struct {
}

// Имплементация Accept для друга
func (f *Friend) Accept(v Visitor) {
	v.VisitFriend(f)
}

// DoneChores удовлетворяет ранее описанному интерфейсу Visitor - позволяет выполнить нужное действие для каждого посещенного объекта.
type DoneChores struct {
	Done []string
}

func (dc *DoneChores) VisitGranny(granny *Granny) {
	dc.Done = append(dc.Done, "Eat some cakes")
}

func (dc *DoneChores) VisitNeighbour(neighbour *Neighbour) {
	dc.Done = append(dc.Done, "Seat with neighbour's dog")
}

func (dc *DoneChores) VisitFriend(friend *Friend) {
	dc.Done = append(dc.Done, "Watch movies")
}

//Пример использования:
// func main() {
//>>>>>>>Создаем необходимые элементы, которые будем посещать
// 	closePeople := []ClosePeople{
// 		&Granny{},
// 		&Neighbour{},
// 		&Friend{},
// 	}
// >>>>>>>Создаем посетителя
// 	doneChores := &DoneChores{Done: make([]string, 0, 3)}
// 	for _, closeHuman := range closePeople {
//>>>>>>>>Выполняем необходимые действия при посещении.
// 		closeHuman.Accept(doneChores)
// 	}

// 	fmt.Println("Done chores: ", doneChores.Done)
// }
