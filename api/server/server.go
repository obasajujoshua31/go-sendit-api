package server

import (
	"fmt"
	"log"
	"net/http"
	"sendit-api/api/config"
	"sendit-api/api/database"
	"sendit-api/api/router"

	"github.com/gorilla/mux"
	httplogger "github.com/jesseokeya/go-httplogger"
)

func Start() {
	config.Load()
	database.MigrateAndSeed()
	fmt.Println("Server started at port 7000")
	Listen(config.PORT)
}

func Listen(port int) {
	r := mux.NewRouter()
	router.SetupRoutes(r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), httplogger.Golog(r)))
}
