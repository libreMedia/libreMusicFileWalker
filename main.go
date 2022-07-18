package main

func main() {
	dbCreate("musicDB")
	walk("../music", "musicDB")
}
