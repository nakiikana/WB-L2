package pattern

 
type Strategy interface {
	Execute(int, int) int
}

 
type AddStrategy struct{}

func (s *AddStrategy) Execute(a, b int) int {
	return a + b
}

type SubtractStrategy struct{}

func (s *SubtractStrategy) Execute(a, b int) int {
	return a - b
}

type MultiplyStrategy struct{}

func (s *MultiplyStrategy) Execute(a, b int) int {
	return a * b
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}

/* пример вывода 

func main() {
	context := &Context{}

	context.SetStrategy(&AddStrategy{})
	result := context.ExecuteStrategy(10, 5)
	fmt.Printf("10 + 5 = %d\n", result)

	context.SetStrategy(&SubtractStrategy{})
	result = context.ExecuteStrategy(10, 5)
	fmt.Printf("10 - 5 = %d\n", result)

	context.SetStrategy(&MultiplyStrategy{})
	result = context.ExecuteStrategy(10, 5)
	fmt.Printf("10 * 5 = %d\n", result)
}

*/



/*
	Паттерн "Стратегия" (Strategy) — это поведенческий паттерн проектирования, который позволяет определить набор алгоритмов, инкапсулировать каждый из них и сделать их взаимозаменяемыми. Этот паттерн позволяет выбирать необходимый алгоритм на лету, в зависимости от конкретной ситуации.
	Когда у вас есть несколько схожих классов с разным поведением и вам нужно, чтобы поведение можно было легко изменять и расширять без изменения основного класса.
	Позволяет легко добавлять новые алгоритмы без изменения существующего кода, но увеличивает количество классов и может привести к увеличению сложности структуры программы.
*/
