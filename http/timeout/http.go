package main

import (
	"io"
	"net/http"
	"time"
)

func main() {

	c := http.Client{Timeout: time.Nanosecond} // set timeout to all request

	res, err := c.Get("http://www.google.com")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))

}
