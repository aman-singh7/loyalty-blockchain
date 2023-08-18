package main

import (
	"fmt"
	"log"

	"github.com/aman-singh7/loyalty-blockchain/infrastructure/rest/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	e := echo.New()

	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("configuration file is missing, err: %v\n", err)
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.ApplicationV1Router(e)

	PORT := viper.GetInt("Port")
	addr := fmt.Sprintf(":%d", PORT)

	// Start server
	e.Logger.Fatal(e.Start(addr))
}
