package main

import (
	"fmt"
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
			if suffixChecker(path, ".mp3") ||
				suffixChecker(path, ".ogg") ||
				suffixChecker(path, ".wav") ||
				suffixChecker(path, ".flac") {
				// fmt.Println(info)
				routePath := slashReplacer(path)[3:]
				fmt.Println(routePath)
				dbInsert(dbName, path, routePath, path)
			} else {
				fmt.Println("things")
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
