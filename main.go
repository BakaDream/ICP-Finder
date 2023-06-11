package main

import (
	"fmt"
	"log"
	"mengdiao/ICP-Finder/Handler"
	"net/http"
)

func main() {
	http.HandleFunc("/geticpinfo", Handler.GetICPInfoHandler)
	fileServe := http.FileServer(http.Dir("www"))
	http.Handle("/", fileServe)
	fmt.Println("Server listening on port 8080", "http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
