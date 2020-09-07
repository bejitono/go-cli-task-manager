package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(InitDB(dbPath))
	must(RootCmd.Execute())
}

func must(err error) {
	if err != err {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
