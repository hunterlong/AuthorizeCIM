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
* Fetch customers billing profiles in a simple["array"]
* Process transactions using customers stored credit card
* Delete and Updating billing profiles
* Add Shipping Profiles into user accounts
* Delete a customers entire account

![alt tag](http://pichoster.net/images/2016/05/30/githubbreakerJKAya.jpg)

## Examples
Below you'll find useful functions to get you up and running in no time!

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


## Included API References

Payment Transactions

:white_check_mark: chargeCard

:white_check_mark: authorizeCard

:white_check_mark: capturePreviousCard

:white_check_mark: captureAuthorizedCardChannel

:white_check_mark: refundTransaction

:white_check_mark: voidTransaction

:white_medium_square: updateSplitTenderGround

:white_medium_square: debitBankAccount

:white_medium_square: creditBankAccount

:white_medium_square: chargeCustomerProfile

:white_medium_square: chargeTokenCard

:white_medium_square: creditAcceptPaymentTransaction

:white_medium_square: getAccessPaymentPage

:white_medium_square: getHostedPaymentPageRequest

Fraud Management

:white_check_mark: getUnsettledTransactionListRequest

:white_medium_square: updateHeldTransactionRequest

Recurring Billing

:white_check_mark: ARBCreateSubscriptionRequest

:white_check_mark: ARBGetSubscriptionRequest

:white_check_mark: ARBGetSubscriptionStatusRequest

:white_check_mark: ARBUpdateSubscriptionRequest

:white_check_mark: ARBCancelSubscriptionRequest

:white_check_mark: ARBGetSubscriptionListRequest

Customer Profile

:white_check_mark: createCustomerProfileRequest

:white_check_mark: getCustomerProfileRequest

:white_check_mark: getCustomerProfileIdsRequest

:white_check_mark: updateCustomerProfileRequest

:white_check_mark: deleteCustomerProfileRequest

:white_medium_square: createCustomerPaymentProfileRequest

:white_medium_square: getCustomerPaymentProfileRequest

:white_medium_square: getCustomerPaymentProfileListRequest

:white_medium_square: validateCustomerPaymentProfileRequest

:white_medium_square: updateCustomerPaymentProfileRequest

:white_medium_square: deleteCustomerPaymentProfileRequest

:white_medium_square: createCustomerShippingAddressRequest

:white_medium_square: getCustomerShippingAddressRequest

:white_medium_square: updateCustomerShippingAddressRequest

:white_medium_square: deleteCustomerShippingAddressRequest

:white_medium_square: getHostedProfilePageRequest

:white_medium_square: createCustomerProfileFromTransactionRequest

Transaction Reporting

:white_medium_square: getSettledBatchListRequest

:white_medium_square: getTransactionListRequest


# License
This golang package is release under MIT license. This software gets reponses in JSON, but Authorize.net currently says "JSON Support is in BETA, please contact us if you intend to use it in production." Make sure you test in sandbox mode!


