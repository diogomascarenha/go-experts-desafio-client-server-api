package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/diogomascarenha/go-experts-desafio-client-server-api/server/controllers/cotacao"
	"github.com/diogomascarenha/go-experts-desafio-client-server-api/server/repositories/exchange_rate"
	"github.com/diogomascarenha/go-experts-desafio-client-server-api/server/services"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

func main() {

	LoadEnvVariables()

	db, err := sql.Open("sqlite3", "./db.sqlite?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	GetCotacaoController := GetCotacaoControllerFactory(db)
	mux.HandleFunc("/cotacao", GetCotacaoController.Execute)

	fmt.Println("Server Start")

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func LoadEnvVariables() {
	godotenv.Load()
	if os.Getenv("EXCHANGE_RATE_API_TIMEOUT_IN_MILLISECOND") == "" {
		log.Fatal("Variable EXCHANGE_RATE_API_TIMEOUT_IN_MILLISECOND must be defined")
	}
}

func GetCotacaoControllerFactory(db *sql.DB) *cotacao.GetCotacaoController {
	exchangeRateApiTimeoutInMillisecond, _ := strconv.ParseInt(os.Getenv("EXCHANGE_RATE_API_TIMEOUT_IN_MILLISECOND"), 10, 64)
	exchangeRateApiTimeoutDuration := (time.Duration(exchangeRateApiTimeoutInMillisecond) * time.Millisecond)
	getUSDollarExchangeRateToBrazilianReais := services.NewGetUSDollarExchangeRateToBrazilianReais(exchangeRateApiTimeoutDuration)

	exchangeRateRepository := exchange_rate.NewExchangeRateSqliteRepository(db)

	GetCotacaoController := cotacao.NewGetCotacaoController(
		getUSDollarExchangeRateToBrazilianReais,
		exchangeRateRepository,
	)
	return GetCotacaoController
}
