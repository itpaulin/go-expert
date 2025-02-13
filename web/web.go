package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	cep_off "github.com/itpaulin/go-expert/cep_off"
)

func Web() {
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	cepStr := r.URL.Query().Get("cep")
	if cepStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c, err := BuscaCep(cepStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(c)
}
func BuscaCep(cepReq string) (*cep_off.CepResponse, error) {
	res, err := http.Get("https://viacep.com.br/ws/" + cepReq + "/json/")
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, fmt.Errorf("failed to get response")
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var c cep_off.CepResponse
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
