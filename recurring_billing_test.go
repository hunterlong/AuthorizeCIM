package AuthorizeCIM

import (
	"testing"
)

func TestCreateSubscription(t *testing.T) {
	subscription := Subscription{
		Name:"New Subscription",
		Amount: "8.00",
		TrialAmount: "0.00",
		PaymentSchedule: PaymentSchedule{
			StartDate: CurrentTime(),
			TotalOccurrences: "9999",
			TrialOccurrences: "0",
			Interval: Interval{
				Length: "1",
				Unit: "months",
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
			LastName: "Long",
		},
	}

	response := subscription.Charge()

	if response.Approved() {
		t.Log("New Subscription: ",response.SubscriptionID)
	} else {
		t.Log(response.ErrorMessage(), "\n")
	}

}

func TestGetSubscription(t *testing.T) {

	sub := SetSubscription{
		Id: "09090990",
	}

	subscriptionInfo := sub.Info()

	t.Log(subscriptionInfo)

}
