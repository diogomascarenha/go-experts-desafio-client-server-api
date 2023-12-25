package exchange_rate

import (
	"context"

	"github.com/diogomascarenha/go-experts-desafio-client-server-api/server/repositories/exchange_rate/dtos"
)

type ExchangeRateRepository interface {
	Save(ctx context.Context, input *dtos.ExchangeRateInputDTO) (*dtos.ExchangeRateOutputDTO, error)
}
