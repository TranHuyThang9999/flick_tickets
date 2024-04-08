package main

import (
	"file_store/routers"
	"log"
	"net/http"
)

func main() {
	apiRouter := routers.NewApiRouter()

	server := &http.Server{
		Addr:    ":1234",
		Handler: apiRouter.Engine,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
