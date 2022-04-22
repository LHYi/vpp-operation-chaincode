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
	ID                   string  `json:"ID"`
	Regulation_direction string  `json:"Regulation_direction"`
	CurrentTime          string  `json:"CurrentTime"`
	NextBidTime          string  `json:"NextBidTime"`
	P_available          float64 `json:"P_available"`
	Price_bid            float64 `json:"Price_bid"`
	Price_cleared        float64 `json:"Price_cleared"`
	P_cap                float64 `json:"P_cap"`
	Regulation_signal    float64 `json:"Regulation_signal"`
	P_reg                float64 `json:"P_reg"`
	P_res                float64 `json:"P_res"`
	P_mis_1              float64 `json:"P_mis_1"`
	P_mis_2              float64 `json:"P_mis_2"`
	Uncertainty          float64 `json:"Uncertainty"`
	Price_penalty        float64 `json:"Price_penalty"`
	Revenue_capacity     float64 `json:"Revenue_capacity"`
	Revenue_mileage      float64 `json:"Revenue_mileage"`
	Penalty              float64 `json:"Penalty"`
}
