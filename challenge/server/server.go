package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CotacaoResponse struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		fmt.Println("Erro ao abrir banco de dados:", err)
		return
	}
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		HandleCotacao(w, r, db)
	})
	fmt.Println("Servidor rodando na porta 8080...")

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
	}

}

func HandleCotacao(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	cotacao, err := GetCotacao(ctx)
	if err != nil {
		http.Error(w, "Erro ao obter cotação", http.StatusInternalServerError)
		fmt.Println("Erro ao obter cotação:", err)
		return
	}
	err = CreateCotacao(db, cotacao)
	if err != nil {
		http.Error(w, "Erro ao salvar no banco", http.StatusInternalServerError)
		fmt.Println("Erro ao salvar cotação no banco:", err)
		return
	}
	w.Write([]byte(cotacao.USDBRL.Bid))
}

func GetCotacao(ctx context.Context) (*CotacaoResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API retornou status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cotacao CotacaoResponse

	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		return nil, err
	}

	return &cotacao, nil
}

func CreateCotacao(db *sql.DB, c *CotacaoResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	stmt, err := db.PrepareContext(ctx, "INSERT INTO exchange_rates (code, codein, name, high, low, var_bid, pct_change, bid, ask, timestamp, create_date) VALUES (? , ? , ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.USDBRL.Code, c.USDBRL.Codein, c.USDBRL.Name, c.USDBRL.High, c.USDBRL.Low, c.USDBRL.VarBid, c.USDBRL.PctChange, c.USDBRL.Bid, c.USDBRL.Ask, c.USDBRL.Timestamp, c.USDBRL.CreateDate)

	return err

}
