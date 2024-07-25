package pattern

import "log"

/* Команда - поведенческий паттерн, который используется для выполнения команд с разной имплементацией, но одинаковыми шагами для ее
реализации.
Составляющие паттерна:
- интерфейс команды (может быть также Undo)
- инициатор (Invoker) - объект, который знает, какую команду выполнять
- цель (Receiver) - объект, который получает запросы от команды
*/

type Command interface {
	Execute()
}

// Цель (receiver) в контексте данного примера
type Enterprise struct {
	TotalDocuments int
	TotalPaper     int
}

// Конструктор объекта receiver'а
func NewEnterprise(docs, paper int) *Enterprise {
	return &Enterprise{TotalDocuments: docs, TotalPaper: paper}
}

// Конкретные команды
type PrintDocumentsCommand struct {
	n          int
	enterprise *Enterprise
}

func (c *PrintDocumentsCommand) Execute() {
	if c.enterprise.TotalPaper-c.n >= 0 {
		c.enterprise.TotalDocuments += c.n
		c.enterprise.TotalPaper -= c.n
	} else {
		log.Println("You don't have enough paper")
	}
}

type HandToSignCommand struct {
	n          int
	enterprise *Enterprise
}

func (c *HandToSignCommand) Execute() {
	if c.n <= c.enterprise.TotalDocuments {
		c.enterprise.TotalDocuments -= c.n
	} else {
		log.Println("You don't have enough documents to sign")
	}
}

type ReceiveDocumentsCommand struct {
	n          int
	enterprise *Enterprise
}

func (c *ReceiveDocumentsCommand) Execute() {
	c.enterprise.TotalDocuments += c.n
}

// Методы для создания объектов команд (из цели)
func (e *Enterprise) PrintDocuments(n int) Command {
	return &PrintDocumentsCommand{enterprise: e, n: n}
}

func (e *Enterprise) HandToSign(n int) Command {
	return &HandToSignCommand{enterprise: e, n: n}
}

func (e *Enterprise) ReceiveDocuments(n int) Command {
	return &ReceiveDocumentsCommand{enterprise: e, n: n}
}

// Инициатор (invoker), содержищий список команд и способный их выполнять
type Employee struct {
	Commands []Command
}

func (e *Employee) ExecuteCommands() {
	for _, c := range e.Commands {
		c.Execute()
	}
}
