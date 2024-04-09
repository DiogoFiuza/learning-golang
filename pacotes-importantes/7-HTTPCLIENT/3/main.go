package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()                          // declarate a new context
	ctx, cancel := context.WithTimeout(ctx, time.Second) // define timeout for context
	//ctx, cancel := context.WithCancel(ctx) // Here the context stop when its cancel
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	print(string(body))
}
