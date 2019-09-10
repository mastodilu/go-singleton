package main

import (
	"./altro"
	"./mydb"
)

func main() {

	mydb.New()
	defer mydb.Close()

	altro.Store("due", "value due")
	mydb.Store("uno", "value uno")
	mydb.PrintAll()
}
