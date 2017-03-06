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
		PaymentSchedule: PaymentSchedule{
			StartDate:        CurrentTime(),
			TotalOccurrences: "9999",
			TrialOccurrences: "0",
			Interval: Interval{
				Length: "1",
				Unit:   "months",
			},
		},
		Payment: Payment{
			CreditCard: CreditCard{
				CardNumber:     "4007000000027",
				ExpirationDate: "10/23",
			},
		},
		BillTo: BillTo{
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

func TestCancelSubscription(t *testing.T) {

	sub := SetSubscription{
		Id: newSubscriptionId,
	}

	subscriptionInfo := sub.Cancel()

	t.Log("Subscription ID has been canceled: ", subscriptionInfo.Messages.Message[0].Text)

}

func TestGetSubscriptionList(t *testing.T) {

	subscriptionList := SubscriptionList("subscriptionInactive")

	t.Log("Amount of Subscriptions: ", subscriptionList.Count())

}
