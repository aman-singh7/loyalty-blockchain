package main

import (
	"fmt"
	"log"

	"github.com/aman-singh7/loyalty-blockchain/application/api"
	"github.com/aman-singh7/loyalty-blockchain/infrastructure/rest/routes"
	"github.com/aman-singh7/loyalty-blockchain/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-playground/validator/v10"
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

	// Validator
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	// Blockchain
	RPC_URL := viper.GetString("RPCUrl")
	PRIVATE_KEY := viper.GetString("PrivateKey")
	ADDRESS := viper.GetString("Address")

	var conn *api.Api

	if RPC_URL != "" && PRIVATE_KEY != "" && ADDRESS != "" {
		client, err := ethclient.Dial(RPC_URL)
		if err != nil {
			log.Fatalf("error while client intialization: %+v\n", err)
		}

		conn, err = api.NewApi(common.HexToAddress(ADDRESS), client)
		if err != nil {
			log.Fatalf("error while connection: %+v\n", err)
		}
	}

	routes.ApplicationV1Router(e, conn)

	PORT := viper.GetInt("Port")
	addr := fmt.Sprintf(":%d", PORT)

	// Start server
	e.Logger.Fatal(e.Start(addr))
}
