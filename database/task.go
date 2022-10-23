package database

import(
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


func InsertTask(task string){
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/go_todo")
	if err!= nil {
        log.Println(err.Error())
    }
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO todo VALUES(0, ?)")
	if err!= nil {
        log.Println(err.Error())
    }
	defer stmt.Close()

	res, err := stmt.Exec(task)
	if err != nil {
		log.Fatal(err)
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(lastInsertID)
}