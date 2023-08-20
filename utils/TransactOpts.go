package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TransactOpts struct {
	client *ethclient.Client
}

func NewTransactOpts(client *ethclient.Client) *TransactOpts {
	return &TransactOpts{
		client: client,
	}
}

func (t *TransactOpts) GetOpts(accountAddr string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(accountAddr)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddr := crypto.PubkeyToAddress(*publicKeyECDSA)

	// fetch the last use nonce of account
	nonce, err := t.client.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("nounce: ", nonce)

	chainId, err := t.client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		panic(err)
	}

	gasPrice, err := t.client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("GasPrice: ", gasPrice)

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	return auth
}
