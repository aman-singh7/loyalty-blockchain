// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"couponId\",\"type\":\"uint256\"}],\"name\":\"CouponsExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"couponId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"InSufficientCoupons\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"InSufficientFunds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumLoyaltyProgramme.AccountType\",\"name\":\"validAuthority\",\"type\":\"uint8\"},{\"internalType\":\"enumLoyaltyProgramme.AccountType\",\"name\":\"recieverAuthority\",\"type\":\"uint8\"}],\"name\":\"InValidRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumLoyaltyProgramme.AccountType\",\"name\":\"existingAuthority\",\"type\":\"uint8\"}],\"name\":\"InValidRegistration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"couponId\",\"type\":\"uint256\"}],\"name\":\"InvalidCoupons\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumLoyaltyProgramme.AccountType\",\"name\":\"validAuthority\",\"type\":\"uint8\"},{\"internalType\":\"enumLoyaltyProgramme.AccountType\",\"name\":\"senderAuthority\",\"type\":\"uint8\"}],\"name\":\"UnAuthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"enumLoyaltyProgramme.AccountType\",\"name\":\"expectedType\",\"type\":\"uint8\"},{\"internalType\":\"enumLoyaltyProgramme.AccountType\",\"name\":\"recievedType\",\"type\":\"uint8\"}],\"name\":\"UnexpectedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"couponId\",\"type\":\"uint256\"}],\"name\":\"CouponCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"couponId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"CouponTransactionComplete\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"memberAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumLoyaltyProgramme.AccountType\",\"name\":\"authority\",\"type\":\"uint8\"}],\"name\":\"MemberRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SuperCoinTransactionComplete\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"consumerTokenPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"couponId\",\"type\":\"uint256\"}],\"name\":\"consumerCouponPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"businessAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"couponId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"purchaseCoupon\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redeemTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"couponId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"productCategory\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"thresholdvalue\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"productId\",\"type\":\"string\"},{\"internalType\":\"enumLoyaltyProgramme.CouponType\",\"name\":\"couponType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"createCoupons\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"memberAddress\",\"type\":\"address\"},{\"internalType\":\"enumLoyaltyProgramme.AccountType\",\"name\":\"accountType\",\"type\":\"uint8\"}],\"name\":\"registerMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"business\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"payBusiness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b503373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250505f805f60805173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f815f01819055506003816002015f6101000a81548160ff021916908360038111156100b4576100b36100bf565b5b0217905550506100ec565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b608051612a0c6200010c5f395f818161063901526113740152612a0c5ff3fe608060405234801561000f575f80fd5b5060043610610086575f3560e01c8063a6e158f811610059578063a6e158f81461010e578063ad466e2c1461012a578063ea747b8c14610146578063ec6976af1461016257610086565b8063398f332f1461008a57806343d24b23146100a65780637fdcffb8146100c257806397304ced146100de575b5f80fd5b6100a4600480360381019061009f9190611c54565b61017e565b005b6100c060048036038101906100bb9190611cc5565b61042f565b005b6100dc60048036038101906100d79190611d15565b610665565b005b6100f860048036038101906100f39190611d53565b610669565b6040516101059190611d8d565b60405180910390f35b61012860048036038101906101239190611d53565b610807565b005b610144600480360381019061013f9190611e2a565b61080a565b005b610160600480360381019061015b9190611f2f565b610c50565b005b61017c60048036038101906101779190611d15565b61126b565b005b600380600381111561019357610192611f93565b5b5f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff1660038111156101f1576101f0611f93565b5b1461028057805f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff166040517f4890faf0000000000000000000000000000000000000000000000000000000008152600401610277929190612006565b60405180910390fd5b5f600381111561029357610292611f93565b5b5f808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff1660038111156102f1576102f0611f93565b5b1461037e575f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff166040517ff2b5772b000000000000000000000000000000000000000000000000000000008152600401610375919061202d565b60405180910390fd5b5f805f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f815f018190555082816002015f6101000a81548160ff021916908360038111156103eb576103ea611f93565b5b02179055507f7496ec113494270625ae4bdb915ff6fdad6f6d96d4c88bd9417727daedcdf3e08484604051610421929190612055565b60405180910390a150505050565b600380600381111561044457610443611f93565b5b5f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff1660038111156104a2576104a1611f93565b5b1461053157805f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff166040517f4890faf0000000000000000000000000000000000000000000000000000000008152600401610528929190612006565b60405180910390fd5b6002600381111561054557610544611f93565b5b5f808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff1660038111156105a3576105a2611f93565b5b146106335760025f808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff166040517f86b15d9200000000000000000000000000000000000000000000000000000000815260040161062a929190612006565b60405180910390fd5b61065f847f0000000000000000000000000000000000000000000000000000000000000000858561139e565b50505050565b5050565b5f600380600381111561067f5761067e611f93565b5b5f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff1660038111156106dd576106dc611f93565b5b1461076c57805f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff166040517f4890faf0000000000000000000000000000000000000000000000000000000008152600401610763929190612006565b60405180910390fd5b825f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8282546107b991906120a9565b925050819055505f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0154915050919050565b50565b600280600381111561081f5761081e611f93565b5b5f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff16600381111561087d5761087c611f93565b5b1461090c57805f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff166040517f4890faf0000000000000000000000000000000000000000000000000000000008152600401610903929190612006565b60405180910390fd5b5f8b1161094e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109459061215c565b60405180910390fd5b60018081111561096157610960611f93565b5b83600181111561097457610973611f93565b5b036109c3575f85859050036109be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109b5906121ea565b60405180910390fd5b610a4c565b5f8888905003610a08576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ff9061229e565b60405180910390fd5b5f861015610a4b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a429061232c565b60405180910390fd5b5b6103e85f610a5a919061234a565b8910158015610a7757506103e86064610a73919061234a565b8911155b610ab6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aad906123fb565b60405180910390fd5b428211610af8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aef90612489565b60405180910390fd5b5f805f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8e81526020019081526020015f2090508a816001018190555033815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508b81600201819055508981600301819055508888826004019182610bb89291906126d8565b508681600501819055508585826006019182610bd59291906126d8565b5083816007015f6101000a81548160ff02191690836001811115610bfc57610bfb611f93565b5b02179055508281600801819055507fb7d0ca4a7a6c4b6645f44871097973843f7fc472371531680e0f8d6a4e4c7b108d604051610c399190611d8d565b60405180910390a150505050505050505050505050565b6001806003811115610c6557610c64611f93565b5b5f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff166003811115610cc357610cc2611f93565b5b14610d5257805f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff166040517f4890faf0000000000000000000000000000000000000000000000000000000008152600401610d49929190612006565b60405180910390fd5b836002806003811115610d6857610d67611f93565b5b5f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff166003811115610dc657610dc5611f93565b5b14610e575781815f808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff166040517fee07ec54000000000000000000000000000000000000000000000000000000008152600401610e4e939291906127a5565b60405180910390fd5b8585855f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8381526020019081526020015f206009015f9054906101000a900460ff16610ef957816040517ff3b77d2c000000000000000000000000000000000000000000000000000000008152600401610ef09190611d8d565b60405180910390fd5b805f808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8481526020019081526020015f20600201541015610fe45781815f808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8581526020019081526020015f20600201546040517f87c6d0cc000000000000000000000000000000000000000000000000000000008152600401610fdb939291906127da565b60405180910390fd5b88885f808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8281526020019081526020015f206009015f9054906101000a900460ff1661108557806040517ff3b77d2c00000000000000000000000000000000000000000000000000000000815260040161107c9190611d8d565b60405180910390fd5b425f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8381526020019081526020015f2060080154111561124f575f805f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8381526020019081526020015f206008015490505f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8381526020019081526020015f205f8082015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182015f9055600282015f9055600382015f9055600482015f6111cb9190611b77565b600582015f9055600682015f6111e19190611b77565b600782015f6101000a81549060ff0219169055600882015f9055600982015f6101000a81549060ff0219169055505080826040517ff9f0fd4500000000000000000000000000000000000000000000000000000000815260040161124692919061280f565b60405180910390fd5b61125c8c338d8d8d611550565b50505050505050505050505050565b60018060038111156112805761127f611f93565b5b5f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff1660038111156112de576112dd611f93565b5b1461136d57805f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff166040517f4890faf0000000000000000000000000000000000000000000000000000000008152600401611364929190612006565b60405180910390fd5b61139983337f00000000000000000000000000000000000000000000000000000000000000008561139e565b505050565b8281805f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0154101561146357805f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f01546040517f8fcd099700000000000000000000000000000000000000000000000000000000815260040161145a92919061280f565b60405180910390fd5b825f808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8282546114b09190612836565b92505081905550825f808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f82825461150491906120a9565b925050819055507f2e578e436685c8ee906431b002661f41d45a790b565ca7b6f06ae35bfae01a76868686866040516115409493929190612869565b60405180910390a1505050505050565b5f8483835f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8381526020019081526020015f206009015f9054906101000a900460ff166115f357816040517ff3b77d2c0000000000000000000000000000000000000000000000000000000081526004016115ea9190611d8d565b60405180910390fd5b805f808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8481526020019081526020015f206002015410156116de5781815f808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8581526020019081526020015f20600201546040517f87c6d0cc0000000000000000000000000000000000000000000000000000000081526004016116d5939291906127da565b60405180910390fd5b5f805f8a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8881526020019081526020015f2090505f805f8a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8981526020019081526020015f20604051806101400160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201548152602001600282015481526020016003820154815260200160048201805461180c9061250b565b80601f01602080910402602001604051908101604052809291908181526020018280546118389061250b565b80156118835780601f1061185a57610100808354040283529160200191611883565b820191905f5260205f20905b81548152906001019060200180831161186657829003601f168201915b50505050508152602001600582015481526020016006820180546118a69061250b565b80601f01602080910402602001604051908101604052809291908181526020018280546118d29061250b565b801561191d5780601f106118f45761010080835404028352916020019161191d565b820191905f5260205f20905b81548152906001019060200180831161190057829003601f168201915b50505050508152602001600782015f9054906101000a900460ff16600181111561194a57611949611f93565b5b600181111561195c5761195b611f93565b5b815260200160088201548152602001600982015f9054906101000a900460ff161515151581525050905086826002015f8282546119999190612836565b92505081905550806101200151156119c95786816040018181516119bd91906120a9565b915081815250506119d4565b868160400181815250505b805f808b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8a81526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155606082015181600301556080820151816004019081611a9a91906128b6565b5060a0820151816005015560c0820151816006019081611aba91906128b6565b5060e0820151816007015f6101000a81548160ff02191690836001811115611ae557611ae4611f93565b5b02179055506101008201518160080155610120820151816009015f6101000a81548160ff0219169083151502179055509050507fc19f12d44b95c1a221ab57e69ec538ff5181f820f28f967108426da285ab9a6b8b8b8b8b8b604051611b4f959493929190612985565b60405180910390a1816001015487611b67919061234a565b9550505050505095945050505050565b508054611b839061250b565b5f825580601f10611b945750611bb1565b601f0160209004905f5260205f2090810190611bb09190611bb4565b5b50565b5b80821115611bcb575f815f905550600101611bb5565b5090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611c0082611bd7565b9050919050565b611c1081611bf6565b8114611c1a575f80fd5b50565b5f81359050611c2b81611c07565b92915050565b60048110611c3d575f80fd5b50565b5f81359050611c4e81611c31565b92915050565b5f8060408385031215611c6a57611c69611bcf565b5b5f611c7785828601611c1d565b9250506020611c8885828601611c40565b9150509250929050565b5f819050919050565b611ca481611c92565b8114611cae575f80fd5b50565b5f81359050611cbf81611c9b565b92915050565b5f805f60608486031215611cdc57611cdb611bcf565b5b5f611ce986828701611cb1565b9350506020611cfa86828701611c1d565b9250506040611d0b86828701611cb1565b9150509250925092565b5f8060408385031215611d2b57611d2a611bcf565b5b5f611d3885828601611cb1565b9250506020611d4985828601611cb1565b9150509250929050565b5f60208284031215611d6857611d67611bcf565b5b5f611d7584828501611cb1565b91505092915050565b611d8781611c92565b82525050565b5f602082019050611da05f830184611d7e565b92915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f840112611dc757611dc6611da6565b5b8235905067ffffffffffffffff811115611de457611de3611daa565b5b602083019150836001820283011115611e0057611dff611dae565b5b9250929050565b60028110611e13575f80fd5b50565b5f81359050611e2481611e07565b92915050565b5f805f805f805f805f805f6101208c8e031215611e4a57611e49611bcf565b5b5f611e578e828f01611cb1565b9b50506020611e688e828f01611cb1565b9a50506040611e798e828f01611cb1565b9950506060611e8a8e828f01611cb1565b98505060808c013567ffffffffffffffff811115611eab57611eaa611bd3565b5b611eb78e828f01611db2565b975097505060a0611eca8e828f01611cb1565b95505060c08c013567ffffffffffffffff811115611eeb57611eea611bd3565b5b611ef78e828f01611db2565b945094505060e0611f0a8e828f01611e16565b925050610100611f1c8e828f01611cb1565b9150509295989b509295989b9093969950565b5f805f8060808587031215611f4757611f46611bcf565b5b5f611f5487828801611cb1565b9450506020611f6587828801611c1d565b9350506040611f7687828801611cb1565b9250506060611f8787828801611cb1565b91505092959194509250565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60048110611fd157611fd0611f93565b5b50565b5f819050611fe182611fc0565b919050565b5f611ff082611fd4565b9050919050565b61200081611fe6565b82525050565b5f6040820190506120195f830185611ff7565b6120266020830184611ff7565b9392505050565b5f6020820190506120405f830184611ff7565b92915050565b61204f81611bf6565b82525050565b5f6040820190506120685f830185612046565b6120756020830184611ff7565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6120b382611c92565b91506120be83611c92565b92508282019050808211156120d6576120d561207c565b5b92915050565b5f82825260208201905092915050565b7f636f756e74206f6620746f6b656e732073686f756c64206265206772656174655f8201527f72207468616e2030000000000000000000000000000000000000000000000000602082015250565b5f6121466028836120dc565b9150612151826120ec565b604082019050919050565b5f6020820190508181035f8301526121738161213a565b9050919050565b7f70726f6475637449642073686f756c64206e6f7420626520656d7074792077685f8201527f656e20636f75706f6e207479706520697320756e697175652070726f64756374602082015250565b5f6121d46040836120dc565b91506121df8261217a565b604082019050919050565b5f6020820190508181035f830152612201816121c8565b9050919050565b7f70726f647563742063617465676f72792073686f756c64206e6f7420626520655f8201527f6d707479207768656e20636f75706f6e20747970652069732063617465676f7260208201527f7900000000000000000000000000000000000000000000000000000000000000604082015250565b5f6122886041836120dc565b915061229382612208565b606082019050919050565b5f6020820190508181035f8301526122b58161227c565b9050919050565b7f7468726573686f6c642076616c756520666f72207468652070726f64756374205f8201527f73686f756c6420626520616c776179732067726561746572207468616e203000602082015250565b5f612316603f836120dc565b9150612321826122bc565b604082019050919050565b5f6020820190508181035f8301526123438161230a565b9050919050565b5f61235482611c92565b915061235f83611c92565b925082820261236d81611c92565b915082820484148315176123845761238361207c565b5b5092915050565b7f646973636f756e742073686f756c642062652076616c696420282062657477655f8201527f656e203020746f20313030202900000000000000000000000000000000000000602082015250565b5f6123e5602d836120dc565b91506123f08261238b565b604082019050919050565b5f6020820190508181035f830152612412816123d9565b9050919050565b7f6578706972792064617465206f662074686520636f75706f6e73206d757374205f8201527f6265206166746572207468652000000000000000000000000000000000000000602082015250565b5f612473602d836120dc565b915061247e82612419565b604082019050919050565b5f6020820190508181035f8301526124a081612467565b9050919050565b5f82905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061252257607f821691505b602082108103612535576125346124de565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026125977fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261255c565b6125a1868361255c565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6125dc6125d76125d284611c92565b6125b9565b611c92565b9050919050565b5f819050919050565b6125f5836125c2565b612609612601826125e3565b848454612568565b825550505050565b5f90565b61261d612611565b6126288184846125ec565b505050565b5b8181101561264b576126405f82612615565b60018101905061262e565b5050565b601f821115612690576126618161253b565b61266a8461254d565b81016020851015612679578190505b61268d6126858561254d565b83018261262d565b50505b505050565b5f82821c905092915050565b5f6126b05f1984600802612695565b1980831691505092915050565b5f6126c883836126a1565b9150826002028217905092915050565b6126e283836124a7565b67ffffffffffffffff8111156126fb576126fa6124b1565b5b612705825461250b565b61271082828561264f565b5f601f83116001811461273d575f841561272b578287013590505b61273585826126bd565b86555061279c565b601f19841661274b8661253b565b5f5b828110156127725784890135825560018201915060208501945060208101905061274d565b8683101561278f578489013561278b601f8916826126a1565b8355505b6001600288020188555050505b50505050505050565b5f6060820190506127b85f830186612046565b6127c56020830185611ff7565b6127d26040830184611ff7565b949350505050565b5f6060820190506127ed5f830186611d7e565b6127fa6020830185611d7e565b6128076040830184611d7e565b949350505050565b5f6040820190506128225f830185611d7e565b61282f6020830184611d7e565b9392505050565b5f61284082611c92565b915061284b83611c92565b92508282039050818111156128635761286261207c565b5b92915050565b5f60808201905061287c5f830187611d7e565b6128896020830186612046565b6128966040830185612046565b6128a36060830184611d7e565b95945050505050565b5f81519050919050565b6128bf826128ac565b67ffffffffffffffff8111156128d8576128d76124b1565b5b6128e2825461250b565b6128ed82828561264f565b5f60209050601f83116001811461291e575f841561290c578287015190505b61291685826126bd565b86555061297d565b601f19841661292c8661253b565b5f5b828110156129535784890151825560018201915060208501945060208101905061292e565b86831015612970578489015161296c601f8916826126a1565b8355505b6001600288020188555050505b505050505050565b5f60a0820190506129985f830188611d7e565b6129a56020830187612046565b6129b26040830186612046565b6129bf6060830185611d7e565b6129cc6080830184611d7e565b969550505050505056fea26469706673582212206c8896bfe339adeeb30fdb5e5c2d8ef01fb9f8712f93269b1b578729b73dfab364736f6c63430008150033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// ConsumerCouponPayment is a paid mutator transaction binding the contract method 0x7fdcffb8.
//
// Solidity: function consumerCouponPayment(uint256 transactionId, uint256 couponId) returns()
func (_Api *ApiTransactor) ConsumerCouponPayment(opts *bind.TransactOpts, transactionId *big.Int, couponId *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "consumerCouponPayment", transactionId, couponId)
}

