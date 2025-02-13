package mux

import "net/http"

func Mux() {
	mux := http.NewServeMux()

	mux.HandleFunc("/welcome", HandleWelcome)
	mux.Handle("/my-name", &name{clientName: "paulin"})
	http.ListenAndServe(":8081", mux)
}

func HandleWelcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Mux Server!"))
}

type name struct {
	clientName string
}

func (n *name) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("My name is Paulo Ricardo, and yours " + n.clientName))
}
