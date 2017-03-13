![alt tag](http://pichoster.net/images/2016/05/30/authcimyjIpi.jpg)

# Authorize.net CIM, AIM, and ARB for Go Language
[![Build Status](https://travis-ci.org/hunterlong/AuthorizeCIM.svg?branch=master)](https://travis-ci.org/hunterlong/AuthorizeCIM)  [![Code Climate](https://lima.codeclimate.com/github/Hunterlong/AuthorizeCIM/badges/gpa.svg)](https://lima.codeclimate.com/github/Hunterlong/AuthorizeCIM) [![Coverage Status](https://coveralls.io/repos/github/hunterlong/AuthorizeCIM/badge.svg?branch=master)](https://coveralls.io/github/hunterlong/AuthorizeCIM?branch=master) [![GoDoc](https://godoc.org/github.com/hunterlong/AuthorizeCIM?status.svg)](https://godoc.org/github.com/hunterlong/AuthorizeCIM) [![Go Report Card](https://goreportcard.com/badge/github.com/hunterlong/AuthorizeCIM)](https://goreportcard.com/report/github.com/hunterlong/AuthorizeCIM)

Give your Go Language applications the ability to store and retrieve credit cards from Authorize.net CIM, AIM, and ARB API.
This golang package lets you create recurring subscriptions, AUTH only transactions, voids, refunds, and other functionality connected to the Authorize.net API.

# Features
* [AIM Payment Transactions](https://github.com/hunterlong/AuthorizeCIM#payment-transactions)
* [CIM Customer Information Manager](https://github.com/hunterlong/AuthorizeCIM#customer-profile)
* [ARB Automatic Recurring Billing](https://github.com/hunterlong/AuthorizeCIM#recurring-billing) (Subscriptions)
* [Transaction Reporting](https://github.com/hunterlong/AuthorizeCIM#transaction-reporting)
* [Fraud Management](https://github.com/hunterlong/AuthorizeCIM#fraud-management)
* Creating Users Accounts based on user's unique ID and/or email address
* Store Payment Profiles (credit card) on Authorize.net using Customer Information Manager (CIM)
* Create Subscriptions (monthly, weekly, days) with Automated Recurring Billing (ARB)
* Process transactions using customers stored credit card
* Delete and Updating payment profiles
* Add Shipping Profiles into user accounts
* Delete a customers entire account
* Tests included and examples below

```go
customer := AuthorizeCIM.Customer{
        ID: "13838",
    }

customerInfo := customer.Info()

paymentProfiles := customerInfo.PaymentProfiles()
shippingProfiles := customerInfo.ShippingProfiles()
subscriptions := customerInfo.Subscriptions()
```

# Usage
* Import package
```
go get github.com/hunterlong/authorizecim
```
```go
import "github.com/hunterlong/authorizecim"
```
###### Or Shorten the Package Name
```go
import auth "github.com/hunterlong/authorizecim"
// auth.SetAPIInfo(apiName,apiKey,"test")
```

## Set Authorize.net API Keys
You can get Sandbox Access at:  https://developer.authorize.net/hello_world/sandbox/
```go
apiName := "auth_name_here"
apiKey := "auth_transaction_key_here"
AuthorizeCIM.SetAPIInfo(apiName,apiKey,"test")
// use "live" to do transactions on production server
```

![alt tag](http://pichoster.net/images/2016/05/30/githubbreakerJKAya.jpg)

## Included API References

:white_check_mark: Set API Creds
```go
func main() {

    apiName := "PQO38FSL"
    apiKey := "OQ8NFBAPA9DS"
    apiMode := "test"

    AuthorizeCIM.SetAPIInfo(apiName,apiKey,apiMode)

}
```

# Payment Transactions

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

:white_check_mark: chargeCustomerProfile
```go
customer := AuthorizeCIM.Customer{
		ID: "49587345",
		PaymentID: "84392124324",
	}

newTransaction := AuthorizeCIM.NewTransaction{
		Amount: "35.00",
	}

response := newTransaction.ChargeProfile(customer)

if response.Ok() {

}
```
:white_medium_square: chargeTokenCard

:white_medium_square: creditAcceptPaymentTransaction

:white_medium_square: getAccessPaymentPage

:white_medium_square: getHostedPaymentPageRequest

## Transaction Responses
```go
response.Ok()                   // bool
response.Approved()             // bool
response.Message()              // string
response.ErrorMessage()         // string
response.TransactionID()        // string
response.AVS()                  // [avsResultCode,cavvResultCode,cvvResultCode]
```

# Fraud Management

:white_check_mark: getUnsettledTransactionListRequest
```go
transactions := AuthorizeCIM.UnsettledBatchList()
fmt.Println("Unsettled Count: ", transactions.Count)
```

:white_check_mark: updateHeldTransactionRequest
```go
oldTransaction := AuthorizeCIM.PreviousTransaction{
		Amount: "49.99",
		RefId:  "39824723983",
	}

	response := oldTransaction.Approve()
	//response := oldTransaction.Decline()

	if response.Ok() {

	}
```

# Recurring Billing

:white_check_mark: ARBCreateSubscriptionRequest
```go
subscription := AuthorizeCIM.Subscription{
		Name:        "New Subscription",
		Amount:      "9.00",
		TrialAmount: "0.00",
		PaymentSchedule: &PaymentSchedule{
			StartDate:        CurrentDate(),
			TotalOccurrences: "9999",
			TrialOccurrences: "0",
			Interval: AuthorizeCIM.IntervalMonthly(),
		},
		Payment: &Payment{
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
###### For Intervals, you can use simple methods
```go
AuthorizeCIM.IntervalWeekly()      // runs every week (7 days)
AuthorizeCIM.IntervalMonthly()     // runs every Month
AuthorizeCIM.IntervalQuarterly()   // runs every 3 months
AuthorizeCIM.IntervalYearly()      // runs every 1 year
AuthorizeCIM.IntervalDays("15")    // runs every 15 days
AuthorizeCIM.IntervalMonths("6")   // runs every 6 months
```

:white_check_mark: ARBCreateSubscriptionRequest from Customer Profile
```go
subscription := AuthorizeCIM.Subscription{
		Name:        "New Customer Subscription",
		Amount:      "12.00",
		TrialAmount: "0.00",
		PaymentSchedule: &PaymentSchedule{
			StartDate:        CurrentDate(),
			TotalOccurrences: "9999",
			TrialOccurrences: "0",
			Interval: AuthorizeCIM.IntervalDays("15"),
		},
		Profile: &CustomerProfiler{
			CustomerProfileID: "823928379",
			CustomerPaymentProfileID: "183949200",
			//CustomerShippingProfileID: "310282443",
		},
	}

	response := subscription.Charge()

	if response.Approved() {
		newSubscriptionId = response.SubscriptionID
		fmt.Println("Customer #",response.CustomerProfileId(), " Created a New Subscription: ", response.SubscriptionID)
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

if response.Ok() {

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

# Customer Profile (CIM)

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

if response.Ok() {
    fmt.Println("New Customer Profile Created #",response.CustomerProfileID)
    fmt.Println("New Customer Payment Profile Created #",response.CustomerPaymentProfileID)
} else {
       fmt.Println(response.ErrorMessage())
   }
```

:white_check_mark: getCustomerProfileRequest
```go
customer := AuthorizeCIM.Customer{
		ID: "13838",
	}

customerInfo := customer.Info()

paymentProfiles := customerInfo.PaymentProfiles()
shippingProfiles := customerInfo.ShippingProfiles()
subscriptions := customerInfo.Subscriptions()
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

if response.Ok() {

}
```

:white_check_mark: deleteCustomerProfileRequest
```go
customer := AuthorizeCIM.Customer{
		ID: "13838",
	}

	response := customer.Delete()

if response.Ok() {

}
```

# Customer Payment Profile

:white_check_mark: createCustomerPaymentProfileRequest
```go
paymentProfile := AuthorizeCIM.CustomerPaymentProfile{
		CustomerProfileID: "32948234232",
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

if response.Ok() {

} else {
    fmt.Println(response.ErrorMessage())
}
```

:white_check_mark: getCustomerPaymentProfileRequest
```go
customer := AuthorizeCIM.Customer{
		ID: "3923482487",
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

if response.Ok() {

}
```

:white_check_mark: updateCustomerPaymentProfileRequest
```go
customer := AuthorizeCIM.Profile{
		CustomerProfileId:  "3838238293",
		PaymentProfileId: "83929382739",
		Email:              "info@updatedemail.com",
		PaymentProfiles: &PaymentProfiles{
			Payment: Payment{
				CreditCard: CreditCard{
					CardNumber: "4007000000027",
					ExpirationDate: "01/26",
				},
			},
			BillTo: &BillTo{
				FirstName:   "newname",
				LastName:    "golang",
				Address:     "2841 purple ct",
				City:        "los angeles",
				State:		  "CA",
				Zip:            "93939",
				Country:     "USA",
				PhoneNumber: "8885555555",
			},
		},
	}

response := customer.UpdatePaymentProfile()

if response.Ok() {
    fmt.Println("Customer Payment Profile was Updated")
} else {
    fmt.Println(response.ErrorMessage())
}
```
:white_check_mark: deleteCustomerPaymentProfileRequest
```go
customer := AuthorizeCIM.Customer{
		ID: "3724823472",
		PaymentID: "98238472349",
	}

response := customer.DeletePaymentProfile()

if response.Ok() {
    fmt.Println("Payment Profile was Deleted")
} else {
    fmt.Println(response.ErrorMessage())
}
```

# Customer Shipping Profile

:white_check_mark: createCustomerShippingAddressRequest
```go
customer := AuthorizeCIM.Profile{
		MerchantCustomerID: "86437",
		CustomerProfileId:  "7832642387",
		Email:              "info@emailhereooooo.com",
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

response := customer.CreateShipping()

if response.Ok() {
    fmt.Println("New Shipping Added: #",response.CustomerAddressID)
} else {
    fmt.Println(response.ErrorMessage())
}
```

:white_check_mark: getCustomerShippingAddressRequest
```go
customer := AuthorizeCIM.Customer{
		ID: "3842934233",
	}

response := customer.Info()

shippingProfiles := response.ShippingProfiles()

fmt.Println("Customer Shipping Profiles", shippingProfiles)
```
:white_check_mark: updateCustomerShippingAddressRequest
```go
customer := AuthorizeCIM.Profile{
		CustomerProfileId:  "398432389",
		CustomerAddressId: "848388438",
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

response := customer.UpdateShippingProfile()

if response.Ok() {
    fmt.Println("Shipping Profile was updated")
}
```
:white_check_mark: deleteCustomerShippingAddressRequest
```go
customer := AuthorizeCIM.Customer{
		ID: "128749382",
		ShippingID: "34892734829",
	}

	response := customer.DeleteShippingProfile()

	if response.Ok() {
		fmt.Println("Shipping Profile was Deleted")
	} else {
		fmt.Println(response.ErrorMessage())
	}
```
:white_medium_square: getHostedProfilePageRequest

:white_medium_square: createCustomerProfileFromTransactionRequest

# Transaction Reporting

:white_check_mark: getSettledBatchListRequest
```go
list := AuthorizeCIM.Range{
		Start: LastWeek(),
		End:   Now(),
	}

batches := list.SettledBatch().List()

for _, v := range batches {
    t.Log("Batch ID: ", v.BatchID, "\n")
    t.Log("Payment Method: ", v.PaymentMethod, "\n")
    t.Log("State: ", v.SettlementState, "\n")
}
```
:white_check_mark: getUnSettledBatchListRequest
```go
batches := AuthorizeCIM.UnSettledBatch().List()

for _, v := range batches {
    t.Log("Status: ",v.TransactionStatus, "\n")
    t.Log("Amount: ",v.Amount, "\n")
    t.Log("Transaction ID: #",v.TransID, "\n")
}

```
:white_check_mark: getTransactionListRequest
```go
list := AuthorizeCIM.Range{
		BatchId: "6933560",
	}

batches := list.Transactions().List()

for _, v := range batches {
    t.Log("Transaction ID: ", v.TransID, "\n")
    t.Log("Amount: ", v.Amount, "\n")
    t.Log("Account: ", v.AccountNumber, "\n")
}
```
:white_check_mark: getTransactionDetails
```go
oldTransaction := AuthorizeCIM.PreviousTransaction{
		RefId: "60019493304",
	}
response := oldTransaction.Info()

fmt.PrintLn("Transaction Status: ",response.TransactionStatus,"\n")
```
:white_check_mark: getBatchStatistics
```go
list := AuthorizeCIM.Range{
		BatchId: "6933560",
	}

batch := list.Statistics()

fmt.PrintLn("Refund Count: ", batch.RefundCount, "\n")
fmt.PrintLn("Charge Count: ", batch.ChargeCount, "\n")
fmt.PrintLn("Void Count: ", batch.VoidCount, "\n")
fmt.PrintLn("Charge Amount: ", batch.ChargeAmount, "\n")
fmt.PrintLn("Refund Amount: ", batch.RefundAmount, "\n")
```
:white_check_mark: getMerchantDetails
```go
info := AuthorizeCIM.GetMerchantDetails()

fmt.PrintLn("Test Mode: ", info.IsTestMode, "\n")
fmt.PrintLn("Merchant Name: ", info.MerchantName, "\n")
fmt.PrintLn("Gateway ID: ", info.GatewayID, "\n")
```

![alt tag](http://pichoster.net/images/2016/05/30/githubbreakerJKAya.jpg)

# ToDo
* Organize and refactor some areas
* Add Bank Account Support
* Make tests fail if transactions fail (skipping 'duplicate transaction')

### Authorize.net CIM Documentation
http://developer.authorize.net/api/reference/#customer-profiles

### Authorize.net Sandbox Access
https://developer.authorize.net/hello_world/sandbox/

# License
This golang package is release under MIT license.
Feel free to submit a Pull Request if you have updates!

*This software gets reponses in JSON, but Authorize.net currently says "JSON Support is in BETA, please contact us if you intend to use it in production."* Make sure you test in sandbox mode!


