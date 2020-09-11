package main

import (
	"fmt"
	"os"

	auth "github.com/hunterlong/authorizecim"
)

var newTransactionId string

func main() {

	apiName := os.Getenv("apiName")
	apiKey := os.Getenv("apiKey")

	auth.SetAPIInfo(apiName, apiKey, "test")

	if auth.IsConnected() {
		fmt.Println("Connected to Authorize.net!")
	}

	ChargeCustomer()

	VoidTransaction()
}

func ChargeCustomer() {

	newTransaction := auth.NewTransaction{
		Amount: "13.75",
		CreditCard: auth.CreditCard{
			CardNumber:     "4012888818888",
			ExpirationDate: "08/25",
			CardCode:       "393",
		},
		BillTo: &auth.BillTo{
			FirstName:   "Timmy",
			LastName:    "Jimmy",
			Address:     "1111 green ct",
			City:        "los angeles",
			State:       "CA",
			Zip:         "43534",
			Country:     "USA",
			PhoneNumber: "8885555555",
		},
	}
	response := newTransaction.Charge()

	if response.Approved() {
		newTransactionId = response.TransactionID()
		fmt.Println("Transaction was Approved! #", response.TransactionID())
	}
}

func VoidTransaction() {

	newTransaction := auth.PreviousTransaction{
		RefID: newTransactionId,
	}
	response := newTransaction.Void()
	if response.Approved() {
		fmt.Println("Transaction was Voided!")
	}

}
