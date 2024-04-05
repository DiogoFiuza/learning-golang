package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Second}  // Defines an HTTP client with a timeout of one second
	resp, err := c.Get("http://google.com") // Sends a GET request to "http://google.com" using the defined client
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // Reads the response body
	if err != nil {
		panic(err)
	}
	fmt.Print(string(body))
}
