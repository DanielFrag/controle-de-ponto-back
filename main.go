package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"bitbucket.org/DanielFrag/gestor-de-ponto/router"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		clean()
	}()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func clean() {
	stopDb()
	os.Exit(1)
}

func stopDb() {
	fmt.Println("DB closed")
}
