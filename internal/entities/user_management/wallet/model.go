package wallet

import "gorm.io/gorm"

type Detail struct {
	gorm.Model
	ID       string `gorm:"type:varchar(20)"`
	UserId   string
	Balance  int  `gorm:"default:0"`
	Currency string	`gorm:"default:'INR'"`
}

const (
	TableName               ="wallet"
	EntityName              ="wallet"
	Prefix                  ="wal"
	IDLengthExcludingPrefix =14 //length used for alphanumeric string (exculding prefix
)
func(sd *Detail) TableName() string{
	return TableName
}
