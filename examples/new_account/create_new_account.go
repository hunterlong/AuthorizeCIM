package main

import (
	"github.com/hunterlong/authorizecim"
	"fmt"
	"os"
	"time"
	"math/rand"
)

func main() {

	apiName := os.Getenv("apiName")
	apiKey := os.Getenv("apiKey")
	AuthorizeCIM.SetAPIInfo(apiName,apiKey,"test")

	connected := AuthorizeCIM.TestConnection()

	// Create random email address so it won't make duplicate records
	newUserEmail := RandomString(7)+"@domain.com"

	if connected {
		fmt.Println("Successful Authorize.net Connection")
	} else {
		fmt.Println("There was an issue connecting to Authorize.net")
	}

	customer_info := AuthorizeCIM.AuthUser{
		"74",
		newUserEmail,
		"Test Account",
	}
	new_customer, success := AuthorizeCIM.CreateCustomerProfile(customer_info)

	if success {
		fmt.Println("New Customer Profile ID: ",new_customer)
	} else {
		fmt.Println("There was an issue creating the Customer Profile")
	}

}


// NOT NEEDED - ONLY FOR CREATING A RANDOM EMAIL ADDRESS
func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
