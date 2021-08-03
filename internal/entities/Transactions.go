package entities

import "github.com/jinzhu/gorm"

type Transactions struct {
	gorm.Model
	Id string `json:"id",gorm:"primaryKey"`
	Amount uint `json:"amount"`
	Currency string `json:"currency"`
	SourceId string `json:"source_id"`
	DestinationId string `json:"destination_id"`
	Type string `json:"type"`
	Status     string `json:"status"`
	PaymentId string `json:"payment_id"`
	Payments Payments
}

func (b *Transactions) TableName() string {
	return "Transactions"
}
//added a comment
