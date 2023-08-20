package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aman-singh7/loyalty-blockchain/application/api"
	"github.com/aman-singh7/loyalty-blockchain/infrastructure/rest/routes"
	"github.com/aman-singh7/loyalty-blockchain/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	CONTRACT_ADDRESS := viper.GetString("ContractAddress")

	if RPC_URL == "" || CONTRACT_ADDRESS == "" {
		log.Fatalln("url or address is empty")
	}

	client, err := ethclient.Dial(RPC_URL)
	if err != nil {
		log.Fatalf("error while client intialization: %+v\n", err)
	}

	contractAddr := common.HexToAddress(CONTRACT_ADDRESS)
	conn, err := api.NewApi(contractAddr, client)
	if err != nil {
		log.Fatalf("error while connection: %+v\n", err)
	}

	opts := utils.NewTransactOpts(client)
	srvc := api.NewService(conn, opts)

	routes.ApplicationV1Router(e, srvc)

	PORT := viper.GetInt("Port")
	addr := fmt.Sprintf(":%d", PORT)

	// Start server
	e.Logger.Fatal(e.Start(addr))

	// Logs
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog)
		}
	}
}
