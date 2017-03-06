package AuthorizeCIM

import (
	"testing"
)

var newCustomerProfileId string

func TestCreateCustomerProfile(t *testing.T) {

	customer := Profile{
		MerchantCustomerID: "86437",
		Email:              "info@emailhereooooo.com",
		PaymentProfiles: PaymentProfiles{
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
