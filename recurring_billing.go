package AuthorizeCIM

import (
	"encoding/json"
	"fmt"
	"time"
)

func (response SubscriptionResponse) Approved() bool {
	if response.Messages.ResultCode == "Ok" {
		return true
	}
	return false
}

func (response SubscriptionResponse) CustomerProfileId() string {
	return response.Profile.CustomerProfileID
}

func (response SubscriptionResponse) CustomerPaymentProfileId() string {
	return response.Profile.CustomerPaymentProfileID
}

func (response SubscriptionResponse) ErrorMessage() string {
	return response.Messages.Message[0].Text
}

func (sub Subscription) Charge() SubscriptionResponse {
	response, _ := SendSubscription(sub)
	return response
}

func (response SubscriptionResponse) Info() string {
	return response.Messages.Message[0].Text
}

type CreateSubscriptionRequest struct {
	ARBCreateSubscriptionRequest ARBCreateSubscriptionRequest `json:"ARBCreateSubscriptionRequest"`
}

type ARBCreateSubscriptionRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	RefID        string `json:"refId,omitempty"`
	Subscription Subscription `json:"subscription"`
}


type GetSubscriptionRequest struct {
	ARBGetSubscriptionRequest ARBGetSubscriptionRequest `json:"ARBGetSubscriptionRequest"`
}

type ARBGetSubscriptionRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	RefID string `json:"refId"`
	SubscriptionID string `json:"subscriptionId"`
}

type SetSubscription struct {
	Id string `json:"subscriptionId"`
}

type Subscription struct {
	Name            string `json:"name"`
	PaymentSchedule PaymentSchedule `json:"paymentSchedule,omitempty"`
	Amount      string `json:"amount"`
	TrialAmount string `json:"trialAmount,omitempty"`
	Payment    Payment `json:"payment"`
	BillTo BillTo `json:"billTo"`
}

type BillTo struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type PaymentSchedule struct {
	Interval Interval `json:"interval,omitempty"`
	StartDate        string `json:"startDate,omitempty"`
	TotalOccurrences string `json:"totalOccurrences,omitempty"`
	TrialOccurrences string `json:"trialOccurrences,omitempty"`
}

type Interval struct {
	Length string `json:"length"`
	Unit   string `json:"unit"`
}

type SubscriptionResponse struct {
	SubscriptionID string `json:"subscriptionId"`
	Profile struct {
			       CustomerProfileID string `json:"customerProfileId"`
			       CustomerPaymentProfileID string `json:"customerPaymentProfileId"`
		       } `json:"profile"`
	Messages struct {
			       ResultCode string `json:"resultCode"`
			       Message []struct {
				       Code string `json:"code"`
				       Text string `json:"text"`
			       } `json:"message"`
		       } `json:"messages"`
}



func SendSubscription(sub Subscription) (SubscriptionResponse, interface{}) {
	action := CreateSubscriptionRequest{
		ARBCreateSubscriptionRequest: ARBCreateSubscriptionRequest {
			MerchantAuthentication: MerchantAuthentication{
				Name:           "8v25DGQq9kf",
				TransactionKey: "5KDX8Vz3mx334aJm",
			},
			Subscription: sub,
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}

	response := SendRequest(jsoned)
	var dat SubscriptionResponse
	fmt.Println(string(response))
	err = json.Unmarshal(response, &dat)
	if err != nil {
		panic(err)
	}
	return dat, err
}


func (sub SetSubscription) Info() (GetSubscriptionResponse) {
	action := GetSubscriptionRequest{
		ARBGetSubscriptionRequest: ARBGetSubscriptionRequest {
			MerchantAuthentication: MerchantAuthentication{
				Name:           "8v25DGQq9kf",
				TransactionKey: "5KDX8Vz3mx334aJm",
			},
			SubscriptionID: sub.Id,
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}
	response := SendRequest(jsoned)
	var dat GetSubscriptionResponse
	fmt.Println(string(response))
	err = json.Unmarshal(response, &dat)
	if err != nil {
		panic(err)
	}
	return dat
}


func (sub GetSubscriptionResponse) ErrorMessage() string {
	return sub.ARBGetSubscriptionResponse.Messages.Message.Text
}


func CurrentTime() string {
	current_time := time.Now().UTC()
	return current_time.Format("2006-01-02")
}


type GetSubscriptionResponse struct {
	ARBGetSubscriptionResponse struct {
					   RefID string `json:"refId"`
					   Messages struct {
							 ResultCode string `json:"resultCode"`
							 Message struct {
									    Code string `json:"code"`
									    Text string `json:"text"`
								    } `json:"message"`
						 } `json:"messages"`
					   Subscription struct {
							 Name string `json:"name"`
							 PaymentSchedule struct {
								      Interval struct {
										       Length string `json:"length"`
										       Unit string `json:"unit"`
									       } `json:"interval"`
								      StartDate string `json:"startDate"`
								      TotalOccurrences string `json:"totalOccurrences"`
								      TrialOccurrences string `json:"trialOccurrences"`
							      } `json:"paymentSchedule"`
							 Amount string `json:"amount"`
							 TrialAmount string `json:"trialAmount"`
							 Status string `json:"status"`
							 Profile struct {
								      Description string `json:"description"`
								      CustomerProfileID string `json:"customerProfileId"`
								      PaymentProfile struct {
											  CustomerType string `json:"customerType"`
											  BillTo struct {
													       FirstName string `json:"firstName"`
													       LastName string `json:"lastName"`
												       } `json:"billTo"`
											  CustomerPaymentProfileID string `json:"customerPaymentProfileId"`
											  Payment struct {
													       CreditCard struct {
																  CardNumber string `json:"cardNumber"`
																  ExpirationDate string `json:"expirationDate"`
															  } `json:"creditCard"`
												       } `json:"payment"`
										  } `json:"paymentProfile"`
							      } `json:"profile"`
						 } `json:"subscription"`
				   } `json:"ARBGetSubscriptionResponse"`
}