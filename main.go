package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sk-7600/Bank_App/BankApp/controllers"
	"github.com/sk-7600/Bank_App/BankApp/repository"
	"github.com/sk-7600/Bank_App/BankApp/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/swabhav?charset=utf8&parseTime=True&loc=Local")
	chekErr(err)
	fmt.Println("Connection Establish...")
	defer db.Close()

	router := mux.NewRouter()
	nuac := service.NewUserAccountService(db, &repository.GormRepository{})
	ucon := controllers.NewUserAccountController(nuac)
	nbas := service.NewBankAccountService(db, &repository.GormRepository{})
	bcon := controllers.NewBankAccountController(nbas)
	bcon.RegisterRoutes(router)
	ucon.RegisterRoutes(router)

	headers := handlers.AllowedHeaders([]string{"Content-Type", "token"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origin := handlers.AllowedOrigins([]string{"*"})
	server := &http.Server{
		Handler:      handlers.CORS(headers, methods, origin)(router),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Addr:         ":8080",
	}

	log.Fatal(server.ListenAndServe())

	//Get All Details
	// data := []model.BankAccount{}
	// nbas.GetAllData(&data)
	// fmt.Println(data)

	// singleData := model.BankAccount{
	// 	CustomModel: model.CustomModel{
	// 		ID: uuid.Must(uuid.FromString("9917d25b-b4c7-482e-b9ff-742f57837955")),
	// 	},
	// }
	// nbas.GetByID(&singleData)
	// fmt.Println(singleData)

	// Update Account
	// data := model.BankAccount{
	// 	CustomModel: model.CustomModel{
	// 		ID: uuid.Must(uuid.FromString("af3dc1ad-2e9c-43e0-9ee2-d76ca1d37224")),
	// 	},
	// 	Name:    "Sumit",
	// 	Balance: 3000,
	// }
	// nbas.UpdateAccount(data)

	// Delete Account
	//nbas.DeleteAccount(model.BankAccount{})
}

func chekErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
