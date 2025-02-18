package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer requisição:", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Erro ao ler o conteúdo:", err)
	}

	file, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = os.WriteFile("cotacao.txt", []byte(fmt.Sprintf("Dólar: %s", string(body))), 0644)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
	}
}
