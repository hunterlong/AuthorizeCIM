package AuthorizeCIM

type MerchantAuthentication struct {
	Name string `json:"name"`
	TransactionKey string `json:"transactionKey"`
}


type User struct {
	ID string
	Email string
	ProfileID string
	BillingProfiles interface{}
	ShippingProfiles interface{}
	Subscriptions map[string]interface{}
}


type AuthUser struct {
	Uuid	string
	Email	string
	Description 	string
}

type DeleteARBSubscriptionRequest struct {
	ARBCancelSubscriptionRequest DeleteSubscriptionRequest `json:"ARBCancelSubscriptionRequest"`
}


type DeleteSubscriptionRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	SubscriptionId string `json:"subscriptionId"`
}


type DeleteCustomerShippingAddressRequest struct {
	GetCustomerShippingAddress GetCustomerShippingAddress `json:"deleteCustomerShippingAddressRequest"`
}

type GetCustomerShippingAddressRequest struct {
	GetCustomerShippingAddress GetCustomerShippingAddress `json:"getCustomerShippingAddressRequest"`
}

type GetCustomerShippingAddress struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileId string `json:"customerProfileId"`
	CustomerShippingId string `json:"customerAddressId"`
}

type CustomerShippingAddressRequest struct {
	CustomerShippingAddress CustomerShippingAddress `json:"createCustomerShippingAddressRequest"`
}

type CustomerShippingAddress struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileId string `json:"customerProfileId"`
	Address Address		`json:"address"`
}


type TransactionRecord struct {

}

type TransactionMessages struct {

}


type TransactionDetailsRequest struct {
	TransactionDetails TransactionDetails `json:"getTransactionDetailsRequest"`
}

type TransactionDetails struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	TransId string `json:"transId"`
}


type TransactionResponse struct {
	ResponseCode string `json:"responseCode"`
	AuthCode string `json:"authCode"`
	AvsResultCode string `json:"avsResultCode"`
	CvvResultCode string `json:"cvvResultCode"`
	CavvResultCode string `json:"cavvResultCode"`
	TransID string `json:"transId"`
	RefTransID string `json:"refTransID"`
	TransHash string `json:"transHash"`
	TestRequest string `json:"testRequest"`
	AccountNumber string `json:"accountNumber"`
	AccountType string `json:"accountType"`
}

type AuthorizeNetTest struct {
	AuthenticateTestRequest AuthenticateTestRequest `json:"authenticateTestRequest"`
}

type AuthenticateTestRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
}

type DoCreateTransaction struct {
	CreateTransactionRequest CreateTransactionRequest `json:"createTransactionRequest"`
}

type VoidTransactionRequestARB struct {
	VoidTransaction VoidTransactionRequest `json:"createTransactionRequest"`
}

type CreateRefundTransactionRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	RefundTransactionRequest RefundTransactionRequest `json:"transactionRequest"`
}

type RefundTransactionRequest struct {
	TransactionType string `json:"transactionType"`
	Amount string `json:"amount"`
	Payment Payment `json:"payment"`
	TransxId string `json:"refTransId"`
}


type VoidTransactionRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	MinTrans MinTrans `json:"transactionRequest"`
}

type MinTrans struct {
	TransactionType string `json:"transactionType"`
	TransxId string `json:"refTransId"`
}


type CreateTransactionRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	RefID string `json:"refId"`
	TransactionRequest TransactionRequest `json:"transactionRequest"`
}


type TransactionRequest struct {
	TransactionType string `json:"transactionType"`
	Amount string `json:"amount"`
	TranProfile TranProfile `json:"profile"`
	LineItems LineItems `json:"lineItems"`
}

type TranProfile struct {
	CustomerProfileId string `json:"customerProfileId"`
	SubProfile SubProfile `json:"paymentProfile"`
}

type SubProfile struct {
	CustomerPaymentProfileId string `json:"paymentProfileId"`
}

type LineItems struct {
	LineItem LineItem `json:"lineItem"`
}

type LineItem struct {
	ItemID string `json:"itemId"`
	Name string `json:"name"`
	Description string `json:"description"`
	Quantity string `json:"quantity"`
	UnitPrice string `json:"unitPrice"`
}


type deleteCustomerPaymentProfileRequest struct {
	DeleteCustomerPaymentProfileRequest deleteCustomerPaymentProfile `json:"deleteCustomerPaymentProfileRequest"`
}



type deleteCustomerPaymentProfile struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileId string `json:"customerProfileId"`
	CustomerPaymentProfileId string `json:"customerPaymentProfileId"`
}



