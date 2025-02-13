package web

import (
	"encoding/json"
	"io"
	"net/http"

	cep "github.com/itpaulin/busca-cep-go/cep"
)

func Web() {
	http.HandleFunc("/")
	http.HandleFunc("/cep", BuscaCep())
	http.ListenAndServe(":8080", nil)
}

func RouteBuscaCep(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello World!"))
	cepStr := r.URL.Query().Get("cep")
	if cepStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	BuscaCep(cepStr)

}
func BuscaCep(cepReq string) (*cep.CepResponse, error) {
	res, err := http.Get("https://viacep.com.br/ws/" + cepReq + "/json/")
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if res != nil {
		return nil, err
	}

	defer res.Body.Close()

	var c cep.CepResponse
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
