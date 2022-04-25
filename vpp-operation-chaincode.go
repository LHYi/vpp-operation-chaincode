package main

import (
	"encoding/json"
	"fmt"

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

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assets := []Asset{
		{ID: "DER-1", Regulation_direction: "Up", CurrentTime: "04-22-01", NextBidTime: "04-22-03",
			P_available: 0, Price_bid: 0, Price_cleared: 0, P_cap: 0, Regulation_signal: 0, P_reg: 0,
			P_res: 0, P_mis_1: 0, P_mis_2: 0, Uncertainty: 0, Price_penalty: 0, Revenue_capacity: 0,
			Revenue_mileage: 0, Penalty: 0},
		{ID: "DER-1", Regulation_direction: "Down", CurrentTime: "04-22-01", NextBidTime: "04-22-03",
			P_available: 0, Price_bid: 0, Price_cleared: 0, P_cap: 0, Regulation_signal: 0, P_reg: 0,
			P_res: 0, P_mis_1: 0, P_mis_2: 0, Uncertainty: 0, Price_penalty: 0, Revenue_capacity: 0,
			Revenue_mileage: 0, Penalty: 0},
		{ID: "DER-2", Regulation_direction: "Up", CurrentTime: "04-22-01", NextBidTime: "04-22-03",
			P_available: 0, Price_bid: 0, Price_cleared: 0, P_cap: 0, Regulation_signal: 0, P_reg: 0,
			P_res: 0, P_mis_1: 0, P_mis_2: 0, Uncertainty: 0, Price_penalty: 0, Revenue_capacity: 0,
			Revenue_mileage: 0, Penalty: 0},
		{ID: "DER-2", Regulation_direction: "Down", CurrentTime: "04-22-01", NextBidTime: "04-22-03",
			P_available: 0, Price_bid: 0, Price_cleared: 0, P_cap: 0, Regulation_signal: 0, P_reg: 0,
			P_res: 0, P_mis_1: 0, P_mis_2: 0, Uncertainty: 0, Price_penalty: 0, Revenue_capacity: 0,
			Revenue_mileage: 0, Penalty: 0},
		{ID: "DER-3", Regulation_direction: "Up", CurrentTime: "04-22-01", NextBidTime: "04-22-03",
			P_available: 0, Price_bid: 0, Price_cleared: 0, P_cap: 0, Regulation_signal: 0, P_reg: 0,
			P_res: 0, P_mis_1: 0, P_mis_2: 0, Uncertainty: 0, Price_penalty: 0, Revenue_capacity: 0,
			Revenue_mileage: 0, Penalty: 0},
		{ID: "DER-3", Regulation_direction: "Down", CurrentTime: "04-22-01", NextBidTime: "04-22-03",
			P_available: 0, Price_bid: 0, Price_cleared: 0, P_cap: 0, Regulation_signal: 0, P_reg: 0,
			P_res: 0, P_mis_1: 0, P_mis_2: 0, Uncertainty: 0, Price_penalty: 0, Revenue_capacity: 0,
			Revenue_mileage: 0, Penalty: 0},
		{ID: "DER-4", Regulation_direction: "Up", CurrentTime: "04-22-01", NextBidTime: "04-22-03",
			P_available: 0, Price_bid: 0, Price_cleared: 0, P_cap: 0, Regulation_signal: 0, P_reg: 0,
			P_res: 0, P_mis_1: 0, P_mis_2: 0, Uncertainty: 0, Price_penalty: 0, Revenue_capacity: 0,
			Revenue_mileage: 0, Penalty: 0},
		{ID: "DER-4", Regulation_direction: "Down", CurrentTime: "04-22-01", NextBidTime: "04-22-03",
			P_available: 0, Price_bid: 0, Price_cleared: 0, P_cap: 0, Regulation_signal: 0, P_reg: 0,
			P_res: 0, P_mis_1: 0, P_mis_2: 0, Uncertainty: 0, Price_penalty: 0, Revenue_capacity: 0,
			Revenue_mileage: 0, Penalty: 0},
		{ID: "DER-5", Regulation_direction: "Up", CurrentTime: "04-22-01", NextBidTime: "04-22-03",
			P_available: 0, Price_bid: 0, Price_cleared: 0, P_cap: 0, Regulation_signal: 0, P_reg: 0,
			P_res: 0, P_mis_1: 0, P_mis_2: 0, Uncertainty: 0, Price_penalty: 0, Revenue_capacity: 0,
			Revenue_mileage: 0, Penalty: 0},
		{ID: "DER-5", Regulation_direction: "Down", CurrentTime: "04-22-01", NextBidTime: "04-22-03",
			P_available: 0, Price_bid: 0, Price_cleared: 0, P_cap: 0, Regulation_signal: 0, P_reg: 0,
			P_res: 0, P_mis_1: 0, P_mis_2: 0, Uncertainty: 0, Price_penalty: 0, Revenue_capacity: 0,
			Revenue_mileage: 0, Penalty: 0},
	}

	for _, asset := range assets {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		eventFrom, err0 := ctx.GetClientIdentity().GetMSPID()
		if err0 != nil {
			return fmt.Errorf("Failed to get client identity")
		}

		payload := "Ledger Initialized by " + eventFrom
		payloadAsBytes := []byte(payload)

		eventErr := ctx.GetStub().SetEvent(eventFrom, payloadAsBytes)
		if eventErr != nil {
			return fmt.Errorf("%v", eventErr)
		}

		err = ctx.GetStub().PutState(asset.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}
