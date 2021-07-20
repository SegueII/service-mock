package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", QueryRateHandlerFn)
	log.Println("Starting mock server ...")
	log.Fatal(http.ListenAndServe(":8083", nil))
}

func QueryRateHandlerFn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("content-type", "application/json")

	println("receive request ...")
	_, _ = w.Write([]byte("100"))
}
