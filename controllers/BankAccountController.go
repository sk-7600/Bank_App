package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sk-7600/Bank_App/BankApp/model"
	"github.com/sk-7600/Bank_App/BankApp/service"
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
	err := UnmarshalJSON(r, &bA)
	if err != nil {
		var x = []byte(err.Error())
		w.Write(x)
	} else {
		var x = []byte("Data Inserted...")
		w.Write(x)
	}
	er := bac.bas.AddBankAccount(bA)
	writeErrorInResponse(er, w)
}

func (bac *BankAccountController) GetAllAcount(w http.ResponseWriter, r *http.Request) {
	content := []model.BankAccount{}
	er := bac.bas.GetAllData(&content)
	writeErrorInResponse(er, w)
	//fmt.Println(content)
	RespondJSON(&w, http.StatusOK, content)
}

func RespondJSON(w *http.ResponseWriter, statusCode int, content interface{}) {
	response, err := json.Marshal(content)
	if err != nil {
		writeToHeader(w, http.StatusInternalServerError, err.Error())
		return
	}
	(*w).Header().Set("Content-Type", "application/json")
	writeToHeader(w, statusCode, response)
}

func writeToHeader(w *http.ResponseWriter, statusCode int, payload interface{}) {
	(*w).WriteHeader(statusCode)
	(*w).Write(payload.([]byte))
}

func writeErrorInResponse(err error, w http.ResponseWriter) {
	if err != nil {
		var x = []byte(err.Error())
		w.Write(x)
	}
}
