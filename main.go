package main

import (
	"context"
	"database/sql"
	"log"
	"synapsis-challenge/config"
	"synapsis-challenge/src/handler"
	"synapsis-challenge/src/repository"
	"synapsis-challenge/src/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	config := config.InitConfig(context.Background())

	db, err := sql.Open("postgres", config.DB.ConnUri)
	if err != nil {
		log.Fatal(err)
	}

	customerRepository := repository.InitCustomerRepository(db)
	customerService := service.InitCustomerService(customerRepository, config.JWT)
	customerHandler := handler.InitCustomerHandler(customerService)

	r := gin.Default()
	r.POST("/register", customerHandler.Register)
	r.POST("/login", customerHandler.Login)

	r.SetTrustedProxies(nil)
	r.Run(":8080")
}
