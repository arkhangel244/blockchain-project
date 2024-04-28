package banking

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type BankingContract struct {
	contractapi.Contract
}

// Auction data
type Transanction struct {
	EmployeeID    string                   `json:"item"`
	Seller        string                   `json:"seller"`
	CompletedPays map[string]CompletedPays `json:"privateBids"`
	AdvancePays   map[string]AdvancePays   `json:"revealedBids"`
	Price         int                      `json:"price"`
	Status        string                   `json:"status"`
}

// FullBid is the structure of a revealed bid
type CompletedPays struct {
	Type   string `json:"objectType"`
	Price  int    `json:"price"`
	Org    string `json:"org"`
	Bidder string `json:"bidder"`
}

// BidHash is the structure of a private bid
type AdvancePays struct {
	Org  string `json:"org"`
	Hash string `json:"hash"`
}

const bidKeyType = "bid"
