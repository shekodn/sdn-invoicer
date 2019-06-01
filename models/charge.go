package models

import (
  "github.com/jinzhu/gorm"
)

type Charge struct {
	gorm.Model
	InvoiceID   int     `gorm:"index"  json:"invoice_id"`
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}
