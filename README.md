# Authorize.net CIM for golang

Give your Go Language applications the ability to store and retrieve credit cards from Authorize.net.


## Usage

* Import package
```
go get github.com/hunterlong/authorize_cim
```

## Features
* Store Billing Profile (credit card) on Authorize.net using CIM
* Fetch customers billing profiles in a simple array
* Process transactions using customers billing profile
* Delete and Updating billing profiles


## Examples
Below you'll find the useful functions to get you up and running in no time!

#### Create new Customer Account
```
customer_info := AuthUser{"54","email@domain.com","Test Account"}
new_customer, err := CreateCustomerProfile(customer_info)
// outputs new user profile ID
```

#### Create Payment Profile for Customer
```
address := Address{FirstName: "Test", LastName: "User", Address: "1234 Road St", City: "City Name", State:" California", Zip: "93063", Country: "USA", PhoneNumber: "5555555555"}
credit_card := CreditCard{CardNumber: "4111111111111111", ExpirationDate: "2020-12"}
profile_id := "53"
CreateCustomerBillingProfile(profile_id, address, credit_card)
```
#### Get Customers stored billing accounts
```
profile_id := "30089822"
billing_accounts := GetCustomerProfile(profile_id)
```

#### Delete Customer Profile
```
profile_id := "30089822"
billing_accounts = DeleteCustomerProfile(profile_id)
```

#### Get detailed information about a billing profile from customer
```
profile_id := "30089822"
payment_id := "1200079812"
stored_card := GetCustomerPaymentProfile(profile_id,payment_id)
```

#### Delete customers billing profile
```
profile_id := "30089822"
payment_id := "1200079812"
DeleteCustomerPaymentProfile(profile_id,payment_id)
```

#### Update a single billing profile with new information
```
new_address := Address{FirstName: "Test", LastName: "User", Address: "1234 Road St", City: "City Name", State:" California", Zip: "93063", Country: "USA", PhoneNumber: "5555555555"}
credit_card := CreditCard{CardNumber: "4111111111111111", ExpirationDate: "2020-12"}
profile_id := "53"
payment_id := "416"
UpdateCustomerPaymentProfile(profile_id,payment_id,new_address,credit_card)
```

#### Create a transaction that will be charged on customers billing profile
```
item := LineItem{ItemID: "55", Name: "item here", Description: "its simple", Quantity: "1", UnitPrice: "9.58"}
profile_id := "53"
payment_id := "416"
amount := "18.99"
CreateTransaction(profile_id,payment_id,item,amount)
```

# ToDo
* Cleanup code and add organized structs for function inputs


# License
This golang package is release under MIT license. This software gets reponses in JSON, but Authorize.net currently says "JSON Support is in BETA, please contact us if you intend to use it in production." Make sure you test in sandbox mode!
