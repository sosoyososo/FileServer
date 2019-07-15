package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func handUploadFile(w http.ResponseWriter, r *http.Request) {
	fileLenStr := r.Header.Get("content-length")
	maxLen := 1024 * 1024 * 100
	fileLen := maxLen
	if len(fileLenStr) > 0 {
		l, err := strconv.Atoi(fileLenStr)
		if nil == err && l > 0 {
			fileLen = l
		}
	}
	if fileLen > maxLen {
		fileLen = maxLen
	}

	err := r.ParseMultipartForm(int64(fileLen))
	if err != nil {
		fmt.Println("Error Parse Form")
		fmt.Println(err)
		return
	}

	file, handler, err := r.FormFile("myfile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)

	tempFile, err := ioutil.TempFile("/tmp", handler.Filename)
	if err != nil {
		fmt.Println("open temp file fail: ")
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	buf, err := ioutil.ReadAll(file)
	if nil != err || len(buf) == 0 {
		fmt.Println("read file err :")
		fmt.Println(err)
		return
	}

	_, err = tempFile.Write(buf)
	if nil != err {
		fmt.Println("write file err :")
		fmt.Println(err)
		return
	}

	http.Redirect(w, r, "./upload.html", 301)
}

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

	fileServer := http.FileServer(http.Dir(path))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handUploadFile(w, r)
		} else {
			fileServer.ServeHTTP(w, r)
		}
	})

	fmt.Printf("Server at path %s with port %s\n", port, path)
	log.Fatal(http.ListenAndServe(port, nil))
}
