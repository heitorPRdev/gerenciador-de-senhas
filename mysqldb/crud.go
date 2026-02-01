package crud

import (
	ConnFile "crud/connection/data/mysqldb/connection"
	"database/sql"
	"log"
)

func InsertPassword(senha_local string, senha string) {
	var conn *sql.DB
	conn = ConnFile.Connection()
	query := "insert into passwords(senha_local,senha) values(?,?)"
	_, err := conn.Exec(query, senha_local, senha)
	if err != nil {
		log.Fatal(err)
	}

}
