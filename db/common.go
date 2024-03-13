package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var err error
var db *sql.DB

func DbConnectSlqServer() error {
	var server = "SEBASTIAN"
	var port = 1433
	var user = "sa"
	var password = "123456$"
	var database = "Gambit"

	var connectionString = fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s;",
		server, port, user, password, database)

	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err.Error())
	}

	// var result string
	// ctx := context.Background()

	// erro := db.QueryRowContext(ctx, "SELECT @@version").Scan(&result)

	// if erro != nil {
	// 	log.Fatal(err.Error())
	// }

	return err
}
