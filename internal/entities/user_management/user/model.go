package user

import "gorm.io/gorm"

type Detail struct {
	gorm.Model
	ID                   string `gorm:"type:varchar(20)"`
	Name                 string
	PhoneNumber          string
	Password             string
	Email                string
	NumberOfTransactions int64 `gorm:"default:0"`
	IsVerified bool `gorm:"default:false"`
}

const (
	TableName  = "user"
	EntityName = "user"
	Prefix     = "user"
	IDLengthExcludingPrefix= 14
)

func (sd *Detail) TableName() string {
	return TableName
}