// ConsumerCouponPayment is a paid mutator transaction binding the contract method 0x7fdcffb8.
//
// Solidity: function consumerCouponPayment(uint256 transactionId, uint256 couponId) returns()
func (_Api *ApiSession) ConsumerCouponPayment(transactionId *big.Int, couponId *big.Int) (*types.Transaction, error) {
	return _Api.Contract.ConsumerCouponPayment(&_Api.TransactOpts, transactionId, couponId)
}

// ConsumerCouponPayment is a paid mutator transaction binding the contract method 0x7fdcffb8.
//
// Solidity: function consumerCouponPayment(uint256 transactionId, uint256 couponId) returns()
func (_Api *ApiTransactorSession) ConsumerCouponPayment(transactionId *big.Int, couponId *big.Int) (*types.Transaction, error) {
	return _Api.Contract.ConsumerCouponPayment(&_Api.TransactOpts, transactionId, couponId)
}

// ConsumerTokenPayment is a paid mutator transaction binding the contract method 0xec6976af.
//
// Solidity: function consumerTokenPayment(uint256 transactionId, uint256 amount) returns()
func (_Api *ApiTransactor) ConsumerTokenPayment(opts *bind.TransactOpts, transactionId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "consumerTokenPayment", transactionId, amount)
}

// ConsumerTokenPayment is a paid mutator transaction binding the contract method 0xec6976af.
//
// Solidity: function consumerTokenPayment(uint256 transactionId, uint256 amount) returns()
func (_Api *ApiSession) ConsumerTokenPayment(transactionId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.ConsumerTokenPayment(&_Api.TransactOpts, transactionId, amount)
}

