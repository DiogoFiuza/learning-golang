package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request started")
	defer log.Println("Request finished")

	select {
	case <-time.After(time.Second * 5):
		// printing on comand line stdout
		log.Println("Request precessed with success")
		// printing on browser
		w.Write([]byte("Request processed with success"))
	case <-ctx.Done():
		// printing on comand line stdout
		log.Println("Resquest canceled by the user")
	}
}
