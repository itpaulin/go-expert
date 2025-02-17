package main

import (
	"io"
	"net/http"
)

func main() {
	c := http.Client{}

	req, err := http.NewRequest("GET", "http://www.google.com", nil)
	if err != nil {
		panic(err)
	}
	res, err := c.Do(req)

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
