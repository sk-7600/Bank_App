package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sk-7600/Bank_App/BankApp/model"
	"github.com/sk-7600/Bank_App/BankApp/service"
	"github.com/sk-7600/Bank_App/BankApp/web"
)

type BankAccountController struct {
	bas *service.BankAccountService
}

func NewBankAccountController(bas *service.BankAccountService) *BankAccountController {
	return &BankAccountController{
		bas: bas,
	}
}

func (bac *BankAccountController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/bank/account/addUserAccount", bac.NewUserAccount).Methods("POST")
	router.HandleFunc("/bank/account/all", bac.GetAllAcount).Methods("GET")
}

func (bac *BankAccountController) NewUserAccount(w http.ResponseWriter, r *http.Request) {
	bA := model.BankAccount{}
	err := web.UnmarshalJSON(r, &bA)
	if err != nil {
		web.WriteErrorInResponse(err, w)
	} else {
		var x = []byte("Data Inserted...")
		w.Write(x)
	}
	er := bac.bas.AddBankAccount(bA)
	web.WriteErrorInResponse(er, w)
}

func (bac *BankAccountController) GetAllAcount(w http.ResponseWriter, r *http.Request) {
	content := []model.BankAccount{}
	er := bac.bas.GetAllData(&content)
	web.WriteErrorInResponse(er, w)
	//fmt.Println(content)
	web.RespondJSON(&w, http.StatusOK, content)
}
