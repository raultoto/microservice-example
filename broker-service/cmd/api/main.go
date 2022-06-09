package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "3000"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting web server on %s", webPort)

	//define http server

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
