package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/sk-7600/Bank_App/BankApp/model"
	"github.com/sk-7600/Bank_App/BankApp/service"
	"github.com/sk-7600/Bank_App/BankApp/web"
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
	err := web.UnmarshalJSON(r, &user)
	if err != nil {
		web.WriteErrorInResponse(err, w)
	} else {
		var x = []byte("Data Deleted...")
		w.Write(x)
	}
	er := uac.uas.DeleteUserAccount(user)
	web.WriteErrorInResponse(er, w)
	web.RespondJSON(&w, http.StatusOK, user)
}

func (uac *UserAccountController) updateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := web.UnmarshalJSON(r, &user)
	if err != nil {
		web.WriteErrorInResponse(err, w)
	} else {
		var x = []byte("Data Updated...")
		w.Write(x)
	}
	er := uac.uas.UpdateUserAccount(user)
	web.WriteErrorInResponse(er, w)
	web.RespondJSON(&w, http.StatusOK, user)
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
	er1 := uac.uas.GetUserByID(&user)
	web.WriteErrorInResponse(er1, w)
	er2 := uac.uas.DeleteUserAccount(user)
	web.WriteErrorInResponse(er2, w)
	err := web.UnmarshalJSON(r, &user)
	if err != nil {
		web.WriteErrorInResponse(err, w)
	} else {
		var x = []byte("Data Deleted...")
		w.Write(x)
	}
}

func (uac *UserAccountController) AllUsers(w http.ResponseWriter, r *http.Request) {
	content := []model.User{}
	er := uac.uas.GetAllUsers(&content)
	web.WriteErrorInResponse(er, w)
	//fmt.Println(content)
	web.RespondJSON(&w, http.StatusOK, content)
}

func (uac *UserAccountController) NewUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := web.UnmarshalJSON(r, &user)
	if err != nil {
		web.WriteErrorInResponse(err, w)
	} else {
		var x = []byte("Data Inserted...")
		w.Write(x)
	}
	er := uac.uas.AddUserAccount(user)
	web.WriteErrorInResponse(er, w)

}