// ConsumerTokenPayment is a paid mutator transaction binding the contract method 0xec6976af.
//
// Solidity: function consumerTokenPayment(uint256 transactionId, uint256 amount) returns()
func (_Api *ApiTransactorSession) ConsumerTokenPayment(transactionId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.ConsumerTokenPayment(&_Api.TransactOpts, transactionId, amount)
}

// CreateCoupons is a paid mutator transaction binding the contract method 0xad466e2c.
//
// Solidity: function createCoupons(uint256 couponId, uint256 count, uint256 cost, uint256 discount, string productCategory, uint256 thresholdvalue, string productId, uint8 couponType, uint256 expiry) returns()
func (_Api *ApiTransactor) CreateCoupons(opts *bind.TransactOpts, couponId *big.Int, count *big.Int, cost *big.Int, discount *big.Int, productCategory string, thresholdvalue *big.Int, productId string, couponType uint8, expiry *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "createCoupons", couponId, count, cost, discount, productCategory, thresholdvalue, productId, couponType, expiry)
}

// CreateCoupons is a paid mutator transaction binding the contract method 0xad466e2c.
//
// Solidity: function createCoupons(uint256 couponId, uint256 count, uint256 cost, uint256 discount, string productCategory, uint256 thresholdvalue, string productId, uint8 couponType, uint256 expiry) returns()
func (_Api *ApiSession) CreateCoupons(couponId *big.Int, count *big.Int, cost *big.Int, discount *big.Int, productCategory string, thresholdvalue *big.Int, productId string, couponType uint8, expiry *big.Int) (*types.Transaction, error) {
	return _Api.Contract.CreateCoupons(&_Api.TransactOpts, couponId, count, cost, discount, productCategory, thresholdvalue, productId, couponType, expiry)
}

