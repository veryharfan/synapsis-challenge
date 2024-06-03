package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"synapsis-challenge/app/handler"
	"synapsis-challenge/app/repository"
	"synapsis-challenge/app/service"
	"synapsis-challenge/config"
	"synapsis-challenge/migration"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	config := config.InitConfig(context.Background())

	args := os.Args
	if migrationArgs, runMigration := migration.IsRunMigration(args); runMigration {
		migration.RunMigration(migrationArgs, config.DB)
		return
	}

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
	r.Run(":" + config.Service.Port)
}
