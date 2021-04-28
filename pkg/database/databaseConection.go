package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Conect(database string) (*sql.DB, error) {
	con := fmt.Sprintf("postgresql://root@localhost:26257/%s_db?sslmode=disable", database)
	db, err := sql.Open("postgres", con)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
		return nil, err
	}
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS tbl_user (
		usr_id SERIAL PRIMARY KEY,
		full_name STRING(50) NOT NULL,
		usrn STRING(30) NOT NULL UNIQUE,
		pwd STRING(30) NOT NULL,
		stat STRING(8) NOT NULL)`); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS tbl_score (
			score_id SERIAL PRIMARY KEY,
			score INT2 NOT NULL,
			usr int references tbl_user(usr_id) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL,
			scored_on TIMESTAMP)`); err != nil {
		log.Fatal(err)
	}
	return db, nil
}
