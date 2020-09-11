package AuthorizeCIM

import (
	"encoding/json"
	"time"
)

func UnsettledBatchList() (*TransactionsList, error) {
	response, err := SendGetUnsettled()
	return response, err
}

func (input TransactionsList) List() ([]BatchTransaction, error) {
	response, err := SendGetUnsettled()
	return response.Transactions, err
}

func updateHeldTransaction() {

}

func (input TransactionsList) Count() int {
	return input.TotalNumInResultSet
}

type UnsettledTransactionsRequest struct {
	GetUnsettledTransactionListRequest GetUnsettledTransactionListRequest `json:"getUnsettledTransactionListRequest"`
}

type GetUnsettledTransactionListRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication,omitempty"`
	Status                 string                 `json:"status,omitempty"`
}

type TransactionsList struct {
	Transactions        []BatchTransaction `json:"transactions"`
	TotalNumInResultSet int                `json:"totalNumInResultSet"`
	MessagesResponse
}

type BatchTransaction struct {
	TransID           string    `json:"transId"`
	SubmitTimeUTC     time.Time `json:"submitTimeUTC"`
	SubmitTimeLocal   string    `json:"submitTimeLocal"`
	TransactionStatus string    `json:"transactionStatus"`
	InvoiceNumber     string    `json:"invoiceNumber"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	AccountType       string    `json:"accountType"`
	AccountNumber     string    `json:"accountNumber"`
	SettleAmount      float64   `json:"settleAmount"`
	MarketType        string    `json:"marketType"`
	Product           string    `json:"product"`
	FraudInformation  struct {
		FraudFilterList []string `json:"fraudFilterList"`
		FraudAction     string   `json:"fraudAction"`
	} `json:"fraudInformation"`
}

type UpdateHeldTransactionRequest struct {
	UpdateHeldTransaction UpdateHeldTransaction `json:"updateHeldTransactionRequest"`
}

type UpdateHeldTransaction struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	RefID                  string                 `json:"refId"`
	HeldTransactionRequest HeldTransactionRequest `json:"heldTransactionRequest"`
}

type HeldTransactionRequest struct {
	Action     string `json:"action"`
	RefTransID string `json:"refTransId"`
}

func SendTransactionUpdate(tranx PreviousTransaction, method string) (*TransactionResponse, error) {
	action := UpdateHeldTransactionRequest{
		UpdateHeldTransaction: UpdateHeldTransaction{
			MerchantAuthentication: GetAuthentication(),
			RefID: tranx.RefID,
			HeldTransactionRequest: HeldTransactionRequest{
				Action:     method,
				RefTransID: tranx.RefID,
			},
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		return nil, err
	}
	response, err := SendRequest(jsoned)
	var dat TransactionResponse
	err = json.Unmarshal(response, &dat)
	if err != nil {
		return nil, err
	}
	return &dat, err
}

func (t PreviousTransaction) Approve() (*TransactionResponse, error) {
	response, err := SendTransactionUpdate(t, "approve")
	return response, err
}

func (t PreviousTransaction) Decline() (*TransactionResponse, error) {
	response, err := SendTransactionUpdate(t, "decline")
	return response, err
}

func SendGetUnsettled() (*TransactionsList, error) {
	action := UnsettledTransactionsRequest{
		GetUnsettledTransactionListRequest: GetUnsettledTransactionListRequest{
			MerchantAuthentication: GetAuthentication(),
			Status:                 "pendingApproval",
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		return nil, err
	}
	response, err := SendRequest(jsoned)
	var dat TransactionsList
	err = json.Unmarshal(response, &dat)
	if err != nil {
		return nil, err
	}
	return &dat, err
}
