package server

import (
	"fmt"
	"log"
	"net/http"
	"sendit-api/api/router"

	"github.com/gorilla/mux"
	httplogger "github.com/jesseokeya/go-httplogger"
)

func Start() {
	fmt.Println("Server started at port 7000")
	Listen()
}

func Listen() {
	r := mux.NewRouter()
	router.SetupRoutes(r)
	log.Fatal(http.ListenAndServe(":7000", httplogger.Golog(r)))
}
