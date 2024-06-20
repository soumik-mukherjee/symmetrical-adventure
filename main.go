package main

import (
	"fmt"
	"log"
	"net/http"
	"symmetrical-adventure/internal/configs"
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
	// load configurations from config file
	config, err := configs.NewConfiguration()
	if err != nil {
		fmt.Errorf("could not load configuration: %v", err)
	}
	http.HandleFunc("/", handler)
	host := config.Fe.Host //os.Getenv("HOST")
	port := config.Fe.Listen_port
	log.Printf("Running server on host address: %s and port: %s\n", host, port)
	http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
}
