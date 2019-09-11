package mydb

import (
	"fmt"
	"sync"
)

var once sync.Once // allow only one operation to be called

var lock = sync.Mutex{}      // allow one operation at a time (thread-safe)
var stored map[string]string // the actual fake db

// New create the new singleton object
func New() {
	lock.Lock()
	defer lock.Unlock()

	// initialize the singleton just once.
	// future calls of New() won't run this block
	once.Do(func() {
		fmt.Println("new class")
		stored = make(map[string]string)
	})
}

// Close delete the singleton object
func Close() {
	lock.Lock()
	defer lock.Unlock()
	once = sync.Once{}
	stored = nil
}

// Store save a key, value pair in the database
func Store(k, v string) {
	lock.Lock()
	defer lock.Unlock()
	stored[k] = v
}

// PrintAll print the database
func PrintAll() {
	lock.Lock()
	defer lock.Unlock()
	for k, v := range stored {
		fmt.Println(k, ":", v)
	}
}
