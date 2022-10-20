package database

import(
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDB(){
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/go_todo")
    if err!= nil {
        log.Println(err.Error())
    }
	defer db.Close()

	err = db.Ping()
    if err!= nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to the database")
	}
}