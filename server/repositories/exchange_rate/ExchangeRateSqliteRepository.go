package exchange_rate

import (
	"context"
	"database/sql"

	"github.com/diogomascarenha/go-experts-desafio-client-server-api/server/repositories/exchange_rate/dtos"
)

type ExchangeRateSqliteRepository struct {
	DB *sql.DB
}

func NewExchangeRateSqliteRepository(db *sql.DB) *ExchangeRateSqliteRepository {
	return &ExchangeRateSqliteRepository{
		DB: db,
	}
}

func (r *ExchangeRateSqliteRepository) Save(ctx context.Context, input *dtos.ExchangeRateInputDTO) (*dtos.ExchangeRateOutputDTO, error) {

	stmt, err := r.DB.PrepareContext(ctx, `
        INSERT INTO exchange_rate (currency_from, currency_to, value) 
                           VALUES (?, ?, ?) 
                        RETURNING id, currency_from, currency_to, value, created_at`,
	)
	defer stmt.Close()

	if err != nil {
		return nil, err
	}

	output := dtos.ExchangeRateOutputDTO{}
	stmt.QueryRowContext(ctx, input.Currency_From, input.Currency_To, input.Value).Scan(
		&output.ID,
		&output.Currency_From,
		&output.Currency_To,
		&output.Value,
		&output.CreatedAt,
	)

	return &output, nil
}
