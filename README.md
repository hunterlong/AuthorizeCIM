![alt tag](http://pichoster.net/images/2016/05/30/authcimyjIpi.jpg)

# Authorize.net CIM for golang
[![Build Status](https://travis-ci.org/hunterlong/AuthorizeCIM.svg?branch=master)](https://travis-ci.org/hunterlong/AuthorizeCIM)  [![Code Climate](https://codeclimate.com/github/Hunterlong/AuthorizeCIM/badges/gpa.svg)](https://codeclimate.com/github/hunterlong/AuthorizeCIM)  [![Coverage Status](https://coveralls.io/repos/github/Hunterlong/AuthorizeCIM/badge.svg?branch=master)](https://coveralls.io/github/hunterlong/AuthorizeCIM?branch=master)  [![GoDoc](https://godoc.org/github.com/hunterlong/AuthorizeCIM?status.svg)](https://godoc.org/github.com/hunterlong/AuthorizeCIM) [![Go Report Card](https://goreportcard.com/badge/github.com/hunterlong/AuthorizeCIM)](https://goreportcard.com/report/github.com/hunterlong/AuthorizeCIM)

Give your Go Language applications the ability to store and retrieve credit cards from Authorize.net CIM API. This golang package lets you create recurring subscriptions, AUTH only transactions, voids, refunds, and other functionality connected to the Authorize.net API.


## Usage
* Import package
```
go get github.com/hunterlong/authorizecim
```
```go
import "github.com/hunterlong/authorizecim"
```

* Set Authorize.net API Keys
```go
// Set your Authorize API name and key here
// You can get Sandbox Access at:  https://developer.authorize.net/hello_world/sandbox/
apiName := "auth_name_here"
apiKey := "auth_transaction_key_here"
AuthorizeCIM.SetAPIInfo(apiName,apiKey,"test")
// use "live" to do transactions on production server
```

## Features
* Creating Users Accounts based on user's unique ID and/or email address
* Store Billing Profiles (credit card) on Authorize.net using Customer Information Manager (CIM)
* Create Subscriptions (monthly, weekly, days) with Automated Recurring Billing (ARB)
* Process transactions using customers stored credit card
* Delete and Updating billing profiles
* Add Shipping Profiles into user accounts
* Delete a customers entire account

![alt tag](http://pichoster.net/images/2016/05/30/githubbreakerJKAya.jpg)

## Examples
Below you'll find useful functions to get you up and running in no time!

## Included API References

:white_check_mark: Set API Creds
```go
apiName := "PQO38FSL"
apiKey := "OQ8NFBAPA9DS"
apiMode := "test"

AuthorizeCIM.SetAPIInfo(apiName,apiKey,apiMode)
```

### Payment Transactions

:white_check_mark: chargeCard
```go
newTransaction := AuthorizeCIM.NewTransaction{
		Amount: "15.90",
		CreditCard: CreditCard{
			CardNumber:     "4007000000027",
			ExpirationDate: "10/23",
		},
	}
response := newTransaction.Charge()
if response.Approved() {

}
```

:white_check_mark: authorizeCard
```go
newTransaction := AuthorizeCIM.NewTransaction{
		Amount: "100.00",
		CreditCard: CreditCard{
			CardNumber:     "4012888818888",
			ExpirationDate: "10/27",
		},
	}
response := newTransaction.AuthOnly()
if response.Approved() {

}
```

:white_check_mark: capturePreviousCard
```go
oldTransaction := AuthorizeCIM.PreviousTransaction{
		Amount: "49.99",
		RefId:  "AUTHCODEHERE001",
	}
response := oldTransaction.Capture()
if response.Approved() {

}
```

:white_check_mark: captureAuthorizedCardChannel
```go
newTransaction := AuthorizeCIM.NewTransaction{
		Amount: "38.00",
		CreditCard: CreditCard{
			CardNumber:     "4012888818888",
			ExpirationDate: "10/24",
		},
		AuthCode: "YOURAUTHCODE",
	}
response := newTransaction.Charge()
if response.Approved() {

}
```

:white_check_mark: refundTransaction
```go
newTransaction := AuthorizeCIM.NewTransaction{
		Amount: "15.00",
		CreditCard: CreditCard{
			CardNumber:     "4012888818888",
			ExpirationDate: "10/24",
		},
		RefTransId: "0392482938402",
	}
response := newTransaction.Refund()
if response.Approved() {

}
```

:white_check_mark: voidTransaction
```go
oldTransaction := AuthorizeCIM.PreviousTransaction{
		RefId: "3987324293834",
	}
response := oldTransaction.Void()
if response.Approved() {

}
```

:white_medium_square: updateSplitTenderGround

:white_medium_square: debitBankAccount

:white_medium_square: creditBankAccount

:white_medium_square: chargeCustomerProfile

:white_medium_square: chargeTokenCard

:white_medium_square: creditAcceptPaymentTransaction

:white_medium_square: getAccessPaymentPage

:white_medium_square: getHostedPaymentPageRequest

### Fraud Management

:white_check_mark: getUnsettledTransactionListRequest
```go
transactions := AuthorizeCIM.UnsettledBatchList()
fmt.Println("Unsettled Count: ", transactions.Count)
```

:white_medium_square: updateHeldTransactionRequest

### Recurring Billing

:white_check_mark: ARBCreateSubscriptionRequest
```go
subscription := AuthorizeCIM.Subscription{
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
		Payment: Payment{
			CreditCard: CreditCard{
				CardNumber:     "4007000000027",
				ExpirationDate: "10/23",
			},
		},
		BillTo: &BillTo{
			FirstName: "Test",
			LastName:  "User",
		},
	}

response := subscription.Charge()

if response.Approved() {
    fmt.Println("New Subscription ID: ",response.SubscriptionID)
}
```

:white_check_mark: ARBGetSubscriptionRequest
```go
sub := AuthorizeCIM.SetSubscription{
		Id: "2973984693",
	}

subscriptionInfo := sub.Info()
```

:white_check_mark: ARBGetSubscriptionStatusRequest
```go
sub := AuthorizeCIM.SetSubscription{
		Id: "2973984693",
	}

subscriptionInfo := sub.Status()

fmt.Println("Subscription ID has status: ",subscriptionInfo.Status)
```

:white_check_mark: ARBUpdateSubscriptionRequest
```go
subscription := AuthorizeCIM.Subscription{
		Payment: Payment{
			CreditCard: CreditCard{
				CardNumber:     "5424000000000015",
				ExpirationDate: "06/25",
			},
		},
		SubscriptionId: newSubscriptionId,
	}

response := subscription.Update()

if response.Approved() {

}
```

:white_check_mark: ARBCancelSubscriptionRequest
```go
sub := AuthorizeCIM.SetSubscription{
		Id: "2973984693",
	}

subscriptionInfo := sub.Cancel()

fmt.Println("Subscription ID has been canceled: ", sub.Id, "\n")
```

:white_check_mark: ARBGetSubscriptionListRequest
```go
inactive := AuthorizeCIM.SubscriptionList("subscriptionInactive")
fmt.Println("Amount of Inactive Subscriptions: ", inactive.Count())

active := AuthorizeCIM.SubscriptionList("subscriptionActive")
fmt.Println("Amount of Active Subscriptions: ", active.Count())
```

### Customer Profile

:white_check_mark: createCustomerProfileRequest
```go
customer := AuthorizeCIM.Profile{
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
    fmt.Println("New Customer Profile Created #", response.CustomerProfileID)

}
```

:white_check_mark: getCustomerProfileRequest
```go
customer := AuthorizeCIM.Customer{
		ID: "13838",
	}

response := customer.Info()
```

:white_check_mark: getCustomerProfileIdsRequest
```go
profiles, _ := AuthorizeCIM.GetProfileIds()
fmt.Println(profiles)
```

:white_check_mark: updateCustomerProfileRequest
```go
customer := AuthorizeCIM.Profile{
		MerchantCustomerID: "13838",
		CustomerProfileId: "13838",
		Description: "Updated Account",
		Email:       "info@updatedemail.com",
	}

	response := customer.Update()

if response.Approved() {

}
```

:white_check_mark: deleteCustomerProfileRequest
```go
customer := AuthorizeCIM.Customer{
		ID: "13838",
	}

	response := customer.Delete()

if response.Approved() {

}
```

### Customer Payment Profile

:white_check_mark: createCustomerPaymentProfileRequest
```go
paymentProfile := AuthorizeCIM.CustomerPaymentProfile{
		CustomerProfileID: newCustomerProfileId,
		PaymentProfile: PaymentProfile{
			BillTo: BillTo{
				FirstName: "okokk",
				LastName: "okok",
				Address: "1111 white ct",
				City: "los angeles",
				Country: "USA",
				PhoneNumber: "8885555555",
			},
			Payment: Payment{
				CreditCard: CreditCard{
					CardNumber: "5424000000000015",
					ExpirationDate: "04/22",
				},
			},
			DefaultPaymentProfile: "true",
		},
	}

	response := paymentProfile.Add()

	if response.Approved() {

	}
```

:white_check_mark: getCustomerPaymentProfileRequest
```go
customer := AuthorizeCIM.Customer{
		ID: newCustomerProfileId,
	}

	response := customer.Info()

	paymentProfiles := response.PaymentProfiles()
```

:white_check_mark: getCustomerPaymentProfileListRequest
```go
profileIds := AuthorizeCIM.GetPaymentProfileIds("2017-03","cardsExpiringInMonth")
```

:white_check_mark: validateCustomerPaymentProfileRequest
```go
customerProfile := AuthorizeCIM.Customer{
		ID: "127723778",
		PaymentID: "984583934",
	}

	response := customerProfile.Validate()

	if response.Approved() {

	}
```

:white_medium_square: updateCustomerPaymentProfileRequest

:white_medium_square: deleteCustomerPaymentProfileRequest

Customer Shipping Profile

:white_medium_square: createCustomerShippingAddressRequest

:white_medium_square: getCustomerShippingAddressRequest

:white_medium_square: updateCustomerShippingAddressRequest

:white_medium_square: deleteCustomerShippingAddressRequest

:white_medium_square: getHostedProfilePageRequest

:white_medium_square: createCustomerProfileFromTransactionRequest

Transaction Reporting

:white_medium_square: getSettledBatchListRequest

:white_medium_square: getTransactionListRequest


![alt tag](http://pichoster.net/images/2016/05/30/githubbreakerJKAya.jpg)

# ToDo
* Make cleaner maps for outputs
* Functions to search Subscriptions (active, expired, etc)
* Void and Refund Transactions
* Add Bank Account Support
* Authorize Only methods

### Authorize.net CIM Documentation
http://developer.authorize.net/api/reference/#customer-profiles

### Authorize.net Sandbox Access
https://developer.authorize.net/hello_world/sandbox/

# License
This golang package is release under MIT license. This software gets reponses in JSON, but Authorize.net currently says "JSON Support is in BETA, please contact us if you intend to use it in production." Make sure you test in sandbox mode!


