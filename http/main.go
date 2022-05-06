package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	body, err := io.ReadAll(res.Body)

	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("%s", body)

	// Another way to parse response
	// byteSize := 99999
	// bs := make([]byte, byteSize)
	// res.Body.Read(bs)
	// fmt.Println(string(bs))

}
