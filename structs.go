package AuthorizeCIM

type MerchantAuthentication struct {
	Name string `json:"name"`
	TransactionKey string `json:"transactionKey"`
}



type AuthUser struct {
	Uuid	string
	Email	string
	Description 	string
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

type AuthorizeNetTest struct {
	AuthenticateTestRequest AuthenticateTestRequest `json:"authenticateTestRequest"`
}

type AuthenticateTestRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
}

type DoCreateTransaction struct {
	CreateTransactionRequest CreateTransactionRequest `json:"createTransactionRequest"`
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