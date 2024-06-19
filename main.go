package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("New Request")
	log.Println("===========================================================")
	for k, v := range r.Header {
		log.Printf("%v: %v\n", k, v)
	}
	log.Println("===========================================================")
}

func main() {
	godotenv.Load()
	http.HandleFunc("/", handler)
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	log.Printf("Running server on host address: %s and port: %s\n", host, port)
	http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)

}
