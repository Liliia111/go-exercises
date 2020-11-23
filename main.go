package main

import (
	"awesomeProject4/task/cmd"
	"awesomeProject4/task/db"
)

func main() {
	dbPath := "C:\\Users\\User\\go\\src\\awesomeProject4\\task\\tasks.db"
	err := db.Init(dbPath)
	if err != nil {
		panic(err)
	}
	cmd.RootCmd.Execute()
}
