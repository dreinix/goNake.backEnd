package database

import (
	"database/sql"
	"testing"
)

func TestTableCreation(t *testing.T) {
	DbTest, err := sql.Open("postgres", "postgresql://root@localhost:26257/gonaketest_db?sslmode=disable")
	if err != nil {
		t.Errorf("connect to database test failed expected no error got %v", err)
	}
	if _, err := DbTest.Exec(`
		CREATE TABLE IF NOT EXISTS tbl_user (
		usr_id SERIAL PRIMARY KEY,
		full_name STRING(50) NOT NULL,
		usrn STRING(30) NOT NULL UNIQUE,
		pwd STRING(30) NOT NULL,
		stat STRING(8) NOT NULL)`); err != nil {
		t.Errorf("create table tbl_user test failed expected no error got %v", err)
	}
	if _, err := DbTest.Exec(`	
			CREATE TABLE IF NOT EXISTS tbl_score (
			score_id SERIAL PRIMARY KEY,
			score INT2 NOT NULL,
			usr int references tbl_user(usr_id) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL,
			scored_on TIMESTAMP)`); err != nil {
		t.Errorf("create table tbl_uscore test failed expected no error got %v", err)
	}
}
