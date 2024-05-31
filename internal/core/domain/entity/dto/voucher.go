package dto

import (
	"fase-4-hf-orch/internal/core/domain/entity"
	vo "fase-4-hf-orch/internal/core/domain/entity/valueObject"
	"time"
)

type (
	RequestVoucher struct {
		UUID       string `json:"uuid,omitempty"`
		Code       string `json:"code,omitempty"`
		Percentage int64  `json:"percentage,omitempty"`
		CreatedAt  string `json:"createdAt,omitempty"`
		ExpiresAt  string `json:"expiresAt,omitempty"`
	}

	OutputVoucher struct {
		UUID       string `json:"uuid,omitempty"`
		Code       string `json:"code,omitempty"`
		Percentage int64  `json:"percentage,omitempty"`
		CreatedAt  string `json:"createdAt,omitempty"`
		ExpiresAt  string `json:"expiresAt,omitempty"`
	}
)

func (r RequestVoucher) Voucher() entity.Voucher {
	expirationTime, _ := time.Parse("02-01-2006 15:04:05", r.ExpiresAt)
	createdAt, _ := time.Parse("02-01-2006 15:04:05", r.CreatedAt)

	return entity.Voucher{
		Code:       r.Code,
		Percentage: r.Percentage,
		ExpiresAt: vo.ExpiresAt{
			Value: &expirationTime,
		},
		CreatedAt: vo.CreatedAt{
			Value: createdAt,
		},
	}
}
