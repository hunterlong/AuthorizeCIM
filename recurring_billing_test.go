package AuthorizeCIM

import (
	"testing"
)

var newSubscriptionId string

func TestCreateSubscription(t *testing.T) {
	subscription := Subscription{
		Name:   "New Subscription",
		Amount: "9.00",
		//TrialAmount: "0.00",
		PaymentSchedule: &PaymentSchedule{
			StartDate:        CurrentDate(),
			TotalOccurrences: "9999",
			//TrialOccurrences: "0",
			Interval: IntervalMonthly(),
		},
		Payment: &Payment{
			CreditCard: CreditCard{
				CardNumber:     "4007000000027",
				ExpirationDate: "10/23",
			},
		},
		BillTo: &BillTo{
			FirstName: "Hunter",
			LastName:  "Long",
		},
		Customer: &Customer{
			ID:    "273287",
			Email: "info@newemailuser.com",
		},
	}

	response := subscription.Charge()

	if response.Approved() {
		newSubscriptionId = response.SubscriptionID
		t.Log("New Subscription: ", response.SubscriptionID)
	} else {
		t.Log(response.ErrorMessage(), "\n")
	}

}

func TestGetSubscription(t *testing.T) {

	sub := SetSubscription{
		Id: newSubscriptionId,
	}

	subscriptionInfo := sub.Info()

	t.Log("Subscription Name: #", subscriptionInfo.Subscription.Name, "\n")
	t.Log("Subscription Status: ", subscriptionInfo.Subscription.Status, "\n")

}

func TestGetSubscriptionStatus(t *testing.T) {

	sub := SetSubscription{
		Id: newSubscriptionId,
	}

	subscriptionInfo := sub.Status()

	t.Log("Subscription ID has status: ", subscriptionInfo.Status)

}

func TestUpdateSubscription(t *testing.T) {

	subscription := Subscription{
		Payment: &Payment{
			CreditCard: CreditCard{
				CardNumber:     "5424000000000015",
				ExpirationDate: "06/25",
			},
		},
		SubscriptionId: newSubscriptionId,
	}

	response := subscription.Update()

	if response.Approved() {
		t.Log("Updated Subscription \n")
	} else {
		t.Log(response.ErrorMessage(), "\n")
		t.Fail()
	}

}

func TestCancelSubscription(t *testing.T) {

	sub := SetSubscription{
		Id: newSubscriptionId,
	}

	subscriptionInfo := sub.Cancel()

	if subscriptionInfo.Approved() {
		t.Log("Subscription ID has been canceled: ", sub.Id, "\n")
		t.Log(subscriptionInfo.ErrorMessage(), "\n")
	} else {
		t.Log(subscriptionInfo.ErrorMessage())
		t.Fail()
	}

}

func TestGetInactiveSubscriptionList(t *testing.T) {

	subscriptionList := SubscriptionList("subscriptionInactive")
	count := subscriptionList.Count()
	t.Log("Amount of Inactive Subscriptions: ", count)

	if count == 0 {
		t.Fail()
	}

}

func TestGetActiveSubscriptionList(t *testing.T) {

	subscriptionList := SubscriptionList("subscriptionActive")
	count := subscriptionList.Count()
	t.Log("Amount of Active Subscriptions: ", count)

	if count == 0 {
		t.Fail()
	}

}

func TestGetExpiringSubscriptionList(t *testing.T) {

	subscriptionList := SubscriptionList("subscriptionExpiringThisMonth")

	t.Log("Amount of Subscriptions Expiring This Month: ", subscriptionList.Count())

}

func TestGetCardExpiringSubscriptionList(t *testing.T) {

	subscriptionList := SubscriptionList("cardExpiringThisMonth")

	t.Log("Amount of Cards Expiring This Month: ", subscriptionList.Count())

}

func TestDeleteCustomerShippingProfile(t *testing.T) {
	customer := Customer{
		ID:         newCustomerProfileId,
		ShippingID: newCustomerShippingId,
	}

	response := customer.DeleteShippingProfile()

	if response.Approved() {
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

	if response.Approved() {
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

	if response.Approved() {
		t.Log("Customer was Deleted")
	} else {
		t.Log(response.ErrorMessage())
		t.Fail()
	}

}
