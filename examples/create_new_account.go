package main

import (
	"github.com/hunterlong/authorizecim"
	"fmt"
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
		"74",
		"email@domain.com",
		"Test Account",
	}
	new_customer, success := AuthorizeCIM.CreateCustomerProfile(customer_info)

	if success {
		fmt.Println("New Customer Profile ID: ",new_customer)
	} else {
		fmt.Println("There was an issue creating the Customer Profile")
	}

}