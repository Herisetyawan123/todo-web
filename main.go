package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
	"os"
	"sync"
	"todo-web/route"
	"todo-web/utils"
)

type Response struct {
	status string `json:"status"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Response{
		status: "201",
	})
}

func main() {
	os.Setenv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/todo-web")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		// connect to server
		mux := http.NewServeMux()

		// connect to postgres
		err := utils.ConnectDb()
		if err != nil {
			panic(err)
		}

		db := utils.GetDBConnection()

		route.Api(mux, db)

		fmt.Println("Server is running on port 8080")
		err = http.ListenAndServe(":8080", mux)
		if err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}
