package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct {
}

func (Database) CreateSingleConnection() {
	fmt.Println("Creating connection")
	time.Sleep(2 * time.Second)
	fmt.Println("Creation Done")
}

var db *Database
var lock sync.Mutex

func getDatabaseInstance() *Database {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("Creando conexion base de datos")
		db = &Database{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("DB Conn already created")
	}

	return db
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDatabaseInstance()
		}()
	}

	wg.Wait()
}
