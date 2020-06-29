package model

import uuid "github.com/satori/go.uuid"

type BankAccount struct {
	BaseModel
	UserID uuid.UUID `gorm:"type:varchar(36)"`
	BName  string
}
