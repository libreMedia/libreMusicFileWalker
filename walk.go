package main

import (
	"log"
	"os"
	"path/filepath"
)

func walk(dirToWalk string, dbName string) {
	err := filepath.Walk(getParentDir()+dirToWalk,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if suffixChecker(path, ".mp3") || suffixChecker(path, ".ogg") || suffixChecker(path, ".wav") {
				// fmt.Println(path, info)
				dbInsert(dbName, path, path)
			} else {
				// fmt.Println("things")
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
