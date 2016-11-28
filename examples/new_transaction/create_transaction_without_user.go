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

	success := AuthorizeCIM.AuthorizeCard(credit_card, amount)

	if success {
		fmt.Println("Transaction was approved! \n")
	} else {
		fmt.Println("Transaction has failed! \n")
	}


}

