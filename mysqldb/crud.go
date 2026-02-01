package crud

import (
	"crud/connection/data/data"
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
func SelectAllPass() []data.ColunsSenhas {
	var allColuns []data.ColunsSenhas
	var conn *sql.DB
	conn = ConnFile.Connection()
	query := "select * from passwords"
	rows, err := conn.Query(query)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var coluns data.ColunsSenhas
		if err := rows.Scan(&coluns.Id, &coluns.SenhaLocal, &coluns.Senha); err != nil {
			log.Fatal(err)
		}
		allColuns = append(allColuns, coluns)
	}
	return allColuns

}
