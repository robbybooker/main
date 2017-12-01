package controller

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"int_type"
	"github.com/ant0ine/go-json-rest/rest"
)

type controller struct {
	counter int
}

func New() *controller {

	//ii := int(400)

	i := int_type.IntType(3)

	i.Increment()

	//fmt.Println(ii)

	moo := controller {}
	moo.counter = 10
	return &moo;

	//return new(controller)
	//return &controller{}
}

func (this *controller) AddNumbers(n1 int, n2 int) int {
	return n1 + n2
}

func (this *controller) IncrementNumber(n1 *int) {
	*n1++
}

func (this *controller) IncrementCounter() {
	this.counter++;
}

func (this *controller) ShowCounter() {
	fmt.Println(this.counter);
}

func (this *controller) Error() {
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

func (this *controller) GetDbVersion(w rest.ResponseWriter, r *rest.Request) {
	db, _ := sql.Open("mysql", "root:woofwoof@/rob")
	defer db.Close()

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)

	w.WriteJson(version)
}

func (this *controller) ChannelPlay() {

}