// CreateCoupons is a paid mutator transaction binding the contract method 0xad466e2c.
//
// Solidity: function createCoupons(uint256 couponId, uint256 count, uint256 cost, uint256 discount, string productCategory, uint256 thresholdvalue, string productId, uint8 couponType, uint256 expiry) returns()
func (_Api *ApiTransactorSession) CreateCoupons(couponId *big.Int, count *big.Int, cost *big.Int, discount *big.Int, productCategory string, thresholdvalue *big.Int, productId string, couponType uint8, expiry *big.Int) (*types.Transaction, error) {
	return _Api.Contract.CreateCoupons(&_Api.TransactOpts, couponId, count, cost, discount, productCategory, thresholdvalue, productId, couponType, expiry)
}

// MintTokens is a paid mutator transaction binding the contract method 0x97304ced.
//
// Solidity: function mintTokens(uint256 amount) returns(uint256)
func (_Api *ApiTransactor) MintTokens(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "mintTokens", amount)
}

// MintTokens is a paid mutator transaction binding the contract method 0x97304ced.
//
// Solidity: function mintTokens(uint256 amount) returns(uint256)
func (_Api *ApiSession) MintTokens(amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.MintTokens(&_Api.TransactOpts, amount)
}

// MintTokens is a paid mutator transaction binding the contract method 0x97304ced.
//
// Solidity: function mintTokens(uint256 amount) returns(uint256)
func (_Api *ApiTransactorSession) MintTokens(amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.MintTokens(&_Api.TransactOpts, amount)
}

