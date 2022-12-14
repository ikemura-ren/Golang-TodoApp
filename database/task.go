package database

import(
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	Id int
	Task string
	Progress string
}

func ConnectionDB() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/go_todo")
    if err!= nil {
        log.Fatal(err)
    }
	return db
}

func InsertTask(task string){
	db := ConnectionDB()
	stmt, err := db.Prepare("INSERT INTO todo VALUES(0, ?, DEFAULT)")
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

func GetAllTasks() []Task {
	db := ConnectionDB()
	rows, err := db.Query("SELECT * FROM todo")
	if err!= nil {
        log.Println(err.Error())
    }
	defer rows.Close()

	var todo [] Task
	for rows.Next() {
		var id int
		var task string
		var progress string
        err := rows.Scan(&id, &task, &progress)
		if err!= nil {
            log.Fatal(err)
		}
		todo = append(todo, Task{id, task, progress})
	}
	return todo
}

func GetOneTask(id int) Task {
	db := ConnectionDB()
    rows, err := db.Query("SELECT * FROM todo WHERE id=?", id)
    if err!= nil {
        log.Println(err.Error())
    }
	defer rows.Close()
	var todo Task
    for rows.Next() {
        var id int
		var task string
		var progress string
        err := rows.Scan(&id, &task, &progress)
		if err!= nil {
            log.Fatal(err)
        }
		todo = Task{id, task, progress}
	}
	return todo
}

func UpdateTask(id int, task string, progress string){
    db := ConnectionDB()
	stmt, err := db.Prepare("UPDATE todo SET task_name=?, progress=? WHERE id=?")
	if err!= nil {
        log.Fatal(err)
    }
	defer stmt.Close()
	res, err := stmt.Exec(task,progress, id)
	if err!= nil {
        log.Fatal(err)
    }
	lastInsertID, err := res.LastInsertId()
	if err!= nil {
        log.Fatal(err)
    }
	log.Println(lastInsertID)
}

func DeleteTask(id int) {
	db := ConnectionDB()
	stmt, err := db.Prepare("DELETE FROM todo WHERE id=?")
	if err!= nil {
        log.Println(err.Error())
    }
	defer stmt.Close()

	res, err := stmt.Exec(id)
    if err!= nil {
		log.Fatal(err)
	}

	affected, err := res.RowsAffected()
    if err!= nil {
		log.Fatal(err)
	}
	log.Println(affected)
}