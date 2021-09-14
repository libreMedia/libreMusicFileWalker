package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func walk() {
	err := filepath.Walk(getParentDir()+"/music",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if suffixChecker(path, ".mp3") || suffixChecker(path, ".ogg") || suffixChecker(path, ".wav") {
				fmt.Println(path, info)
			} else {
				fmt.Println("things")
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
