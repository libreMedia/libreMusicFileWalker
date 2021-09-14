package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func getCwd() (p string) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path // for example /home/user+
}

func getParentDir() (p string) {
	parent := filepath.Dir(getCwd())
	return parent
}

func deleteFile(inputFile string) {
	cmdOutput, err := exec.Command("rm", inputFile).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", cmdOutput)
}
