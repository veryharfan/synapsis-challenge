package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"synapsis-challenge/app/handler"
	"synapsis-challenge/app/middleware"
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

	productRepo := repository.InitProductRepository(db)
	productService := service.InitProductService(productRepo)
	productHandler := handler.InitProductHandler(productService)

	cartRepo := repository.InitCartRepository(db)
	cartService := service.InitCartService(cartRepo, productRepo)
	cartHandler := handler.InitCartHandler(cartService)

	r := gin.Default()
	r.POST("/register", customerHandler.Register)
	r.POST("/login", customerHandler.Login)

	r.GET("/products", productHandler.GetByCategory)

	authorized := r.Group("/")
	authorized.Use(middleware.Authentication(config.JWT))
	{
		authorized.POST("/carts", cartHandler.AddCart)
		authorized.GET("/carts", cartHandler.GetByCustomerId)
		authorized.DELETE("/carts/products/:productId", cartHandler.DeleteByCustomerIdAndProductId)
	}

	r.SetTrustedProxies(nil)
	r.Run(":" + config.Service.Port)
}
