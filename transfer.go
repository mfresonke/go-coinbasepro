package coinbasepro

import (
	"fmt"
)

type Transfer struct {
	Type              string `json:"type"`
	Amount            string `json:"amount"`
	CoinbaseAccountID string `json:"coinbase_account_id,string"`
}

func (c *Client) CreateTransfer(newTransfer *Transfer) (Transfer, error) {
	var savedTransfer Transfer

	url := fmt.Sprintf("/transfers")
	_, err := c.Request("POST", url, newTransfer, &savedTransfer)
	return savedTransfer, err
}

type PaymentMethod struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Currency string `json:"currency"`
}

func (c *Client) PaymentMethods() ([]PaymentMethod, error) {
	var paymentMethods []PaymentMethod

	url := fmt.Sprintf("/payment-methods")
	_, err := c.Request("GET", url, nil, &paymentMethods)
	return paymentMethods, err
}

type PaymentMethodDeposit struct {
	Amount          string `json:"amount"`
	Currency        string `json:"currency"`
	PaymentMethodID string `json:"payment_method_id"`
	// Response  Fields
	ID string `json:"id,omitempty"`
}

func (c *Client) PaymentMethodDeposit(deposit *PaymentMethodDeposit) (PaymentMethodDeposit, error) {
	var savedDeposit PaymentMethodDeposit

	url := fmt.Sprintf("/deposits/payment-method")
	_, err := c.Request("POST", url, deposit, &savedDeposit)
	return savedDeposit, err
}
