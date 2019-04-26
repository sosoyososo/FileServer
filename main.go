package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("\nYou can use first parameter as port (:8080 for instance), second parameter as path (relative or absolute)\n ")

	port := ":8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	path := "./"
	if len(os.Args) > 2 {
		path = os.Args[2]
	}
	fmt.Printf("Server at path %s with port %s\n", port, path)
	log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir(path))))
}