// PayBusiness is a paid mutator transaction binding the contract method 0x43d24b23.
//
// Solidity: function payBusiness(uint256 transactionId, address business, uint256 amount) returns()
func (_Api *ApiTransactor) PayBusiness(opts *bind.TransactOpts, transactionId *big.Int, business common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "payBusiness", transactionId, business, amount)
}

// PayBusiness is a paid mutator transaction binding the contract method 0x43d24b23.
//
// Solidity: function payBusiness(uint256 transactionId, address business, uint256 amount) returns()
func (_Api *ApiSession) PayBusiness(transactionId *big.Int, business common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.PayBusiness(&_Api.TransactOpts, transactionId, business, amount)
}

// PayBusiness is a paid mutator transaction binding the contract method 0x43d24b23.
//
// Solidity: function payBusiness(uint256 transactionId, address business, uint256 amount) returns()
func (_Api *ApiTransactorSession) PayBusiness(transactionId *big.Int, business common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.PayBusiness(&_Api.TransactOpts, transactionId, business, amount)
}

// PurchaseCoupon is a paid mutator transaction binding the contract method 0xea747b8c.
//
// Solidity: function purchaseCoupon(uint256 transactionId, address businessAddress, uint256 couponId, uint256 count) returns()
func (_Api *ApiTransactor) PurchaseCoupon(opts *bind.TransactOpts, transactionId *big.Int, businessAddress common.Address, couponId *big.Int, count *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "purchaseCoupon", transactionId, businessAddress, couponId, count)
}

