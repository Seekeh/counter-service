package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type StateData struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

func main() {
	count := 0

	http.HandleFunc("/increment",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				log.Printf("Error: Got method %s instead of POST\n", r.Method)
				w.WriteHeader(400)
				return
			}

			count = count + 1
			log.Printf("Current counter incremented to: %d\n", count)
		})

	http.HandleFunc("/counter",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				log.Printf("Error: Got method %s instead of GET\n", r.Method)
				w.WriteHeader(400)
				return
			}

			log.Printf("Current counter is: %d\n", count)
			w.WriteHeader(200)
			fmt.Fprintf(w, strconv.Itoa(count))
		})

	log.Printf("Starting counter-service...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
