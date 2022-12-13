package mysqlConnection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Inicia a conexão com o banco de dados
func InitConnection() {
	user := os.Getenv("mysqlUser")
	pass := os.Getenv("mysqlPass")
	stringConnection := user + ":" + pass + "@tcp(0.0.0.0:3306)/devbook"
	db, erro := sql.Open("mysql", stringConnection)
	if erro != nil {
		log.Fatalln(erro)
	}
	defer db.Close()
	if erro = db.Ping(); erro != nil {
		log.Fatalln(erro)
	}
	fmt.Println("Conexão com sucesso")
}

func CreateUser() {

}
