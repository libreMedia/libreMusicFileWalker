package main

func maine() {
	dbCreate("musicDB")
	walk("../music", "musicDB")
}
