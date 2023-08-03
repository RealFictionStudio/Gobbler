package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func checkIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func fileExist(path string, fname string) bool {

	dir, err := os.ReadDir(path)
	checkIfError(err)

	for k, v := range dir {
		if v.Name() == fname {
			return true
		}
		_ = k
	}

	return false
}

func main() {
	var sep string

	if len(os.Args) > 1 {
		if os.Args[1] == "windows" {
			sep = "\\"
		} else if os.Args[1] == "unix" {
			sep = "/"
		} else {
			log.Fatal("Unsupported/Unknown OS specify if you're on\nwindows   - go run setup.go windows\nlinux/mac - go run setup.go unix")
		}
	} else {
		log.Fatal("Unspecified OS\nTry typing (if you're on windows) \"go run setup.go windows\"\nor if you're on linux/mac \"go run setup.go unix\"")
	}

	fmt.Println(`
	  _____       _     _     _           
	 / ____|     | |   | |   | |          
   	| |  __  ___ | |__ | |__ | | ___ _ __ 
   	| | |_ |/ _ \| '_ \| '_ \| |/ _ \ '__|
   	| |__| | (_) | |_) | |_) | |  __/ |   
	 \_____|\___/|_.__/|_.__/|_|\___|_|
                 _             
	 ___ ___| |_ _  _ _ __ 
	(_-</ -_)  _| || | '_ \
	/__/\___|\__|\_,_| .__/
			 |_|
	`)

	cwd, err := os.Getwd()
	checkIfError(err)

	var dirs []string = strings.Split(cwd, sep)
	err = os.Chdir(strings.Join(dirs, sep) + sep + "src")
	checkIfError(err)
	cwd, err = os.Getwd()
	checkIfError(err)

	var files [][]string = [][]string{
		{"sended.log", ""}, {".env", "TOKEN=\nCHANNEL="},
	}

	for k, v := range files {
		if !fileExist(cwd, v[0]) {
			f, err := os.Create(v[0])
			checkIfError(err)
			_, err = f.WriteString(v[1])
			checkIfError(err)
			_ = k
			f.Close()
			fmt.Println("Created FILE " + v[0])
		} else {
			fmt.Println("FILE " + v[0] + " already exist")
		}
	}

	exec.Command("go mod tidy")
}
