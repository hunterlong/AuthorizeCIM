package AuthorizeCIM

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

var newCustomerProfileId string
var newCustomerPaymentId string
var newCustomerShippingId string
var newSecondCustomerProfileId string

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestSetAPIInfo(t *testing.T) {
	apiName := os.Getenv("apiName")
	apiKey := os.Getenv("apiKey")
	//apiMode := os.Getenv("mode")
	SetAPIInfo(apiName, apiKey, "test")
	t.Log("API Info Set")
}

func TestIsConnected(t *testing.T) {
	authenticated, err := IsConnected()
	if err != nil {
		t.Fail()
	}
	if !authenticated {
		t.Fail()
	}
}

func TestCreateCustomerProfile(t *testing.T) {

	customer := Profile{
		MerchantCustomerID: RandomNumber(1000, 9999),
		Email:              "info@" + RandomString(8) + ".com",
		PaymentProfiles: &PaymentProfiles{
			CustomerType: "individual",
			Payment: Payment{
				CreditCard: CreditCard{
					CardNumber:     "4007000000027",
					ExpirationDate: "10/26",
					//CardCode: "384",
				},
			},
		},
	}

	response, err := customer.CreateProfile()
	if err != nil {
		t.Fail()
	}

	if response.Ok() {
		newCustomerProfileId = response.CustomerProfileID
		t.Log("New Customer Profile Created #", response.CustomerProfileID)
	} else {
		t.Fail()
		t.Log(response.ErrorMessage())
	}

}

func TestGetProfileIds(t *testing.T) {
	profiles, _ := GetProfileIds()

	for _, p := range profiles {
		t.Log("Profile ID #", p)
	}

	if len(profiles) == 0 {
		t.Fail()
	}

	t.Log(profiles)
}

func TestUpdateCustomerProfile(t *testing.T) {

	customer := Profile{
		MerchantCustomerID: newCustomerProfileId,
		CustomerProfileId:  newCustomerProfileId,
		Description:        "Updated Account",
		Email:              "info@updatedemail.com",
	}

	response, err := customer.UpdateProfile()
	if err != nil {
		t.Fail()
	}

	if response.Ok() {
		t.Log("Customer Profile was Updated")
	} else {
		t.Log(response.ErrorMessage())
		t.Fail()
	}

}

func TestCreateCustomerPaymentProfile(t *testing.T) {

	paymentProfile := CustomerPaymentProfile{
		CustomerProfileID: newCustomerProfileId,
		PaymentProfile: PaymentProfile{
			BillTo: &BillTo{
				FirstName:   "okokk",
				LastName:    "okok",
				Address:     "1111 white ct",
				City:        "los angeles",
				Country:     "USA",
				PhoneNumber: "8885555555",
			},
			Payment: &Payment{
				CreditCard: CreditCard{
					CardNumber:     "5424000000000015",
					ExpirationDate: "04/22",
				},
			},
			DefaultPaymentProfile: "true",
		},
	}

	response, err := paymentProfile.Add()
	if err != nil {
		t.Fail()
	}

	if response.Ok() {
		newCustomerPaymentId = response.CustomerPaymentProfileID
		t.Log("Created new Payment Profile #", response.CustomerPaymentProfileID, "for Customer ID: ", response.CustomerProfileId)
	} else {
		t.Log(response.ErrorMessage())
		t.Fail()
	}

}

func TestGetCustomerPaymentProfile(t *testing.T) {

	customer := Customer{
		ID: newCustomerProfileId,
	}

	response, err := customer.Info()
	if err != nil {
		t.Fail()
	}

	paymentProfiles := response.PaymentProfiles()

	t.Log("Customer Payment Profiles", paymentProfiles)

	if len(paymentProfiles) == 0 {
		t.Fail()
	}

}

func TestGetCustomerPaymentProfileList(t *testing.T) {

	profileIds, err := GetPaymentProfileIds("2020-03", "cardsExpiringInMonth")
	if err != nil {
		t.Fail()
	}

	t.Log(profileIds)
}

func TestValidateCustomerPaymentProfile(t *testing.T) {

	customerProfile := Customer{
		ID:        newCustomerProfileId,
		PaymentID: newCustomerPaymentId,
	}

	response, err := customerProfile.Validate()
	if err != nil {
		t.Fail()
	}

	if response.Ok() {
		t.Log("Customer Payment Profile is VALID")
	} else {
		t.Log(response.ErrorMessage())
		t.Fail()
	}

}

