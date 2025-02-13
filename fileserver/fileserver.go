package fileserver

import (
	"log"
	"net/http"
)

func FileServer() {
	fileServer := http.FileServer(http.Dir("./fileserver/public/"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to File Server!"))
	})
	log.Fatal(http.ListenAndServe(":8080", mux))

}
