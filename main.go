package main

import (
	"fmt"
	"mengdiao/ICP-Finder/Handler"
	"net/http"
)

func main() {
	http.HandleFunc("/geticpinfo", Handler.GetICPInfoHandler)
	fileServe := http.FileServer(http.Dir("www"))
	http.Handle("/", fileServe)
	fmt.Println("Server listening on port 8080", "http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
