package main

import (
	"./altro"
	"./mydb"
)

func main() {

	// test if close actually deletes the singleton and
	// it can be re-instantiated in the future
	mydb.New()
	mydb.Close()

	mydb.New()
	mydb.New() // this does nothing
	defer mydb.Close()

	altro.Store("due", "value due")
	mydb.Store("uno", "value uno")
	mydb.PrintAll()
}