type changeCustomerPaymentProfileRequest struct {
	UpdateCustomerPaymentProfileRequest updateCustomerPaymentProfileRequest `json:"updateCustomerPaymentProfileRequest"`
}


type updateCustomerPaymentProfileRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileId string `json:"customerProfileId"`
	PaymentProfile UpdatePaymentBillingProfile `json:"paymentProfile"`
	ValidationMode string `json:"validationMode"`
}

type UpdatePaymentBillingProfile struct {
	Address Address		`json:"billTo"`
	Payment Payment `json:"payment"`
	CustomerPaymentProfileId string `json:"customerPaymentProfileId"`
}

type getCustomerPaymentProfileRequest struct {
	CustomerPaymentProfileRequest CustomerPaymentProfileRequest `json:"getCustomerPaymentProfileRequest"`
}

type CustomerPaymentProfileRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileId string `json:"customerProfileId"`
	CustomerPaymentProfileId string `json:"customerPaymentProfileId"`
}

type deleteCustomerProfile struct {
	DeleteCustomerProfileRequest deleteCustomerProfileRequest `json:"deleteCustomerProfileRequest"`
}

type deleteCustomerProfileRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileId string `json:"customerProfileId"`
}

type AllCustomerProfileIds struct {
	CustomerProfileIdsRequest getCustomerProfileIdsRequest `json:"getCustomerProfileIdsRequest"`
}


type getCustomerProfileIdsRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
}

type BillTo struct {
	Address Address		`json:"billTo"`
}

type Address struct {
	FirstName string	`json:"firstName"`
	LastName string		`json:"lastName"`
	Address string		`json:"address"`
	City string		`json:"city"`
	State string		`json:"state"`
	Zip string		`json:"zip"`
	Country string		`json:"country"`
	PhoneNumber string	`json:"phoneNumber"`
}


type CustomerProfile struct {
	CustomerProfileRequest getCustomerProfileRequest	`json:"getCustomerProfileRequest"`

}

type getCustomerProfileRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileId string `json:"customerProfileId"`
}

type BillingProfile struct {
	Address Address		`json:"billTo"`
	PaymentProfile PaymentBillingProfile `json:"paymentProfile"`
}


type Profile struct {
	MerchantCustomerID string `json:"merchantCustomerId"`
	Description string `json:"description"`
	Email string `json:"email"`
}


type PaymentBillingProfile struct {
	Address Address		`json:"billTo"`
	Payment Payment `json:"payment"`
}


type PaymentProfiles struct {
	CustomerType string `json:"customerType"`
	Payment Payment `json:"payment"`
}

type Payment struct {
	CreditCard CreditCard `json:"creditCard"`
}


type CreditCard struct {
	CardNumber string `json:"cardNumber"`
	ExpirationDate string `json:"expirationDate"`
}

type CreateCustomerProfileRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	Profile Profile `json:"profile"`
}

type CreateCustomerBillingProfileRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileId string `json:"customerProfileId"`
	Profile                PaymentBillingProfile `json:"paymentProfile"`
	ValidationMode         string `json:"validationMode"`
}

type NewCustomerBillingProfile struct {
	CreateCustomerProfileRequest CreateCustomerBillingProfileRequest `json:"createCustomerPaymentProfileRequest"`
}

type NewCustomerProfile struct {
	CreateCustomerProfileRequest CreateCustomerProfileRequest `json:"createCustomerProfileRequest"`
}

type CreateSubscriptionRequest struct {
	CreateSubscription ARBCreateSubscription `json:"ARBCreateSubscriptionRequest"`
}

type ARBCreateSubscription struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	Subscription Subscription `json:"subscription"`
}

type Subscription struct {
	Name string `json:"name"`
	PaymentSchedule PaymentSchedule `json:"paymentSchedule"`
	Amount string `json:"amount"`
	TrialAmount string `json:"trialAmount"`
	FullProfile FullProfile `json:"profile"`
}

type Interval struct {
	Length string `json:"length"`
	Unit string `json:"unit"`
}

type PaymentSchedule struct {
	Interval Interval `json:"interval"`
	StartDate string `json:"startDate"`
	TotalOccurrences string `json:"totalOccurrences"`
	TrialOccurrences string `json:"trialOccurrences"`
}


type FullProfile struct {
	CustomerProfileID string `json:"customerProfileId"`
	CustomerPaymentProfileID string `json:"customerPaymentProfileId"`
	CustomerAddressID string `json:"customerAddressId"`
}