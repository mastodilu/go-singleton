package mydb

import (
	"fmt"
	"sync"
)

var once sync.Once

var lock = sync.Mutex{}
var stored map[string]string

func New() {
	lock.Lock()
	defer lock.Unlock()

	once.Do(func() {
		fmt.Println("new class")
		stored = make(map[string]string)
	})
}

func Close() {
	lock.Lock()
	defer lock.Unlock()
	stored = nil
}

func Store(k, v string) {
	lock.Lock()
	defer lock.Unlock()
	stored[k] = v
}

func PrintAll() {
	lock.Lock()
	defer lock.Unlock()
	for k, v := range stored {
		fmt.Println(k, ":", v)
	}
}
