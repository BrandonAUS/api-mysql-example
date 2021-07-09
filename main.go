package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/gorilla/mux"
)

var counter uint64

func main() {
	atomic.AddUint64(&counter, 0)
	router := mux.NewRouter()
	router.HandleFunc("/hello", Counter).Methods("GET")
	fmt.Println("GO REST server running on http://localhost:30456 ")
	log.Fatal(http.ListenAndServe(":30456", router))
}

// Counter increments and get the number of users who hit the pahge
func Counter(w http.ResponseWriter, r *http.Request) {
	atomic.AddUint64(&counter, 1)
	json.NewEncoder(w).Encode(counter)
}
