package mariadb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ant0ine/go-json-rest/rest"
)

type database struct {
}

func New() *database {
	return &database{}
}

func (this *database) GetDbVersion(w rest.ResponseWriter, r *rest.Request) {
	db, _ := sql.Open("mysql", "root:woofwoof@/rob")
	defer db.Close()

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)

	w.WriteJson(version)
}
