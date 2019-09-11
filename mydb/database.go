package mydb

import (
	"fmt"
	"sync"
)

var isActive bool            // indicate whether the singleton was instantiated or not
var lock = sync.Mutex{}      // allow one operation at a time
var stored map[string]string // the fake db

// New create a new instance of the singleton object
func New() {
	lock.Lock()
	defer lock.Unlock()

	// if the singleton doesn't exist
	if !isActive {
		// create a singleton
		stored = make(map[string]string)
		isActive = true
	} else {
		// if the singleton exists do nothing
		fmt.Println("You can't do this!")
	}
}

// Close terminate the singleton instance
func Close() {
	lock.Lock()
	defer lock.Unlock()
	// if the singleton doesn't exist, do nothing
	if !isActive {
		return
	}
	// if the singleton exists, delete it
	isActive = false
	stored = nil
}

// Store store a pair key, value into the database
func Store(k, v string) {
	lock.Lock()
	defer lock.Unlock()
	stored[k] = v
}

// PrintAll print the content of the database
func PrintAll() {
	lock.Lock()
	defer lock.Unlock()
	for k, v := range stored {
		fmt.Println(k, ":", v)
	}
}
