package persistency

import (
	"Api-Aula1-golang/config"
	"database/sql"
	"log"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.Cfg.FormatDSN())
	if err != nil {
		log.Println("ERRO ao conectar ao banco de dados:")
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		log.Println("ERRO ao pingar o banco de dados:")
		return nil, err
	}
	return db, nil
}
