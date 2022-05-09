package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	res, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// body, err := io.ReadAll(res.Body)

	// res.Body.Close()

	// if res.StatusCode > 299 {
	// 	log.Fatalf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	// }
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// fmt.Printf("%s", body)

	// Another way to parse response
	// byteSize := 99999
	// bs := make([]byte, byteSize)
	// res.Body.Read(bs)
	// fmt.Println(string(bs))

	// implements both Writer and Reader interfaces
	// io.Copy(os.Stdout, res.Body)

	// with custom writter implementation
	lw := logWriter{}
	io.Copy(lw, res.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
