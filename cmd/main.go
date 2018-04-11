package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alexj01/go-api-breakfast-app/pkg/controller"
	_ "github.com/lib/pq"
)

const (
	dbHost  = "localhost"
	dbPort  = 5432
	dbName  = "breakfast_db"
	dbUser  = "breakfast"
	sslMode = "disable"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s "+
		"port=%d "+
		"user=%s "+
		"dbname=%s "+
		"sslmode=%s", dbHost, dbPort, dbUser, dbName, sslMode)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Printf("Successfully connected to %s.", dbName)

	controller.Startup()

	log.Fatal(http.ListenAndServe(":"+os.Getenv("port"), nil))
}
