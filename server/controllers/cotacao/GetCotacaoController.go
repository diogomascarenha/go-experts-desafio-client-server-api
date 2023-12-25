package cotacao

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/diogomascarenha/go-experts-desafio-client-server-api/server/repositories/exchange_rate"
	"github.com/diogomascarenha/go-experts-desafio-client-server-api/server/repositories/exchange_rate/dtos"
	"github.com/diogomascarenha/go-experts-desafio-client-server-api/server/services"
	"github.com/gookit/goutil/dump"
	log "github.com/sirupsen/logrus"
)

type GetCotacaoController struct {
	GetExchangeRateService *services.GetUSDollarExchangeRateToBrazilianReais
	ExchangeRateRepository exchange_rate.ExchangeRateRepository
}

type GetCotacaoResponseError struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type GetCotacaoResponseSuccess struct {
	Bid string `json:"bid"`
}

func NewGetCotacaoController(
	getExchangeRateService *services.GetUSDollarExchangeRateToBrazilianReais,
	exchangeRateRepository exchange_rate.ExchangeRateRepository,
) *GetCotacaoController {
	return &GetCotacaoController{
		GetExchangeRateService: getExchangeRateService,
		ExchangeRateRepository: exchangeRateRepository,
	}
}

func (c *GetCotacaoController) Execute(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	bid, err := c.GetExchangeRateService.Execute(ctx)

	if err != nil {
		log.Error(err)
		errorResponse := &GetCotacaoResponseError{
			Error:   true,
			Message: err.Error(),
		}
		jsonResponse, _ := json.Marshal(errorResponse)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, string(jsonResponse))
		return
	}

	bidFloat, _ := strconv.ParseFloat(bid, 32)

	input := &dtos.ExchangeRateInputDTO{
		Currency_From: "USD",
		Currency_To:   "BRL",
		Value:         float32(bidFloat),
	}

	resp, err := c.ExchangeRateRepository.Save(ctx, input)

	if err != nil {
		log.Error(err)
		errorResponse := &GetCotacaoResponseError{
			Error:   true,
			Message: err.Error(),
		}
		jsonResponse, _ := json.Marshal(errorResponse)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, string(jsonResponse))
		return
	}

	dump.Println(resp)

	response := &GetCotacaoResponseSuccess{
		Bid: bid,
	}
	jsonResponse, _ := json.Marshal(response)

	log.Infof("Cotação atual: %v", bid)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(jsonResponse))
}
