package main

import (
	"net/http"
	"time"
)

func main() {

	http.Handle("/", http.HandlerFunc(handlerCtx))
	http.ListenAndServe(":8080", nil)
}

func handlerCtx(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-time.After(5 * time.Second):
		println("Request completed!")
		w.Write([]byte("Request completed!"))
		return
	case <-ctx.Done():
		println("Timeout, request canceled!")
		w.Write([]byte("Timeout, request canceled!"))
		return
	}
}
