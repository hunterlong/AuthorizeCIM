package AuthorizeCIM

import (
	"encoding/json"
	"time"
)

func UnsettledBatchList() TransactionsList {
	response, _ := SendGetUnsettled()
	return response
}

func (input TransactionsList) List() []BatchTransaction {
	response, _ := SendGetUnsettled()
	return response.Transactions
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
	Messages            struct {
		ResultCode string `json:"resultCode"`
		Message    []struct {
			Code string `json:"code"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"messages"`
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
	SettleAmount      int       `json:"settleAmount"`
	MarketType        string    `json:"marketType"`
	Product           string    `json:"product"`
	FraudInformation  struct {
		FraudFilterList []string `json:"fraudFilterList"`
		FraudAction     string   `json:"fraudAction"`
	} `json:"fraudInformation"`
}

func SendGetUnsettled() (TransactionsList, interface{}) {
	action := UnsettledTransactionsRequest{
		GetUnsettledTransactionListRequest: GetUnsettledTransactionListRequest{
			MerchantAuthentication: GetAuthentication(),
			Status:                 "pendingApproval",
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}
	response := SendRequest(jsoned)
	var dat TransactionsList
	err = json.Unmarshal(response, &dat)
	if err != nil {
		panic(err)
	}
	return dat, err
}
