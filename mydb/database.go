package mydb

import (
	"fmt"
	"sync"
)

var once sync.Once

var isActive bool
var lock = sync.Mutex{}
var stored map[string]string

func New() {
	lock.Lock()
	defer lock.Unlock()

	if !isActive {
		stored = make(map[string]string)
		isActive = true
	} else {
		fmt.Println("Sono un singleton!!")
	}
}

func Close() {
	lock.Lock()
	defer lock.Unlock()
	if !isActive {
		return
	}
	isActive = false
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
