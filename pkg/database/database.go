package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // O underline import é necessário
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "postgres"
)

func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao testar conexão: %v", err)
	}

	fmt.Println("Conexão com PostgreSQL estabelecida!")
	return db
}
