package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
type Asset struct {
	ID               string  `json:"ID"`
	currentTime      string  `json:"currentTime"`
	nextBidTime      string  `json:"nextBidTime"`
	P_up_available   float64 `json:"P_up_available"`
	P_down_available float64 `json:"P_down_available"`
	price_up_bid     float64 `json:"price_up_bid"`
	price_down_bid   float64 `json:"price_down_bid"`
}
