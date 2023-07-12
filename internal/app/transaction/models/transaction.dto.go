package models

import (
	"time"
)

type CreateTransactionDTO struct {
	UserID     string                    `json:"user_id"`
	SupplierID string                    `json:"supplier_id"`
	Data       []GetTransactionDetailDto `json:"data"`
}

type CreateTransactionDetailDTO struct {
	ProductID string `json:"product_id"`
	Qty       int64  `json:"qty"`
}

type GetTransactionDTO struct {
	ID              string                     `json:"id"`
	UserID          string                     `json:"user_id"`
	SupplierID      string                     `json:"supplier_id"`
	TransactionDate time.Time                  `json:"transaction_date"`
	Data            []*GetTransactionDetailDto `json:"data"`
}

type GetTransactionDetailDto struct {
	ID            string `json:"id"`
	TransactionID string `json:"transaction_id"`
	ProductID     string `json:"product_id"`
	Qty           int64  `json:"qty"`
}
