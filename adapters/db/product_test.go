package db

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func setUp() {
	DB, _ = sql.Open("sqlite3", "memory:")
	createTable(DB)
	createProduct(DB)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
			"id" string,
			"name" string,
			"price" float,
			"status" string
            );`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()

}

func createProduct(db *sql.DB) {
	insert := `insert into products values ("abc", "product Test", 0, "disable")`
	Stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	Stmt.Exec()

}
