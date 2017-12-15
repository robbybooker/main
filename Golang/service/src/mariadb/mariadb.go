package mariadb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type database struct {
}

func New() *database {
	return &database{}
}

func (this *database) GetDbVersion() string {
	db, _ := sql.Open("mysql", "root:woofwoof@/rob")
	defer db.Close()

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)

	return version
}
