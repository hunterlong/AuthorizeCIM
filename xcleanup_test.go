package AuthorizeCIM

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestCancelSubscription(t *testing.T) {

	sub := SetSubscription{
		Id: newSubscriptionId,
	}

	subscriptionInfo := sub.Cancel()

	if subscriptionInfo.Ok() {
		t.Log("Subscription ID has been canceled: ", sub.Id, "\n")
		t.Log(subscriptionInfo.ErrorMessage(), "\n")
	} else {
		t.Log(subscriptionInfo.ErrorMessage())
		t.Fail()
	}

}

func TestCancelSecondSubscription(t *testing.T) {

	sub := SetSubscription{
		Id: newSecondSubscriptionId,
	}

	subscriptionInfo := sub.Cancel()

	if subscriptionInfo.Ok() {
		t.Log("Second Subscription ID has been canceled: ", sub.Id, "\n")
		t.Log(subscriptionInfo.ErrorMessage(), "\n")
	} else {
		t.Log(subscriptionInfo.ErrorMessage())
		t.Fail()
	}

}

func TestDeleteCustomerShippingProfile(t *testing.T) {
	customer := Customer{
		ID:         newCustomerProfileId,
		ShippingID: newCustomerShippingId,
	}

	response := customer.DeleteShippingProfile()

	if response.Ok() {
		t.Log("Shipping Profile was Deleted")
	} else {
		t.Log(response.ErrorMessage())
		t.Fail()
	}
}

func TestDeleteCustomerPaymentProfile(t *testing.T) {
	customer := Customer{
		ID:        newCustomerProfileId,
		PaymentID: newCustomerPaymentId,
	}

	response := customer.DeletePaymentProfile()

	if response.Ok() {
		t.Log("Payment Profile was Deleted")
	} else {
		t.Log(response.ErrorMessage())
		t.Fail()
	}
}

func TestDeleteCustomerProfile(t *testing.T) {

	customer := Customer{
		ID: newCustomerProfileId,
	}

	response := customer.DeleteProfile()

	if response.Ok() {
		t.Log("Customer was Deleted")
	} else {
		t.Log(response.ErrorMessage())
		t.Fail()
	}

}

func TestDeleteSecondCustomerProfile(t *testing.T) {

	customer := Customer{
		ID: newSecondCustomerProfileId,
	}

	response := customer.DeleteProfile()

	if response.Ok() {
		t.Log("Second Customer was Deleted")
	} else {
		t.Log(response.ErrorMessage())
		t.Fail()
	}

}

func RandomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandomNumber(min, max int) string {
	num := rand.Intn(max-min) + min
	return strconv.Itoa(num)
}
