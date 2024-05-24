package entity

import (
	vo "fase-4-hf-orch/internal/core/domain/entity/valueObject"
)

type Voucher struct {
	ID         string       `json:"id,omitempty"`
	Code       string       `json:"code,omitempty"`
	Percentage int64        `json:"percentage,omitempty"`
	CreatedAt  vo.CreatedAt `json:"createdAt,omitempty"`
	ExpiresAt  vo.ExpiresAt `json:"expiresAt,omitempty"`
}
