package db_client

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DBClient *sqlx.DB

// var DBClient *sql.DB

func InitialiseDBConnection() {
	// sqlx
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/test_db?parseTime=true")

	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	// sqlx
	// db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test_db?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	DBClient = db
}
