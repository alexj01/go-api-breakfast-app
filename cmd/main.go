package main

import (
	"net/http"
	"os"
	"log"
	"github.com/alexj01/go-api-breakfast-app/pkg/controller"
)

func main() {
	controller.Startup()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("port"), nil))
}
