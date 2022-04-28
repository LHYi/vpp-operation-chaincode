package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
type Asset struct {
	ID                string  `json:"ID"`
	CurrentTime       string  `json:"CurrentTime"`
	NextBidTime       string  `json:"NextBidTime"`
	P_available       float64 `json:"P_available"`
	Price_bid         float64 `json:"Price_bid"`
	Price_cleared     float64 `json:"Price_cleared"`
	P_cap             float64 `json:"P_cap"`
	Regulation_signal float64 `json:"Regulation_signal"`
	P_reg             float64 `json:"P_reg"`
	P_res             float64 `json:"P_res"`
	P_mis_1           float64 `json:"P_mis_1"`
	P_mis_2           float64 `json:"P_mis_2"`
	Uncertainty       float64 `json:"Uncertainty"`
	Price_penalty     float64 `json:"Price_penalty"`
	Revenue_capacity  float64 `json:"Revenue_capacity"`
	Revenue_mileage   float64 `json:"Revenue_mileage"`
	Penalty           float64 `json:"Penalty"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assets := []Asset{
		{ID: "DER-1-Up", CurrentTime: "04-22-01", NextBidTime: "04-22-03",
			P_available: 0, Price_bid: 0, Price_cleared: 0, P_cap: 0, Regulation_signal: 0, P_reg: 0,
			P_res: 0, P_mis_1: 0, P_mis_2: 0, Uncertainty: 0, Price_penalty: 0, Revenue_capacity: 0,
			Revenue_mileage: 0, Penalty: 0},
		{ID: "DER-1-Down", CurrentTime: "04-22-01", NextBidTime: "04-22-03",
			P_available: 0, Price_bid: 0, Price_cleared: 0, P_cap: 0, Regulation_signal: 0, P_reg: 0,
			P_res: 0, P_mis_1: 0, P_mis_2: 0, Uncertainty: 0, Price_penalty: 0, Revenue_capacity: 0,
			Revenue_mileage: 0, Penalty: 0},
		{ID: "DER-2-Up", CurrentTime: "04-22-01", NextBidTime: "04-22-03",
			P_available: 0, Price_bid: 0, Price_cleared: 0, P_cap: 0, Regulation_signal: 0, P_reg: 0,
			P_res: 0, P_mis_1: 0, P_mis_2: 0, Uncertainty: 0, Price_penalty: 0, Revenue_capacity: 0,
			Revenue_mileage: 0, Penalty: 0},
		{ID: "DER-2-Down", CurrentTime: "04-22-01", NextBidTime: "04-22-03",
			P_available: 0, Price_bid: 0, Price_cleared: 0, P_cap: 0, Regulation_signal: 0, P_reg: 0,
			P_res: 0, P_mis_1: 0, P_mis_2: 0, Uncertainty: 0, Price_penalty: 0, Revenue_capacity: 0,
			Revenue_mileage: 0, Penalty: 0},
		{ID: "DER-3-Up", CurrentTime: "04-22-01", NextBidTime: "04-22-03",
			P_available: 0, Price_bid: 0, Price_cleared: 0, P_cap: 0, Regulation_signal: 0, P_reg: 0,
			P_res: 0, P_mis_1: 0, P_mis_2: 0, Uncertainty: 0, Price_penalty: 0, Revenue_capacity: 0,
			Revenue_mileage: 0, Penalty: 0},
		{ID: "DER-3-Down", CurrentTime: "04-22-01", NextBidTime: "04-22-03",
			P_available: 0, Price_bid: 0, Price_cleared: 0, P_cap: 0, Regulation_signal: 0, P_reg: 0,
			P_res: 0, P_mis_1: 0, P_mis_2: 0, Uncertainty: 0, Price_penalty: 0, Revenue_capacity: 0,
			Revenue_mileage: 0, Penalty: 0},
		{ID: "DER-4-Up", CurrentTime: "04-22-01", NextBidTime: "04-22-03",
			P_available: 0, Price_bid: 0, Price_cleared: 0, P_cap: 0, Regulation_signal: 0, P_reg: 0,
			P_res: 0, P_mis_1: 0, P_mis_2: 0, Uncertainty: 0, Price_penalty: 0, Revenue_capacity: 0,
			Revenue_mileage: 0, Penalty: 0},
		{ID: "DER-4-Down", CurrentTime: "04-22-01", NextBidTime: "04-22-03",
			P_available: 0, Price_bid: 0, Price_cleared: 0, P_cap: 0, Regulation_signal: 0, P_reg: 0,
			P_res: 0, P_mis_1: 0, P_mis_2: 0, Uncertainty: 0, Price_penalty: 0, Revenue_capacity: 0,
			Revenue_mileage: 0, Penalty: 0},
		{ID: "DER-5-Up", CurrentTime: "04-22-01", NextBidTime: "04-22-03",
			P_available: 0, Price_bid: 0, Price_cleared: 0, P_cap: 0, Regulation_signal: 0, P_reg: 0,
			P_res: 0, P_mis_1: 0, P_mis_2: 0, Uncertainty: 0, Price_penalty: 0, Revenue_capacity: 0,
			Revenue_mileage: 0, Penalty: 0},
		{ID: "DER-5-Down", CurrentTime: "04-22-01", NextBidTime: "04-22-03",
			P_available: 0, Price_bid: 0, Price_cleared: 0, P_cap: 0, Regulation_signal: 0, P_reg: 0,
			P_res: 0, P_mis_1: 0, P_mis_2: 0, Uncertainty: 0, Price_penalty: 0, Revenue_capacity: 0,
			Revenue_mileage: 0, Penalty: 0},
	}

	for _, asset := range assets {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// AddDER adds a new DER to the ledger
func (s *SmartContract) AddDER(ctx contractapi.TransactionContextInterface, id string, CurrentTime string, NextBidTime string) error {
	exists, err := s.DERExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the DER %s already exists", id)
	}

	asset := Asset{
		ID: id, CurrentTime: CurrentTime, NextBidTime: NextBidTime,
		P_available: 0, Price_bid: 0, Price_cleared: 0, P_cap: 0, Regulation_signal: 0, P_reg: 0,
		P_res: 0, P_mis_1: 0, P_mis_2: 0, Uncertainty: 0, Price_penalty: 0, Revenue_capacity: 0,
		Revenue_mileage: 0, Penalty: 0,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// ReadDER returns the asset stored in the world state with given id.
func (s *SmartContract) ReadDER(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the DER %s does not exist", id)
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

// GetAllDERs returns all DERs found in world state
func (s *SmartContract) GetAllDERs(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Asset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

// UpdateTime updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdateTime(ctx contractapi.TransactionContextInterface, currentTime string, nextBidTime string) error {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return err
	}
	defer resultsIterator.Close()

	var id string
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return err
		}

		var asset Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return err
		}
		id = asset.ID
		asset.CurrentTime = currentTime
		asset.NextBidTime = nextBidTime
		assetJSON, err1 := json.Marshal(asset)
		if err1 != nil {
			return err1
		}
		err = ctx.GetStub().PutState(id, assetJSON)
		if err != nil {
			return err
		}
	}

	return nil
}

// Update market clearing result to all teh DERs
func (s *SmartContract) UpdateMarketResult(ctx contractapi.TransactionContextInterface, id string, price_cleared float64, P_cap float64) error {
	asset, err := s.ReadDER(ctx, id)
	if err != nil {
		return err
	}
	asset.Price_cleared = price_cleared
	asset.P_cap = P_cap

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// Set regulating capacity for each asset
func (s *SmartContract) UpdateReg(ctx contractapi.TransactionContextInterface, id string, regulationSignal float64, P_reg float64) error {
	asset, err := s.ReadDER(ctx, id)
	if err != nil {
		return err
	}
	asset.Regulation_signal = regulationSignal
	asset.P_reg = P_reg
	asset.P_mis_1 = math.Max(asset.P_cap*regulationSignal-asset.P_reg, 0)

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// Update response of the DERs
func (s *SmartContract) UpdateRes(ctx contractapi.TransactionContextInterface, id string, P_res float64) error {
	asset, err := s.ReadDER(ctx, id)
	if err != nil {
		return err
	}
	asset.P_res = P_res
	asset.P_mis_2 = math.Abs(P_res - asset.P_reg)
	asset.Uncertainty = asset.Uncertainty + 0.5*asset.P_mis_1 + 0.5*asset.P_mis_2

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// Update bidding information
func (s *SmartContract) UpdateBid(ctx contractapi.TransactionContextInterface, id string)

// DERExists returns true when DER with given ID exists in world state
func (s *SmartContract) DERExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

func main() {
	VPPChaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("Error creating VPPChaincode: %v", err)
	}

	if err := VPPChaincode.Start(); err != nil {
		log.Panicf("Error starting VPPChaincode: %v", err)
	}
}
