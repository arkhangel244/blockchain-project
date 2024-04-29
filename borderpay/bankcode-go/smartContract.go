/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"

	banking "bankcode-go/contracts"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	auctionSmartContract, err := contractapi.NewChaincode(&banking.PaymentContract{})
	if err != nil {
		log.Panicf("Error creating auction chaincode: %v", err)
	}

	if err := auctionSmartContract.Start(); err != nil {
		log.Panicf("Error starting auction chaincode: %v", err)
	}
}
