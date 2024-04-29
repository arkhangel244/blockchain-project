package banking

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// QueryAuction allows all members of the channel to read a public auction
func (s *PaymentContract) QueryInvoice(ctx contractapi.TransactionContextInterface, transactionID string) (*PaymentRequest, error) {

	auctionJSON, err := ctx.GetStub().GetState(transactionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get auction object %v: %v", transactionID, err)
	}
	if auctionJSON == nil {
		return nil, fmt.Errorf("auction does not exist")
	}

	var payment *PaymentRequest
	err = json.Unmarshal(auctionJSON, &payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *PaymentContract) QueryAllInvoices(ctx contractapi.TransactionContextInterface) ([]*PaymentRequest, error) {
	// Get a range iterator for all keys in the ledger
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to get state by range: %v", err)
	}
	defer resultsIterator.Close()

	var invoices []*PaymentRequest

	// Iterate over the iterator to retrieve each invoice
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to get next item from iterator: %v", err)
		}

		// Unmarshal the value (invoice) from JSON
		var payment PaymentRequest
		err = json.Unmarshal(queryResponse.Value, &payment)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal invoice JSON: %v", err)
		}

		// Append the invoice to the slice
		invoices = append(invoices, &payment)
	}

	return invoices, nil
}
