package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DbConnection() (*sql.DB , error) {
	dbUrl := "root:root@tcp(localhost:3306)/go_backend_dev?parseTime=true"
	return sql.Open("mysql" , dbUrl);
}
