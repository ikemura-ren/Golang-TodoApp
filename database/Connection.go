package database

import(
	"database/sql"//データベースとの接続のためのインターフェース
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQLのコネクションドライバ
)

func ConnectionDB(){
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/go_todo") // ローカルのMySQLのDSN
    if err!= nil {
        log.Println(err.Error())//もし変数errがnilではない場合、エラーログを表示
    }
	defer db.Close()//遅延実行で最終的にはコネクションを閉じている

	err = db.Ping()//この関数はDBとの接続を監視しており、必要であれば接続を管理する
    if err!= nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to the database")//接続が確立していたら出る内容
	}
}