package http

import (
	"context"
	"fase-4-hf-orch/internal/core/domain/entity/dto"
)

type PaymentAPI interface {
	DoPayment(ctx context.Context, input dto.InputPaymentAPI) (*dto.OutputPaymentAPI, error)
}
