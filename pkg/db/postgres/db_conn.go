package postgres

import (
	"database/sql"
	"example/config"
	"fmt"
)

func NewDbConn() *sql.DB {
	cnf := config.GetConfig()
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cnf.Host, cnf.Port, cnf.User, cnf.Password, cnf.Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}
