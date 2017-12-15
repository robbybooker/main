package maths

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"int_type"
)

type maths struct {
	counter int
}

func New() *maths {
	i := int_type.IntType(3)
	i.Increment()

	return &maths{counter: 10}
}

func (this *maths) AddNumbers(n1 int, n2 int) int {
	return n1 + n2
}

func (this *maths) IncrementNumber(n1 *int) {
	*n1++
}

func (this *maths) IncrementCounter() {
	this.counter++;
}

func (this *maths) GetCounter() int {
	return this.counter;
}

func (this *maths) Error() {
	defer func() {
		fmt.Println("DEFERRING")
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	b := 0
	a := 1 / b

	fmt.Println(a)
	panic("ERROR")
	recover()
}
