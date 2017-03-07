package AuthorizeCIM

import (
	"os"
	"testing"
)

var newCustomerProfileId string
var newCustomerPaymentId string

func TestSetAPIInfo(t *testing.T) {
	apiName := os.Getenv("apiName")
	apiKey := os.Getenv("apiKey")
	apiMode := os.Getenv("mode")
	SetAPIInfo(apiName, apiKey, apiMode)
	t.Log("API Info Set")
}

func TestCreateCustomerProfile(t *testing.T) {

	customer := Profile{
		MerchantCustomerID: "86437",
		Email:              "info@emailhereooooo.com",
		PaymentProfiles: &PaymentProfiles{
			CustomerType: "individual",
			Payment: Payment{
				CreditCard: CreditCard{
					CardNumber:     "4007000000027",
					ExpirationDate: "10/23",
				},
			},
		},
	}

	response := customer.Create()

	if response.Approved() {
		newCustomerProfileId = response.CustomerProfileID
		t.Log("New Customer Profile Created #", response.CustomerProfileID)
	} else {
		t.Log(response.ErrorMessage())
	}

}

func TestGetCustomerProfile(t *testing.T) {

	customer := Customer{
		ID: newCustomerProfileId,
	}

	response := customer.Info()

	paymentProfiles := response.PaymentProfiles()

	t.Log("Customer Profile", response)

	t.Log("Customer Payment Profiles", paymentProfiles)

}

func TestGetProfileIds(t *testing.T) {
	profiles, _ := GetProfileIds()

	for _, p := range profiles {
		t.Log("Profile ID #", p)
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

	response := customer.Update()

	if response.Approved() {
		t.Log("Customer Profile was Updated")
	} else {
		t.Log(response.ErrorMessage())
	}

}

func TestCreateCustomerPaymentProfile(t *testing.T) {

	paymentProfile := CustomerPaymentProfile{
		CustomerProfileID: newCustomerProfileId,
		PaymentProfile: PaymentProfile{
			BillTo: BillTo{
				FirstName:   "okokk",
				LastName:    "okok",
				Address:     "1111 white ct",
				City:        "los angeles",
				Country:     "USA",
				PhoneNumber: "8885555555",
			},
			Payment: Payment{
				CreditCard: CreditCard{
					CardNumber:     "5424000000000015",
					ExpirationDate: "04/22",
				},
			},
			DefaultPaymentProfile: "true",
		},
	}

	response := paymentProfile.Add()

	if response.Approved() {
		newCustomerPaymentId = response.CustomerPaymentProfileID
		t.Log("Created new Payment Profile #", response.CustomerPaymentProfileID, "for Customer ID: ", response.CustomerProfileId)
	} else {
		t.Log(response.ErrorMessage())
	}

}

func TestGetCustomerPaymentProfile(t *testing.T) {

	customer := Customer{
		ID: newCustomerProfileId,
	}

	response := customer.Info()

	paymentProfiles := response.PaymentProfiles()

	t.Log("Customer Payment Profiles", paymentProfiles)

}

func TestGetCustomerPaymentProfileList(t *testing.T) {

	profileIds := GetPaymentProfileIds("2017-03", "cardsExpiringInMonth")

	t.Log(profileIds)
}

func TestValidateCustomerPaymentProfile(t *testing.T) {

	customerProfile := Customer{
		ID:        newCustomerProfileId,
		PaymentID: newCustomerPaymentId,
	}

	response := customerProfile.Validate()

	if response.Approved() {
		t.Log("Customer Payment Profile is VALID")
	} else {
		t.Log(response.ErrorMessage())
	}

}

func TestUpdateCustomerPaymentProfile(t *testing.T) {

}

func TestDeleteCustomerPaymentProfile(t *testing.T) {

}

func TestCreateCustomerShippingProfile(t *testing.T) {

}

func TestGetCustomerShippingProfile(t *testing.T) {

}

func TestUpdateCustomerShippingProfile(t *testing.T) {

}

func TestDeleteCustomerShippingProfile(t *testing.T) {

}

func TestAcceptProfilePage(t *testing.T) {

}

func TestCreateCustomerProfileFromTransaction(t *testing.T) {

}

func TestDeleteCustomerProfile(t *testing.T) {

	customer := Customer{
		ID: "1810878365",
	}

	response := customer.Delete()

	if response.Approved() {
		t.Log("Customer was Deleted")
	} else {
		t.Log(response.ErrorMessage())
	}

}
