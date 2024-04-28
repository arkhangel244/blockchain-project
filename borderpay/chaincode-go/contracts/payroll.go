package borderpay

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type PaymentContract struct {
	contractapi.Contract
}

// Payment represents an individual payment to an employee
type Payment struct {
	EmployeeID string `json:"employeeID"`
	Month      string `json:"month"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"` // Status can be "pending", "paid"
}

// CreatePayment creates a new payment entry for an employee
func (s *PaymentContract) CreatePayment(ctx contractapi.TransactionContextInterface, employeeID string, month string, amount int) error {
	payment := Payment{
		EmployeeID: employeeID,
		Month:      month,
		Amount:     amount,
		Status:     "pending",
	}

	paymentJSON, err := json.Marshal(payment)
	if err != nil {
		return fmt.Errorf("failed to marshal payment: %v", err)
	}

	err = ctx.GetStub().PutState(employeeID+month, paymentJSON)
	if err != nil {
		return fmt.Errorf("failed to put payment into state: %v", err)
	}

	return nil
}

// GetPayment retrieves a payment entry for an employee for a specific month
func (s *PaymentContract) GetPayment(ctx contractapi.TransactionContextInterface, employeeID string, month string) (*Payment, error) {
	paymentJSON, err := ctx.GetStub().GetState(employeeID + month)
	if err != nil {
		return nil, fmt.Errorf("failed to read payment from state: %v", err)
	}
	if paymentJSON == nil {
		return nil, fmt.Errorf("payment not found for employee %s for month %s", employeeID, month)
	}

	var payment Payment
	err = json.Unmarshal(paymentJSON, &payment)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal payment: %v", err)
	}

	return &payment, nil
}

// UpdatePaymentStatus updates the status of a payment to "paid"
func (s *PaymentContract) UpdatePaymentStatus(ctx contractapi.TransactionContextInterface, employeeID string, month string) error {
	payment, err := s.GetPayment(ctx, employeeID, month)
	if err != nil {
		return err
	}

	payment.Status = "paid"

	paymentJSON, err := json.Marshal(payment)
	if err != nil {
		return fmt.Errorf("failed to marshal payment: %v", err)
	}

	err = ctx.GetStub().PutState(employeeID+month, paymentJSON)
	if err != nil {
		return fmt.Errorf("failed to update payment status in state: %v", err)
	}

	return nil
}

// GetAllPayments retrieves all payment entries for an employee
func (s *PaymentContract) GetAllPayments(ctx contractapi.TransactionContextInterface, employeeID string) ([]*Payment, error) {
	resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(employeeID, []string{})
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var payments []*Payment

	for resultsIterator.HasNext() {
		item, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var payment Payment
		err = json.Unmarshal(item.Value, &payment)
		if err != nil {
			return nil, err
		}

		payments = append(payments, &payment)
	}

	return payments, nil
}
