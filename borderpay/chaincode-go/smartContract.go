/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"

	borderpay "chaincode-go/contracts"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	auctionSmartContract, err := contractapi.NewChaincode(&borderpay.HiringContract{})
	if err != nil {
		log.Panicf("Error creating auction chaincode: %v", err)
	}

	if err := auctionSmartContract.Start(); err != nil {
		log.Panicf("Error starting auction chaincode: %v", err)
	}
}
