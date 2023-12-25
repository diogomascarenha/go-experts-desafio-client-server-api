package services

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type USDBRL struct {
	Bid string `json:"bid"`
}
type ExchangeResponse struct {
	USDBRL USDBRL `json:"USDBRL"`
}

type GetUSDollarExchangeRateToBrazilianReais struct {
	Timeout time.Duration
}

func NewGetUSDollarExchangeRateToBrazilianReais(timeout time.Duration) *GetUSDollarExchangeRateToBrazilianReais {
	return &GetUSDollarExchangeRateToBrazilianReais{
		Timeout: timeout,
	}
}

func (g *GetUSDollarExchangeRateToBrazilianReais) Execute(ctx context.Context) (string, error) {

	//ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)

	ctx, cancel := context.WithTimeout(ctx, g.Timeout)
	defer cancel()

	httpClient := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return "", err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var exchangeResponse ExchangeResponse
	err = json.Unmarshal(body, &exchangeResponse)

	if err != nil {
		return "", err
	}

	return exchangeResponse.USDBRL.Bid, nil
}
