package AuthorizeCIM

import (
	"testing"
)

var newSubscriptionId string

func TestCreateSubscription(t *testing.T) {
	subscription := Subscription{
		Name:        "New Subscription",
		Amount:      "9.00",
		TrialAmount: "0.00",
		PaymentSchedule: &PaymentSchedule{
			StartDate:        CurrentTime(),
			TotalOccurrences: "9999",
			TrialOccurrences: "0",
			Interval: Interval{
				Length: "1",
				Unit:   "months",
			},
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

	t.Log(subscriptionInfo)

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
	}

}

func TestCancelSubscription(t *testing.T) {

	sub := SetSubscription{
		Id: newSubscriptionId,
	}

	subscriptionInfo := sub.Cancel()

	t.Log("Subscription ID has been canceled: ", sub.Id, "\n")
	t.Log(subscriptionInfo.ErrorMessage(), "\n")

}

func TestGetInactiveSubscriptionList(t *testing.T) {

	subscriptionList := SubscriptionList("subscriptionInactive")

	t.Log("Amount of Inactive Subscriptions: ", subscriptionList.Count())

}

func TestGetActiveSubscriptionList(t *testing.T) {

	subscriptionList := SubscriptionList("subscriptionActive")

	t.Log("Amount of Active Subscriptions: ", subscriptionList.Count())

}

func TestGetExpiringSubscriptionList(t *testing.T) {

	subscriptionList := SubscriptionList("subscriptionExpiringThisMonth")

	t.Log("Amount of Subscriptions Expiring This Month: ", subscriptionList.Count())

}

func TestGetCardExpiringSubscriptionList(t *testing.T) {

	subscriptionList := SubscriptionList("cardExpiringThisMonth")

	t.Log("Amount of Cards Expiring This Month: ", subscriptionList.Count())

}
