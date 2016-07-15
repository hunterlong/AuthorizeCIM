package main

import (
	"github.com/hunterlong/authorizecim"
	"fmt"
	"time"
)

func main() {

	apiName := os.Getenv("apiName")
	apiKey := os.Getenv("apiKey")
	AuthorizeCIM.SetAPIInfo(apiName,apiKey,"test")

	connected := AuthorizeCIM.TestConnection()

	if connected {
		fmt.Println("Successful Authorize.net Connection")
	} else {
		fmt.Println("There was an issue connecting to Authorize.net")
	}

	customer_info := AuthorizeCIM.AuthUser{
		"399",
		"ncxokpllai@domain.com",
		"Test Account",
	}

	new_customer, success := AuthorizeCIM.CreateCustomerProfile(customer_info)

	if success {
		fmt.Println("New Customer Profile ID: ",new_customer)
	} else {
		fmt.Println("There was an issue creating the Customer Profile")
	}


	address := AuthorizeCIM.Address{
		FirstName: "Test",
		LastName: "User",
		Address: "1234 Road St",
		City: "City Name",
		State:" California",
		Zip: "93063",
		Country: "USA",
		PhoneNumber: "5555555555",
	}
	credit_card := AuthorizeCIM.CreditCard{
		CardNumber: "4111111111111111",
		ExpirationDate: "2020-12",
	}
	profile_id := new_customer

	newPaymentID, success := AuthorizeCIM.CreateCustomerBillingProfile(profile_id, credit_card, address)

	if success {
		fmt.Println("New Credit Card was added, Billing ID: ",newPaymentID)
	} else {
		fmt.Println("There was an issue inserting a credit card into the user account")
	}



	fmt.Println("Waiting for 10 seconds to allow Authorize.net to keep up")
	time.Sleep(10000 * time.Millisecond)



	item := AuthorizeCIM.LineItem{
		ItemID: "S0897",
		Name: "New Product",
		Description: "brand new",
		Quantity: "1",
		UnitPrice: "14.43",
	}
	amount := "14.43"


	payment_id := newPaymentID

	response, approved, success := AuthorizeCIM.CreateTransaction(profile_id, payment_id, item, amount)
	// outputs transaction response, approved status (true/false), and success status (true/false)

	var tranxID string
	if success {
		tranxID = response["transId"].(string)
		if approved {
			fmt.Println("Transaction was approved! "+tranxID+"\n")
		} else {
			fmt.Println("Transaction was denied! "+tranxID+"\n")
		}
	} else {
		fmt.Println("Transaction has failed! \n")
		fmt.Println(response)
	}


}