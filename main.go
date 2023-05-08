package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type TFMContract struct {
	contractapi.Contract
}

func (c *TFMContract) Sum(ctx contractapi.TransactionContextInterface, value1Str string, value2Str string) (int, error) {
	value1, err := strconv.Atoi(value1Str)
	if err != nil {
		return 0, fmt.Errorf("failed to convert value1 to integer: %v", err)
	}

	value2, err := strconv.Atoi(value2Str)
	if err != nil {
		return 0, fmt.Errorf("failed to convert value2 to integer: %v", err)
	}

	return value1 + value2, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(&TFMContract{})
	if err != nil {
		fmt.Printf("Error creating chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting chaincode: %s", err.Error())
	}
}