// PurchaseCoupon is a paid mutator transaction binding the contract method 0xea747b8c.
//
// Solidity: function purchaseCoupon(uint256 transactionId, address businessAddress, uint256 couponId, uint256 count) returns()
func (_Api *ApiSession) PurchaseCoupon(transactionId *big.Int, businessAddress common.Address, couponId *big.Int, count *big.Int) (*types.Transaction, error) {
	return _Api.Contract.PurchaseCoupon(&_Api.TransactOpts, transactionId, businessAddress, couponId, count)
}

// PurchaseCoupon is a paid mutator transaction binding the contract method 0xea747b8c.
//
// Solidity: function purchaseCoupon(uint256 transactionId, address businessAddress, uint256 couponId, uint256 count) returns()
func (_Api *ApiTransactorSession) PurchaseCoupon(transactionId *big.Int, businessAddress common.Address, couponId *big.Int, count *big.Int) (*types.Transaction, error) {
	return _Api.Contract.PurchaseCoupon(&_Api.TransactOpts, transactionId, businessAddress, couponId, count)
}

// RedeemTokens is a paid mutator transaction binding the contract method 0xa6e158f8.
//
// Solidity: function redeemTokens(uint256 amount) returns()
func (_Api *ApiTransactor) RedeemTokens(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "redeemTokens", amount)
}

// RedeemTokens is a paid mutator transaction binding the contract method 0xa6e158f8.
//
// Solidity: function redeemTokens(uint256 amount) returns()
func (_Api *ApiSession) RedeemTokens(amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.RedeemTokens(&_Api.TransactOpts, amount)
}

// RedeemTokens is a paid mutator transaction binding the contract method 0xa6e158f8.
//
// Solidity: function redeemTokens(uint256 amount) returns()
func (_Api *ApiTransactorSession) RedeemTokens(amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.RedeemTokens(&_Api.TransactOpts, amount)
}

// RegisterMember is a paid mutator transaction binding the contract method 0x398f332f.
//
// Solidity: function registerMember(address memberAddress, uint8 accountType) returns()
func (_Api *ApiTransactor) RegisterMember(opts *bind.TransactOpts, memberAddress common.Address, accountType uint8) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "registerMember", memberAddress, accountType)
}

// RegisterMember is a paid mutator transaction binding the contract method 0x398f332f.
//
// Solidity: function registerMember(address memberAddress, uint8 accountType) returns()
func (_Api *ApiSession) RegisterMember(memberAddress common.Address, accountType uint8) (*types.Transaction, error) {
	return _Api.Contract.RegisterMember(&_Api.TransactOpts, memberAddress, accountType)
}

// RegisterMember is a paid mutator transaction binding the contract method 0x398f332f.
//
// Solidity: function registerMember(address memberAddress, uint8 accountType) returns()
func (_Api *ApiTransactorSession) RegisterMember(memberAddress common.Address, accountType uint8) (*types.Transaction, error) {
	return _Api.Contract.RegisterMember(&_Api.TransactOpts, memberAddress, accountType)
}

