package service

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/sk-7600/Bank_App/BankApp/model"
	"github.com/sk-7600/Bank_App/BankApp/repository"
)

type BankAccountService struct {
	DB         *gorm.DB
	Repository *repository.GormRepository
}

func NewBankAccountService(db *gorm.DB, repo *repository.GormRepository) *BankAccountService {
	return &BankAccountService{
		DB:         db.AutoMigrate(&model.BankAccount{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE"),
		Repository: repo,
	}
}

//Add Account
func (bas *BankAccountService) AddBankAccount(bankAccount model.BankAccount) error {
	uow := repository.NewUnitOfWork(bas.DB, false)
	bankAccount.ID = uuid.NewV4()
	err := bas.Repository.Add(uow, bankAccount)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

//Delete Account
func (bas *BankAccountService) DeleteAccount(bankAccount model.BankAccount) error {
	uow := repository.NewUnitOfWork(bas.DB, false)
	err := bas.Repository.Delete(uow, bankAccount)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

//Update Account
func (bas *BankAccountService) UpdateAccount(bankAccount model.BankAccount) error {
	uow := repository.NewUnitOfWork(bas.DB, false)
	err := bas.Repository.Update(uow, bankAccount)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

//Get data by ID
func (bas *BankAccountService) GetByID(bankAccount *model.BankAccount) error {
	uow := repository.NewUnitOfWork(bas.DB, true)
	pA := make([]string, 0)
	err := bas.Repository.Get(uow, bankAccount, bankAccount.ID, pA)
	if err != nil {
		return err
	}
	return err
}

//Get All data
func (bas *BankAccountService) GetAllData(bA *[]model.BankAccount) error {
	uow := repository.NewUnitOfWork(bas.DB, true)
	pA := make([]string, 0)
	err := bas.Repository.GetAll(uow, &bA, pA)
	if err != nil {
		return err
	}
	return err
}
