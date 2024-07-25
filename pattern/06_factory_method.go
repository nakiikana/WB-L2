package pattern

type Object interface {
    GetName() string
}

type ObjectOne struct{}

func (p ObjectOne) GetName() string {
    return "Object One"
}

type ObjectTwo struct{}

func (p ObjectTwo) GetName() string {
    return "Object Two"
}

type Factory interface {
    CreateObject() Object
}

type FactoryOne struct{}

func (f FactoryOne) CreateObject() Object {
    return ObjectOne{}
}

type FactoryTwo struct{}

func (f FactoryTwo) CreateObject() Object {
    return ObjectTwo{}
}
/* пример вывода 

func main() {
    factory1 := ConcreteFactory1{}
    product1 := factory1.CreateProduct()
    fmt.Println("Product:", product1.GetName())

    factory2 := ConcreteFactory2{}
    product2 := factory2.CreateProduct()
    fmt.Println("Product:", product2.GetName())
}

*/


/*
	Паттерн "фабричный метод" (Factory Method) является порождающим паттерном проектирования, который предоставляет интерфейс для создания объектов, но оставляет решение о конкретных классах создаваемых объектов на подклассах. Этот паттерн позволяет делегировать создание объектов подклассам, что способствует более гибкой архитектуре программы.


*/
