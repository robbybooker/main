package maths

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"int_type"
	"github.com/ant0ine/go-json-rest/rest"
)

type maths struct {
	counter int
}

func New() *maths {

	//ii := int(400)

	i := int_type.IntType(3)

	i.Increment()

	//fmt.Println(ii)

	moo := maths {}
	moo.counter = 10
	return &moo;

	//return new(controller)
	//return &controller{}
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

func (this *maths) ShowCounter() {
	fmt.Println(this.counter);
}

func (this *maths) Error() {
	defer func() {
		fmt.Println("DEFERRING")
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("wwwwwwwwwww")

	b := 0
	a := 1 / b

	fmt.Println(a)
	panic("ERROR")
	recover()
}

func (this *maths) GetDbVersion(w rest.ResponseWriter, r *rest.Request) {
	db, _ := sql.Open("mysql", "root:woofwoof@/rob")
	defer db.Close()

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)

	w.WriteJson(version)
}

func (this *maths) ChannelPlay() {

}
