package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.Open("post.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	resp, err := http.Post("http://localhost:18888", "text/plain", file)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", resp.Status)
}
