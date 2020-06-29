package service

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/sk-7600/Bank_App/BankApp/model"
	"github.com/sk-7600/Bank_App/BankApp/repository"
)

type UserAccountService struct {
	DB         *gorm.DB
	Repository *repository.GormRepository
}

func NewUserAccountService(db *gorm.DB, repo *repository.GormRepository) *UserAccountService {
	return &UserAccountService{
		DB:         db.AutoMigrate(&model.User{}),
		Repository: repo,
	}
}

//Add Account
func (uas *UserAccountService) AddUserAccount(user model.User) error {
	uow := repository.NewUnitOfWork(uas.DB, false)
	user.ID = uuid.NewV4()
	// for i := range user.BankAccounts {
	// 	//user.BankAccounts[i].UserID = uuid.Must(uuid.NewV4())
	// 	bac := BankAccountService{}
	// 	bac.AddBankAccount(user.BankAccounts[i])
	// 	user.BankAccounts[i].UserID = user.BankAccounts[i].UserID
	// }
	err := uas.Repository.Add(uow, user)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

//Get User Details
func (uas *UserAccountService) GetAllUsers(user *[]model.User) error {
	uow := repository.NewUnitOfWork(uas.DB, true)
	pA := []string{"BankAccounts"}
	err := uas.Repository.GetAll(uow, &user, pA)
	if err != nil {
		return err
	}
	return err
}

//Update User Account
func (uas *UserAccountService) UpdateUserAccount(user model.User) error {
	uow := repository.NewUnitOfWork(uas.DB, false)
	err := uas.Repository.Update(uow, user)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

//Delete User Account
func (uas *UserAccountService) DeleteUserAccount(user model.User) error {
	uow := repository.NewUnitOfWork(uas.DB, false)
	err := uas.Repository.Delete(uow, user)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

//Get User by ID
func (uas *UserAccountService) GetUserByID(user *model.User) error {
	uow := repository.NewUnitOfWork(uas.DB, true)
	pA := make([]string, 0)
	err := uas.Repository.Get(uow, user, user.ID, pA)
	if err != nil {
		return err
	}
	return err
}