func TestUpdateCustomerPaymentProfile(t *testing.T) {

	customer := Profile{
		CustomerProfileId: newCustomerProfileId,
		PaymentProfileId:  newCustomerPaymentId,
		Description:       "Updated Account",
		Email:             "info@" + RandomString(8) + ".com",
		PaymentProfiles: &PaymentProfiles{
			Payment: Payment{
				CreditCard: CreditCard{
					CardNumber:     "4007000000027",
					ExpirationDate: "01/26",
				},
			},
			BillTo: &BillTo{
				FirstName:   "newname",
				LastName:    "golang",
				Address:     "2841 purple ct",
				City:        "los angeles",
				State:       "CA",
				Country:     "USA",
				PhoneNumber: "8885555555",
			},
		},
	}

	response, err := customer.UpdatePaymentProfile()
	if err != nil {
		t.Fail()
	}

	if response.Ok() {
		t.Log("Customer Payment Profile was Updated")
	} else {
		t.Log(response.ErrorMessage())
		t.Fail()
	}

}

func TestCreateCustomerShippingProfile(t *testing.T) {

	customer := Profile{
		MerchantCustomerID: "86437",
		CustomerProfileId:  newCustomerProfileId,
		Email:              "info@" + RandomString(8) + ".com",
		Shipping: &Address{
			FirstName:   "My",
			LastName:    "Name",
			Company:     "none",
			Address:     "1111 yellow ave.",
			City:        "Los Angeles",
			State:       "CA",
			Zip:         "92039",
			Country:     "USA",
			PhoneNumber: "8885555555",
		},
	}

	response, err := customer.CreateShipping()
	if err != nil {
		t.Fail()
	}

	if response.Ok() {
		newCustomerShippingId = response.CustomerAddressID
		t.Log("New Shipping Added: #", response.CustomerAddressID)
	} else {
		t.Log(response.ErrorMessage())
		t.Fail()
	}
}

func TestGetCustomerShippingProfile(t *testing.T) {

	customer := Customer{
		ID: newCustomerProfileId,
	}

	response, err := customer.Info()
	if err != nil {
		t.Fail()
	}

	shippingProfiles := response.ShippingProfiles()

	t.Log("Customer Shipping Profiles", shippingProfiles)

	if shippingProfiles[0].Zip != "92039" {
		t.Fail()
	}

}

func TestUpdateCustomerShippingProfile(t *testing.T) {

	customer := Profile{
		CustomerProfileId: newCustomerProfileId,
		CustomerAddressId: newCustomerShippingId,
		Shipping: &Address{
			FirstName:   "My",
			LastName:    "Name",
			Company:     "none",
			Address:     "1111 yellow ave.",
			City:        "Los Angeles",
			State:       "CA",
			Zip:         "92039",
			Country:     "USA",
			PhoneNumber: "8885555555",
		},
	}

	response, err := customer.UpdateShippingProfile()
	if err != nil {
		t.Fail()
	}

	if response.Ok() {
		t.Log("Shipping Address Profile was updated")
	} else {
		t.Log(response.ErrorMessage())
		t.Fail()
	}

}

func TestAcceptProfilePage(t *testing.T) {

}

func TestCreateCustomerProfileFromTransaction(t *testing.T) {

}

func TestCreateSubscriptionCustomerProfile(t *testing.T) {

	amount := RandomNumber(5, 99) + "." + RandomNumber(10, 99)

	subscription := Subscription{
		Name:   "New Customer Profile Subscription",
		Amount: amount,
		//TrialAmount: "0.00",
		PaymentSchedule: &PaymentSchedule{
			StartDate:        CurrentDate(),
			TotalOccurrences: "9999",
			//TrialOccurrences: "0",
			Interval: IntervalMonthly(),
		},
		Profile: &CustomerProfiler{
			CustomerProfileID:         newCustomerProfileId,
			CustomerPaymentProfileID:  newCustomerPaymentId,
			CustomerShippingProfileID: newCustomerShippingId,
		},
	}

	response, err := subscription.Charge()
	if err != nil {
		t.Fail()
	}

	if response.Approved() {
		newSubscriptionId = response.SubscriptionID
		t.Log("Customer #", response.CustomerProfileId(), " Created a New Subscription: ", response.SubscriptionID)
	} else {
		t.Log(response.ErrorMessage(), "\n")
		t.Fail()
	}

}

func TestGetCustomerProfile(t *testing.T) {

	customer := Customer{
		ID: newCustomerProfileId,
	}

	response, err := customer.Info()
	if err != nil {
		t.Fail()
	}

	paymentProfiles := response.PaymentProfiles()
	shippingProfiles := response.ShippingProfiles()
	subscriptions := response.Subscriptions()

	t.Log("Customer Profile", response)

	t.Log("Customer Payment Profiles", paymentProfiles)
	t.Log("Customer Shipping Profiles", shippingProfiles)
	t.Log("Customer Subscription IDs", subscriptions)

}
