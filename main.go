package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"sync"
	"todo-web/utils"
)

func main() {
	os.Setenv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/todo-web")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := utils.ConnectDb()
		if err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}
