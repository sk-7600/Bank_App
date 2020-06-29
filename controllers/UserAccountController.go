package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/sk-7600/Bank_App/BankApp/model"
	"github.com/sk-7600/Bank_App/BankApp/service"
)

type UserAccountController struct {
	uas *service.UserAccountService
}

func NewUserAccountController(uas *service.UserAccountService) *UserAccountController {
	return &UserAccountController{
		uas: uas,
	}
}

func (uac *UserAccountController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/bank/users/addUser", uac.NewUser).Methods("POST")
	router.HandleFunc("/bank/users/getAllUsers", uac.AllUsers).Methods("GET")
	router.HandleFunc("/bank/users/updateUser", uac.updateUser).Methods("PUT")
	//router.HandleFunc("/bank/users/deleteUser", uac.deleteUser).Methods("DELETE")
	router.HandleFunc("/bank/users/{ID}", uac.DeleteUserByID).Methods("DELETE")
	//router.HandleFunc("/bank/users/getUserByID", uac.UserByID).Methods("GET")
}

func (uac *UserAccountController) deleteUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := UnmarshalJSON(r, &user)
	if err != nil {
		var x = []byte(err.Error())
		w.Write(x)
	} else {
		var x = []byte("Data Deleted...")
		w.Write(x)
	}
	uac.uas.DeleteUserAccount(user)
	RespondJSON(&w, http.StatusOK, user)
}

func (uac *UserAccountController) updateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := UnmarshalJSON(r, &user)
	if err != nil {
		var x = []byte(err.Error())
		w.Write(x)
	} else {
		var x = []byte("Data Updated...")
		w.Write(x)
	}
	uac.uas.UpdateUserAccount(user)
	RespondJSON(&w, http.StatusOK, user)
}

func getIdFromRequest(req *http.Request) string {
	vars := mux.Vars(req)
	ID, _ := vars["ID"]
	return ID
}

func (uac *UserAccountController) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	id := getIdFromRequest(r)
	userid, _ := uuid.FromString(id)
	user := model.User{}
	user.ID = userid
	uac.uas.GetUserByID(&user)
	uac.uas.DeleteUserAccount(user)
	err := UnmarshalJSON(r, &user)
	if err != nil {
		var x = []byte(err.Error())
		w.Write(x)
	} else {
		var x = []byte("Data Deleted...")
		w.Write(x)
	}
}

func (uac *UserAccountController) AllUsers(w http.ResponseWriter, r *http.Request) {
	content := []model.User{}
	uac.uas.GetAllUsers(&content)
	//fmt.Println(content)
	RespondJSON(&w, http.StatusOK, content)
}

func (uac *UserAccountController) NewUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := UnmarshalJSON(r, &user)
	if err != nil {
		var x = []byte(err.Error())
		w.Write(x)
	} else {
		var x = []byte("Data Inserted...")
		w.Write(x)
	}
	uac.uas.AddUserAccount(user)

}

func UnmarshalJSON(r *http.Request, target interface{}) error {
	if r.Body == nil {
		return errors.New("There is problem while reading data")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.New("Can't handle data")
	}

	if len(body) == 0 {
		return errors.New("Empty Data")
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return errors.New("Unable to Parse Data")
	}
	return nil
}
