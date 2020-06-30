package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
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
	defer func() {
		fmt.Println("Closing db..")
		db.Close()
	}()

	router := mux.NewRouter()
	if router == nil {
		log.Fatal("No router Created")
	}
	fmt.Println("Server Started")
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

	//log.Fatal(server.ListenAndServe())
	var wait time.Duration

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt)

	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	server.Shutdown(ctx)
	func() {
		fmt.Println("Closing DB")
		db.Close()
	}()
	fmt.Println("Server ShutDown....")

	os.Exit(0)
}

func chekErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
