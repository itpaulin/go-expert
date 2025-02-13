package cep_off

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func GetInfoCep(cep string) *CepResponse {

	res, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		fmt.Fprint(os.Stderr, "Erro ao fazer get no viacep", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprint(os.Stderr, "Erro ao ler body", err)
	}
	defer res.Body.Close()

	var data CepResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Fprint(os.Stderr, "Erro ao fazer o bind para json", err)
	}

	file, err := os.Create("cidade.txt")
	if err != nil {
		fmt.Fprint(os.Stderr, "Erro ao criar o arquivo", err)
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s\n", data.Cep, data.Localidade, data.Uf))
	if err != nil {
		fmt.Fprint(os.Stderr, "Erro ao escrever o arquivo", err)
	}
	fmt.Println("Arquivo cidade.txt criado com sucesso!")

	fmt.Printf("O resultado da sua consulta:\n %v", data)

	return &data
}
