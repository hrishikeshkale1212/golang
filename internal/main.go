package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash((true))
	myRouter.HandleFunc("/kafka/publisher", pushDataToKafka)
	myRouter.HandleFunc("/kafka", index)
	http.Handle("/", myRouter)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the indexPage!")
	fmt.Println("Endpoint Hit: /")
}

func pushDataToKafka(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pushing data to kafka...")
	fmt.Println("Endpoint Hit: /kafka/publisher")
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequest()
	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
