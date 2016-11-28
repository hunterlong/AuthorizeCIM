package main

import (
	"github.com/hunterlong/authorizecim"
	"fmt"
	"time"
	"os"
	"math/rand"
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

	credit_card := AuthorizeCIM.CreditCardCVV{
		CardNumber: "4111111111111111",
		ExpirationDate: "2020-12",
		CardCode: "123",
	}
	amount := "14.43"

	response, approved, success := AuthorizeCIM.AuthorizeCard(credit_card, amount)
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

