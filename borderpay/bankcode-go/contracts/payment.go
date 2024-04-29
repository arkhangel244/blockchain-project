package banking

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Status string

const (
	Open        Status = "open"
	Rejected    Status = "rejected"
	Transferred Status = "closed"
)

type Currency string

const (
	Dollar Currency = "dollar"
	Rupees Currency = "rupees"
	Euro   Currency = "euro"
	Yen    Currency = "yen"
)

// PaymentContract defines the smart contract
type PaymentContract struct {
	contractapi.Contract
}

// PaymentRequest defines the structure for a payment request
type PaymentRequest struct {
	FromAccount string   `json:"fromAccount"`
	ToAccount   string   `json:"toAccount"`
	Amount      float64  `json:"amount"`
	CreatedBy   string   `json:"createdby"`
	Status      Status   `json:"Status"`
	Currency    Currency `json:"Currency"`
}

// PaymentResponse defines the structure for a payment response
type PaymentResponse struct {
	TransactionID string `json:"transactionID"`
	FromAccount   string `json:"fromAccount"`
	ToAccount     string `json:"toAccount"`
	Amount        int    `json:"amount"`
	Message       string `json:"message"`
}

func (s *PaymentContract) CreateInvoice(ctx contractapi.TransactionContextInterface, invoiceID, fromAccount, toAccount string, amount float64) error {

	// get ID of submitting client
	clientID, err := s.GetSubmittingClientIdentity(ctx)
	if err != nil {
		return fmt.Errorf("failed to get client identity %v", err)
	}

	// get org of submitting client
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("failed to get client identity %v", err)
	}

	invoice := PaymentRequest{
		CreatedBy:   clientID,
		FromAccount: fromAccount,
		ToAccount:   toAccount,
		Amount:      amount,
	}

	invoiceJSON, err := json.Marshal(invoice)
	if err != nil {
		return err
	}

	// put auction into state
	err = ctx.GetStub().PutState(invoiceID, invoiceJSON)
	if err != nil {
		return fmt.Errorf("failed to put auction in public data: %v", err)
	}

	// set the seller of the auction as an endorser
	err = setAssetStateBasedEndorsement(ctx, invoiceID, clientOrgID)
	if err != nil {

		return fmt.Errorf("failed setting state based endorsement for new organization: %v", err)
	}

	return nil
}
func (s *PaymentContract) MakePayment(ctx contractapi.TransactionContextInterface, invoiceID string) error {

	// get ID of submitting client
	clientID, err := s.GetSubmittingClientIdentity(ctx)
	if err != nil {
		return fmt.Errorf("failed to get client identity %v", err)
	}

	// get org of submitting client
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("failed to get client identity %v", err)
	}
	invoiceJson, err := s.QueryInvoice(ctx, invoiceID)
	if err != nil {
		return fmt.Errorf("failed to get client identity %v", err)
	}

	// Unmarshal response JSON
	var paymentResponse PaymentResponse
	if err := json.Unmarshal(body, &paymentResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payment response: %v", err)
	}

	return &paymentResponse, nil

	invoiceJSON, err := json.Marshal(invoiceJson)
	if err != nil {
		return err
	}

	// put auction into state
	err = ctx.GetStub().PutState(invoiceID, invoiceJSON)
	if err != nil {
		return fmt.Errorf("failed to put auction in public data: %v", err)
	}

	// set the seller of the auction as an endorser
	err = setAssetStateBasedEndorsement(ctx, invoiceID, clientOrgID)
	if err != nil {

		return fmt.Errorf("failed setting state based endorsement for new organization: %v", err)
	}

	return nil
}
