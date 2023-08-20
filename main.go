package main

import (
	db "cc/DB"
	"cc/cmd"
)

func main() {
	db.GenerateHashID()
	db.GET()
	cmd.Execute()

}
