package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Conect() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/gonake_db?sslmode=disable")
	if err != nil {

		log.Fatal("error connecting to the database: ", err)
		return nil, err
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS tbl_user (usr_id SERIAL PRIMARY KEY,
		full_name STRING(50) NOT NULL,
		usrn STRING(30) NOT NULL,
		pwd STRING(30) NOT NULL)`); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS tbl_score (score_id SERIAL PRIMARY KEY,
			value INT2 NOT NULL,
			usr int references tbl_user(usr_id) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL)`); err != nil {
		log.Fatal(err)
	}
	return db, nil
}

/*
func insert(tbl string, User user) {
	//command := fmt.Sprintf("INSERT INTO %s (full_name, department, designation, created_at, updated_at) VALUES ('Irshad', 'IT', 'Product Manager', NOW(), NOW());", tbl)
	if _, err := db.Exec(
		`INSERT INTO tbl_employee (full_name, department, designation, created_at, updated_at)
		VALUES ('Irshad', 'IT', 'Product Manager', NOW(), NOW());`); err != nil {
		log.Fatal(err)
	}
}*/
