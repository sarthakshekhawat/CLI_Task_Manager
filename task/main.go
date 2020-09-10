package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/sarthakshekhawat/CLI_Task_Manager/task/cmd"
	"github.com/sarthakshekhawat/CLI_Task_Manager/task/db"
)

func main() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	dbPath := filepath.Join(home, "tasks.db")
	err = db.Init(dbPath)
	if err != nil {
		fmt.Println("Something went wrong. Error: ", err)
	}
	cmd.RootCmd.Execute()
}
