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

type dbModel struct {
	Names     string
	RoutePath string
	Path      string
}

func dbCreate(dbName string) {

	db, err := sql.Open("sqlite3", "./"+dbName+".db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table musicDB (name text not null primary key, routePath text, path text);
	delete from musicDB;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func readDb() []dbModel {
	database, _ := sql.Open("sqlite3", "./musicDB.db")
	rows, _ := database.Query("SELECT name, routePath, path FROM musicDB")
	var modelArray []dbModel
	var names string
	var routePath string
	var path string
	for rows.Next() {
		rows.Scan(&names, &routePath, &path)
		modelRow := dbModel{Names: names, RoutePath: routePath, Path: path}
		modelArray = append(modelArray, modelRow)
	}
	return modelArray
}

func dbInsert(dbName string, name string, routePath string, path string) {
	db, err := sql.Open("sqlite3", "./"+dbName+".db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	vals := `VALUES("` + name + `" , "` + routePath + `" , "` + path + `")`
	sqlStmt := `
	INSERT INTO musicDB (name,routePath,path) ` + vals
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}
