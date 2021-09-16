package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func removeDb(dbName string) {
	os.Remove("./" + dbName + ".db")
}

func dbCreate(dbName string) {

	db, err := sql.Open("sqlite3", "./"+dbName+".db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (name text not null primary key, path text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func dbInsert(dbName string, name string, path string) {
	db, err := sql.Open("sqlite3", "./"+dbName+".db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	vals := `VALUES("` + name + `" , "` + path + `")`
	sqlStmt := `
	INSERT INTO foo (name,path) ` + vals
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}
