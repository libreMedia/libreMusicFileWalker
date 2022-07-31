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
	name       string
	routePath  string
	path       string
	title      string
	artist     string
	album      string
	year       string
	givenGenre string
	votedGenre string
	comment    string
	composer   string
	lyrics     string
}

func dbCreate(dbName string) {

	db, err := sql.Open("sqlite3", "./"+dbName+".db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table musicDB (name text not null primary key, routePath text, path text, title text, artist text, album text, year text, givenGenre text, votedGenre text, comment text, composer text, lyrics text);
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
		modelRow := dbModel{name: names, routePath: routePath, path: path}
		modelArray = append(modelArray, modelRow)
	}
	return modelArray
}

func dbInsert(dbName string, name string, routePath string, path string, title string, artist string, album string, year string, givenGenre string, votedGenre string, comment string, composer string, lyrics string) {
	db, err := sql.Open("sqlite3", "./"+dbName+".db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	vals := `VALUES("` + name + `" , "` + routePath + `" , "` + path + `" , "` + title + `" , "` + artist + `" , "` + album + `" , "` + year + `" , "` + givenGenre + `" , "` + votedGenre + `" , "` + comment + `" , "` + composer + `" , "` + lyrics + `" )`
	sqlStmt := `
	INSERT INTO musicDB (name,routePath,path,title,artist,album,year,givenGenre,votedGenre,comment,composer,lyrics) ` + vals
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}
