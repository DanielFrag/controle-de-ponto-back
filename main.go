package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"bitbucket.org/DanielFrag/gestor-de-ponto/repository"
	"bitbucket.org/DanielFrag/gestor-de-ponto/router"
)

func main() {
	defer mainRecover()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		clean()
	}()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	repository.StartDB()
	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func clean() {
	stopDb()
	os.Exit(1)
}

func stopDb() {
	repository.StopDB()
	fmt.Println("DB closed")
}

func mainRecover() {
	rec := recover()
	if rec != nil {
		log.Println(rec)
		clean()
	}
}
