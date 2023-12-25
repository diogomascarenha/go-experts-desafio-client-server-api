package dtos

import "time"

type ExchangeRateOutputDTO struct {
	ID            int32
	Currency_From string
	Currency_To   string
	Value         float32
	CreatedAt     time.Time
}
