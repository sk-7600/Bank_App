package model

type User struct {
	BaseModel
	UName        string
	BankAccounts []BankAccount
}
