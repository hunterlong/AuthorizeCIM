package main

import (
	"github.com/hunterlong/authorizecim"
	"fmt"
	"os"
	"time"
	"math/rand"
	"strconv"
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

	// Create random email address so it won't make duplicate records
	newUserEmail := RandomString(7)+"@domain.com"

	customer_info := AuthorizeCIM.AuthUser{
		"70",
		newUserEmail,
		"Test Account",
	}

	customerProfileID, success := AuthorizeCIM.CreateCustomerProfile(customer_info)

	if success {
		fmt.Println("New Customer Profile ID: ",customerProfileID)
	} else {
		fmt.Println("There was an issue creating the Customer Profile")
	}

	// User's Address for billing and shipping
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

	// User's credit card
	credit_card := AuthorizeCIM.CreditCard{
		CardNumber: "4111111111111111",
		ExpirationDate: "2020-12",
	}

	newPaymentID, success := AuthorizeCIM.CreateCustomerBillingProfile(customerProfileID, credit_card, address)

	if success {
		fmt.Println("New Billing ID: ",newPaymentID)
	}

	newShippingID, success := AuthorizeCIM.CreateShippingAddress(customerProfileID, address)

	if success {
		fmt.Println("New Shipping ID: ",newShippingID)
	}

	startTime := time.Now().Format("2006-01-02")
	//startTime := "2016-06-02"
	totalRuns := "9999" //means forever
	trialRuns := "0"

	// amount here
	amount := RandomDollar(10,90)

	// Users Full Details
	userFullProfile := AuthorizeCIM.FullProfile{
		CustomerProfileID: customerProfileID,
		CustomerAddressID: newShippingID,
		CustomerPaymentProfileID: newPaymentID,
	}

	// Subscription Payment schedule
	paymentSchedule := AuthorizeCIM.PaymentSchedule{
		Interval: AuthorizeCIM.Interval{"1","months"},
		StartDate:startTime,
		TotalOccurrences:totalRuns, TrialOccurrences:trialRuns,
	}

	subscriptionInput := AuthorizeCIM.Subscription{
		"New Subscription",
		paymentSchedule,
		amount,
		"0.00",
		userFullProfile,
	}

	newSubscriptionID, success := AuthorizeCIM.CreateSubscription(subscriptionInput)

	if success {
		fmt.Println("New Subscription was created, ID: ",newSubscriptionID)
	} else {
		fmt.Println("There was an issue creating the new subscription")
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

func RandomDollar(min int, max int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	amount := float64(rand.Intn(max - min) + min)
	f := strconv.FormatFloat(amount, 'f', 2, 64)
	return f
}
