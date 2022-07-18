package main

import (
	"strings"
)

func suffixChecker(inputName string, suffix string) (r bool) {
	res := strings.HasSuffix(inputName, suffix)
	return res
}

func slashReplacer(inputFilePath string) (r string) {

	res := strings.Replace(inputFilePath, "\\", "/", -1)
	return res
}

// func runCliCmd(inputFilePath string) (r string) {
// 	cmd := exec.Command("gulp", "serv.dev")
// 	if err := cmd.Run(); err != nil {
// 		log.Fatal(err)
// 	}
// 	return
// }
