package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	name := os.Getenv("NAME")
	fmt.Println("hi", name)

	www := os.Getenv("WWW1")

	if len(www) > 0 {
		r, err := http.Get(www)
		if err != nil {
			log.Fatalf("failed to curl WWW1: %v", err)
		}

		body := r.Body
		defer body.Close()

		content, err := io.ReadAll(body)
		if err != nil {
			log.Fatalf("failed to read response of WWW1: %v", err)
		}

		log.Println("RESPONSE -> ", string(content))
	}
}