// ApiCouponCreatedIterator is returned from FilterCouponCreated and is used to iterate over the raw logs and unpacked data for CouponCreated events raised by the Api contract.
type ApiCouponCreatedIterator struct {
	Event *ApiCouponCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiCouponCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiCouponCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiCouponCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiCouponCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiCouponCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiCouponCreated represents a CouponCreated event raised by the Api contract.
type ApiCouponCreated struct {
	CouponId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCouponCreated is a free log retrieval operation binding the contract event 0xb7d0ca4a7a6c4b6645f44871097973843f7fc472371531680e0f8d6a4e4c7b10.
//
// Solidity: event CouponCreated(uint256 couponId)
func (_Api *ApiFilterer) FilterCouponCreated(opts *bind.FilterOpts) (*ApiCouponCreatedIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "CouponCreated")
	if err != nil {
		return nil, err
	}
	return &ApiCouponCreatedIterator{contract: _Api.contract, event: "CouponCreated", logs: logs, sub: sub}, nil
}

// WatchCouponCreated is a free log subscription operation binding the contract event 0xb7d0ca4a7a6c4b6645f44871097973843f7fc472371531680e0f8d6a4e4c7b10.
//
// Solidity: event CouponCreated(uint256 couponId)
func (_Api *ApiFilterer) WatchCouponCreated(opts *bind.WatchOpts, sink chan<- *ApiCouponCreated) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "CouponCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiCouponCreated)
				if err := _Api.contract.UnpackLog(event, "CouponCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCouponCreated is a log parse operation binding the contract event 0xb7d0ca4a7a6c4b6645f44871097973843f7fc472371531680e0f8d6a4e4c7b10.
//
// Solidity: event CouponCreated(uint256 couponId)
func (_Api *ApiFilterer) ParseCouponCreated(log types.Log) (*ApiCouponCreated, error) {
	event := new(ApiCouponCreated)
	if err := _Api.contract.UnpackLog(event, "CouponCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiCouponTransactionCompleteIterator is returned from FilterCouponTransactionComplete and is used to iterate over the raw logs and unpacked data for CouponTransactionComplete events raised by the Api contract.
type ApiCouponTransactionCompleteIterator struct {
	Event *ApiCouponTransactionComplete // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiCouponTransactionCompleteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiCouponTransactionComplete)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiCouponTransactionComplete)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiCouponTransactionCompleteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiCouponTransactionCompleteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiCouponTransactionComplete represents a CouponTransactionComplete event raised by the Api contract.
type ApiCouponTransactionComplete struct {
	TransactionId *big.Int
	Sender        common.Address
	Receiver      common.Address
	CouponId      *big.Int
	Count         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCouponTransactionComplete is a free log retrieval operation binding the contract event 0xc19f12d44b95c1a221ab57e69ec538ff5181f820f28f967108426da285ab9a6b.
//
// Solidity: event CouponTransactionComplete(uint256 transactionId, address sender, address receiver, uint256 couponId, uint256 count)
func (_Api *ApiFilterer) FilterCouponTransactionComplete(opts *bind.FilterOpts) (*ApiCouponTransactionCompleteIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "CouponTransactionComplete")
	if err != nil {
		return nil, err
	}
	return &ApiCouponTransactionCompleteIterator{contract: _Api.contract, event: "CouponTransactionComplete", logs: logs, sub: sub}, nil
}

// WatchCouponTransactionComplete is a free log subscription operation binding the contract event 0xc19f12d44b95c1a221ab57e69ec538ff5181f820f28f967108426da285ab9a6b.
//
// Solidity: event CouponTransactionComplete(uint256 transactionId, address sender, address receiver, uint256 couponId, uint256 count)
func (_Api *ApiFilterer) WatchCouponTransactionComplete(opts *bind.WatchOpts, sink chan<- *ApiCouponTransactionComplete) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "CouponTransactionComplete")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiCouponTransactionComplete)
				if err := _Api.contract.UnpackLog(event, "CouponTransactionComplete", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCouponTransactionComplete is a log parse operation binding the contract event 0xc19f12d44b95c1a221ab57e69ec538ff5181f820f28f967108426da285ab9a6b.
//
// Solidity: event CouponTransactionComplete(uint256 transactionId, address sender, address receiver, uint256 couponId, uint256 count)
func (_Api *ApiFilterer) ParseCouponTransactionComplete(log types.Log) (*ApiCouponTransactionComplete, error) {
	event := new(ApiCouponTransactionComplete)
	if err := _Api.contract.UnpackLog(event, "CouponTransactionComplete", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiMemberRegisteredIterator is returned from FilterMemberRegistered and is used to iterate over the raw logs and unpacked data for MemberRegistered events raised by the Api contract.
type ApiMemberRegisteredIterator struct {
	Event *ApiMemberRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiMemberRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiMemberRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiMemberRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiMemberRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiMemberRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiMemberRegistered represents a MemberRegistered event raised by the Api contract.
type ApiMemberRegistered struct {
	MemberAddress common.Address
	Authority     uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMemberRegistered is a free log retrieval operation binding the contract event 0x7496ec113494270625ae4bdb915ff6fdad6f6d96d4c88bd9417727daedcdf3e0.
//
// Solidity: event MemberRegistered(address memberAddress, uint8 authority)
func (_Api *ApiFilterer) FilterMemberRegistered(opts *bind.FilterOpts) (*ApiMemberRegisteredIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "MemberRegistered")
	if err != nil {
		return nil, err
	}
	return &ApiMemberRegisteredIterator{contract: _Api.contract, event: "MemberRegistered", logs: logs, sub: sub}, nil
}

// WatchMemberRegistered is a free log subscription operation binding the contract event 0x7496ec113494270625ae4bdb915ff6fdad6f6d96d4c88bd9417727daedcdf3e0.
//
// Solidity: event MemberRegistered(address memberAddress, uint8 authority)
func (_Api *ApiFilterer) WatchMemberRegistered(opts *bind.WatchOpts, sink chan<- *ApiMemberRegistered) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "MemberRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiMemberRegistered)
				if err := _Api.contract.UnpackLog(event, "MemberRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMemberRegistered is a log parse operation binding the contract event 0x7496ec113494270625ae4bdb915ff6fdad6f6d96d4c88bd9417727daedcdf3e0.
//
// Solidity: event MemberRegistered(address memberAddress, uint8 authority)
func (_Api *ApiFilterer) ParseMemberRegistered(log types.Log) (*ApiMemberRegistered, error) {
	event := new(ApiMemberRegistered)
	if err := _Api.contract.UnpackLog(event, "MemberRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiSuperCoinTransactionCompleteIterator is returned from FilterSuperCoinTransactionComplete and is used to iterate over the raw logs and unpacked data for SuperCoinTransactionComplete events raised by the Api contract.
type ApiSuperCoinTransactionCompleteIterator struct {
	Event *ApiSuperCoinTransactionComplete // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiSuperCoinTransactionCompleteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiSuperCoinTransactionComplete)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiSuperCoinTransactionComplete)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiSuperCoinTransactionCompleteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiSuperCoinTransactionCompleteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiSuperCoinTransactionComplete represents a SuperCoinTransactionComplete event raised by the Api contract.
type ApiSuperCoinTransactionComplete struct {
	TransactionId *big.Int
	Sender        common.Address
	Receiver      common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSuperCoinTransactionComplete is a free log retrieval operation binding the contract event 0x2e578e436685c8ee906431b002661f41d45a790b565ca7b6f06ae35bfae01a76.
//
// Solidity: event SuperCoinTransactionComplete(uint256 transactionId, address sender, address receiver, uint256 amount)
func (_Api *ApiFilterer) FilterSuperCoinTransactionComplete(opts *bind.FilterOpts) (*ApiSuperCoinTransactionCompleteIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "SuperCoinTransactionComplete")
	if err != nil {
		return nil, err
	}
	return &ApiSuperCoinTransactionCompleteIterator{contract: _Api.contract, event: "SuperCoinTransactionComplete", logs: logs, sub: sub}, nil
}

// WatchSuperCoinTransactionComplete is a free log subscription operation binding the contract event 0x2e578e436685c8ee906431b002661f41d45a790b565ca7b6f06ae35bfae01a76.
//
// Solidity: event SuperCoinTransactionComplete(uint256 transactionId, address sender, address receiver, uint256 amount)
func (_Api *ApiFilterer) WatchSuperCoinTransactionComplete(opts *bind.WatchOpts, sink chan<- *ApiSuperCoinTransactionComplete) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "SuperCoinTransactionComplete")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiSuperCoinTransactionComplete)
				if err := _Api.contract.UnpackLog(event, "SuperCoinTransactionComplete", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSuperCoinTransactionComplete is a log parse operation binding the contract event 0x2e578e436685c8ee906431b002661f41d45a790b565ca7b6f06ae35bfae01a76.
//
// Solidity: event SuperCoinTransactionComplete(uint256 transactionId, address sender, address receiver, uint256 amount)
func (_Api *ApiFilterer) ParseSuperCoinTransactionComplete(log types.Log) (*ApiSuperCoinTransactionComplete, error) {
	event := new(ApiSuperCoinTransactionComplete)
	if err := _Api.contract.UnpackLog(event, "SuperCoinTransactionComplete", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
