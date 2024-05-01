package borderpay

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type HiringContract struct {
	contractapi.Contract
}

type AccountContract struct {
	contractapi.Contract
}

type Status string

const (
	Open     Status = "open"
	Accepted Status = "accepted"
	Rejected Status = "rejected"
	Closed   Status = "closed"
)

type Currency string

const (
	Dollar Currency = "dollar"
	Rupees Currency = "rupees"
	Euro   Currency = "euro"
	Yen    Currency = "yen"
)

func ParseCurrency(str string) Currency {

	switch str {
	case "dollar":
		return Dollar
	case "rupees":
		return Rupees
	case "euro":
		return Euro
	case "yen":
		return Yen
	default:
		// Return a default value or handle error
		return ""
	}
}

func ParseStatus(str string) Status {

	switch str {
	case "open":
		return Open
	case "accepted":
		return Accepted
	case "rejected":
		return Rejected
	case "closed":
		return Closed
	default:
		// Return a default value or handle error
		return ""
	}
}

type Hiring struct {
	SubmittedBy   string   `json:"submittedby"`
	HiringID      string   `json:"hiringID"`
	EmployeeID    string   `json:"employeeID"`
	Company       string   `json:"Company"`
	Salary        float64  `json:"Salary"`
	AccountHiring Account  `json:"AccountHiring"`
	VariablePay   float64  `json:"VariablePay"`
	Currency      Currency `json:"Currency"`
	Status        Status   `json:"Status"`
}

type Account struct {
	Name              string   `json:"Name"`
	BankName          string   `json:"BankName"`
	PreferredCurrency Currency `json:"PreferredCurrency"`
	BankAccount       string   `json:"BankAccount"`
}

func (s *HiringContract) CreateHiring(ctx contractapi.TransactionContextInterface, hiringID string, AccountID string, salary int, variablepay int, curr string, comapny string) error {
	// get ID of submitting client
	clientID, err := s.GetSubmittingClientIdentity(ctx)
	if err != nil {
		return fmt.Errorf("failed to get client identity %v", err)
	}
	lower := strings.ToLower(string(curr))
	currency := ParseCurrency(lower)
	emptyaccount := Account{}
	hiring := Hiring{
		HiringID:      hiringID,
		SubmittedBy:   clientID,
		EmployeeID:    AccountID,
		Salary:        float64(salary),
		VariablePay:   float64(variablepay),
		Currency:      currency,
		Status:        Open,
		Company:       comapny,
		AccountHiring: emptyaccount,
	}

	hiringJSON, err := json.Marshal(hiring)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(hiringID, hiringJSON)
	if err != nil {
		return fmt.Errorf("failed to put auction in public data: %v", err)
	}
	// set the seller of the auction as an endorser
	err = setAssetStateBasedEndorsement(ctx, hiringID, AccountID)
	if err != nil {
		return fmt.Errorf("failed setting state based endorsement for new organization: %v", err)
	}
	return nil

}

func (s *HiringContract) GetAllHirings(ctx contractapi.TransactionContextInterface) ([]*Hiring, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var hirings []*Hiring
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var hiring Hiring
		err = json.Unmarshal(queryResponse.Value, &hiring)
		if err != nil {
			return nil, err
		}
		hirings = append(hirings, &hiring)
	}

	return hirings, nil
}

func (s *HiringContract) GetHiringContractsByEmployeeID(ctx contractapi.TransactionContextInterface, employeeID string) ([]byte, error) {
	allAssets, err := s.GetAllHirings(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve assets: %v", err)
	}

	// Filter assets by color
	var tempassets []*Hiring
	for _, asset := range allAssets {
		if asset.EmployeeID == employeeID {
			tempassets = append(tempassets, asset)
		}
	}

	// Marshal greenAssets to JSON
	AssetsJSON, err := json.Marshal(tempassets)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal assets to JSON: %v", err)
	}

	return AssetsJSON, nil
}

func (s *HiringContract) SubmitHiring(ctx contractapi.TransactionContextInterface, name, preferredCurrency, bankAccount, bankName, hiringID string, status string) error {

	// get the auction from public state
	hiring, err := s.QueryHiring(ctx, hiringID)
	if err != nil {
		return fmt.Errorf("failed to get auction from public state %v", err)
	}

	lower := strings.ToLower(string(preferredCurrency))
	currency := ParseCurrency(lower)
	// the auction needs to be open for users to add their bid
	statuss := strings.ToLower(status)
	Status := ParseStatus(statuss)

	if Status == Closed {
		return fmt.Errorf("cannot take action, action already taken")
	} else if Status == Accepted {
		account := Account{
			Name:              name,
			PreferredCurrency: currency,
			BankAccount:       bankAccount,
			BankName:          bankName,
		}
		hiring.Status = Closed
		hiring.AccountHiring = account
	} else {
		hiring.Status = Closed
	}

	newHiringJSON, _ := json.Marshal(hiring)

	err = ctx.GetStub().PutState(hiringID, newHiringJSON)
	if err != nil {
		return fmt.Errorf("failed to update auction: %v", err)
	}

	return nil
}

// QueryAuction allows all members of the channel to read a public auction
func (s *HiringContract) QueryHiring(ctx contractapi.TransactionContextInterface, hiringID string) (*Hiring, error) {

	auctionJSON, err := ctx.GetStub().GetState(hiringID)
	if err != nil {
		return nil, fmt.Errorf("failed to get auction object %v: %v", hiringID, err)
	}
	if auctionJSON == nil {
		return nil, fmt.Errorf("auction does not exist")
	}

	var auction *Hiring
	err = json.Unmarshal(auctionJSON, &auction)
	if err != nil {
		return nil, err
	}

	return auction, nil
}

// QueryAuction allows all members of the channel to read a public auction
func (s *HiringContract) AskForAdvance(ctx contractapi.TransactionContextInterface, hiringID string) (*Hiring, error) {

	auctionJSON, err := ctx.GetStub().GetState(hiringID)
	if err != nil {
		return nil, fmt.Errorf("failed to get auction object %v: %v", hiringID, err)
	}
	if auctionJSON == nil {
		return nil, fmt.Errorf("auction does not exist")
	}

	var auction *Hiring
	err = json.Unmarshal(auctionJSON, &auction)
	if err != nil {
		return nil, err
	}

	return auction, nil
}
