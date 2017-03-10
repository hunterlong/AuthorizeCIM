package AuthorizeCIM

import (
	"encoding/json"
)

func GetPaymentProfileIds(month string, method string) GetCustomerPaymentProfileListResponse {
	action := GetCustomerPaymentProfileListRequest{
		GetCustomerPaymentProfileList: GetCustomerPaymentProfileList{
			MerchantAuthentication: GetAuthentication(),
			SearchType:             method,
			Month:                  month,
			Sorting: Sorting{
				OrderBy:         "id",
				OrderDescending: "false",
			},
			Paging: Paging{
				Limit:  "1000",
				Offset: "1",
			},
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}
	response := SendRequest(jsoned)
	var dat GetCustomerPaymentProfileListResponse
	err = json.Unmarshal(response, &dat)
	if err != nil {
		panic(err)
	}
	return dat
}

func (profile Profile) CreateProfile() CustomProfileResponse {
	response, _ := CreateProfile(profile)
	return response
}

func (profile Profile) CreateShipping() CreateCustomerShippingAddressResponse {
	response, _ := CreateShipping(profile)
	return response
}

func (customer Customer) Info() GetCustomerProfileResponse {
	response, _ := GetProfile(customer)
	return response
}

func (customer Customer) Validate() ValidateCustomerPaymentProfileResponse {
	response, _ := ValidatePaymentProfile(customer)
	return response
}

func (customer Customer) DeleteProfile() MessagesResponse {
	response, _ := DeleteProfile(customer)
	return response
}

func (customer Customer) DeletePaymentProfile() MessagesResponse {
	response, _ := DeletePaymentProfile(customer)
	return response
}

func (customer Customer) DeleteShippingProfile() MessagesResponse {
	response, _ := DeleteShippingProfile(customer)
	return response
}

func (payment CustomerPaymentProfile) Add() CustomerPaymentProfileResponse {
	response, _ := CreatePaymentProfile(payment)
	return response
}

func (response GetCustomerProfileResponse) PaymentProfiles() []GetPaymentProfiles {
	return response.Profile.PaymentProfiles
}

func (response GetCustomerProfileResponse) ShippingProfiles() []GetShippingProfiles {
	return response.Profile.ShippingProfiles
}

func (response GetCustomerProfileResponse) Subscriptions() []string {
	return response.SubscriptionIds
}

func (profile Profile) UpdateProfile() MessagesResponse {
	response, _ := UpdateProfile(profile)
	return response
}

func GetProfileIds() ([]string, interface{}) {
	action := GetCustomerProfileIdsRequest{
		CustomerProfileIdsRequest: CustomerProfileIdsRequest{
			MerchantAuthentication: GetAuthentication(),
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}
	response := SendRequest(jsoned)
	var dat CustomerProfileIdsResponse
	err = json.Unmarshal(response, &dat)
	if err != nil {
		panic(err)
	}
	return dat.Ids, err
}

func ValidatePaymentProfile(customer Customer) (ValidateCustomerPaymentProfileResponse, interface{}) {
	action := ValidateCustomerPaymentProfileRequest{
		ValidateCustomerPaymentProfile: ValidateCustomerPaymentProfile{
			MerchantAuthentication:   GetAuthentication(),
			CustomerProfileID:        customer.ID,
			CustomerPaymentProfileID: customer.PaymentID,
			ValidationMode:           testMode,
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}
	response := SendRequest(jsoned)
	var dat ValidateCustomerPaymentProfileResponse
	err = json.Unmarshal(response, &dat)
	if err != nil {
		panic(err)
	}
	return dat, err
}

func GetProfile(customer Customer) (GetCustomerProfileResponse, interface{}) {
	action := CustomerProfileRequest{
		GetCustomerProfile: GetCustomerProfile{
			MerchantAuthentication: GetAuthentication(),
			CustomerProfileID:      customer.ID,
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}
	response := SendRequest(jsoned)
	var dat GetCustomerProfileResponse
	err = json.Unmarshal(response, &dat)
	if err != nil {
		panic(err)
	}
	return dat, err
}

func CreateProfile(profile Profile) (CustomProfileResponse, interface{}) {
	action := CreateCustomerProfileRequest{
		CreateCustomerProfile: CreateCustomerProfile{
			MerchantAuthentication: GetAuthentication(),
			Profile:                profile,
			ValidationMode:         testMode,
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}

	response := SendRequest(jsoned)
	var dat CustomProfileResponse
	err = json.Unmarshal(response, &dat)
	if err != nil {
		panic(err)
	}
	return dat, err
}

func CreateShipping(profile Profile) (CreateCustomerShippingAddressResponse, interface{}) {
	action := CreateCustomerShippingAddressRequest{
		CreateCustomerShippingAddress: CreateCustomerShippingAddress{
			MerchantAuthentication: GetAuthentication(),
			Address:                profile.Shipping,
			CustomerProfileID:      profile.CustomerProfileId,
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}
	response := SendRequest(jsoned)
	var dat CreateCustomerShippingAddressResponse
	err = json.Unmarshal(response, &dat)
	if err != nil {
		panic(err)
	}
	return dat, err
}

func UpdateProfile(profile Profile) (MessagesResponse, interface{}) {
	action := UpdateCustomerProfileRequest{
		UpdateCustomerProfile: UpdateCustomerProfile{
			MerchantAuthentication: GetAuthentication(),
			Profile:                profile,
		},
	}
	dat, err := MessageResponder(action)
	return dat, err
}

func DeleteProfile(customer Customer) (MessagesResponse, interface{}) {
	action := DeleteCustomerProfileRequest{
		DeleteCustomerProfile: DeleteCustomerProfile{
			MerchantAuthentication: GetAuthentication(),
			CustomerProfileID:      customer.ID,
		},
	}
	dat, err := MessageResponder(action)
	return dat, err
}

func MessageResponder(d interface{}) (MessagesResponse, interface{}) {
	jsoned, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	response := SendRequest(jsoned)
	var dat MessagesResponse
	err = json.Unmarshal(response, &dat)
	if err != nil {
		panic(err)
	}
	return dat, err
}

func DeletePaymentProfile(customer Customer) (MessagesResponse, interface{}) {
	action := DeleteCustomerPaymentProfileRequest{
		DeleteCustomerPaymentProfile: DeleteCustomerPaymentProfile{
			MerchantAuthentication:   GetAuthentication(),
			CustomerProfileID:        customer.ID,
			CustomerPaymentProfileID: customer.PaymentID,
		},
	}
	dat, err := MessageResponder(action)
	return dat, err
}

func DeleteShippingProfile(customer Customer) (MessagesResponse, interface{}) {
	action := DeleteCustomerShippingProfileRequest{
		DeleteCustomerShippingProfile: DeleteCustomerShippingProfile{
			MerchantAuthentication: GetAuthentication(),
			CustomerProfileID:      customer.ID,
			CustomerShippingID:     customer.ShippingID,
		},
	}
	dat, err := MessageResponder(action)
	return dat, err
}

func CreatePaymentProfile(profile CustomerPaymentProfile) (CustomerPaymentProfileResponse, interface{}) {
	action := CreateCustomerPaymentProfile{
		CreateCustomerPaymentProfileRequest: CreateCustomerPaymentProfileRequest{
			MerchantAuthentication: GetAuthentication(),
			CustomerProfileID:      profile.CustomerProfileID,
			PaymentProfile: PaymentProfile{
				BillTo:                profile.PaymentProfile.BillTo,
				Payment:               profile.PaymentProfile.Payment,
				DefaultPaymentProfile: profile.PaymentProfile.DefaultPaymentProfile,
			},
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}
	response := SendRequest(jsoned)
	var dat CustomerPaymentProfileResponse
	err = json.Unmarshal(response, &dat)
	if err != nil {
		panic(err)
	}
	return dat, err
}

type CreateCustomerProfileRequest struct {
	CreateCustomerProfile CreateCustomerProfile `json:"createCustomerProfileRequest"`
}

type CreateCustomerProfile struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	Profile                Profile                `json:"profile"`
	ValidationMode         string                 `json:"validationMode"`
}

type CustomerProfiler struct {
	CustomerProfileID         string `json:"customerProfileId,omitempty"`
	CustomerPaymentProfileID  string `json:"customerPaymentProfileId,omitempty"`
	CustomerShippingProfileID string `json:"customerAddressId,omitempty"`
}

type Profile struct {
	MerchantCustomerID string           `json:"merchantCustomerId,omitempty"`
	Description        string           `json:"description,omitempty"`
	Email              string           `json:"email,omitempty"`
	CustomerProfileId  string           `json:"customerProfileId,omitempty"`
	PaymentProfiles    *PaymentProfiles `json:"paymentProfiles,omitempty"`
	Shipping           *Address         `json:"address,omitempty"`
}

type PaymentProfiles struct {
	CustomerType string  `json:"customerType,omitempty"`
	Payment      Payment `json:"payment,omitempty"`
}

type CustomProfileResponse struct {
	CustomerProfileID             string        `json:"customerProfileId"`
	CustomerPaymentProfileIDList  []string      `json:"customerPaymentProfileIdList"`
	CustomerShippingAddressIDList []interface{} `json:"customerShippingAddressIdList"`
	ValidationDirectResponseList  []string      `json:"validationDirectResponseList"`
	MessagesResponse
}

type CustomerProfileRequest struct {
	GetCustomerProfile GetCustomerProfile `json:"getCustomerProfileRequest"`
}

type GetCustomerProfile struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileID      string                 `json:"customerProfileId"`
}

type GetCustomerProfileResponse struct {
	Profile struct {
		PaymentProfiles    []GetPaymentProfiles  `json:"paymentProfiles,omitempty"`
		ShippingProfiles   []GetShippingProfiles `json:"shipToList,omitempty"`
		CustomerProfileID  string                `json:"customerProfileId"`
		MerchantCustomerID string                `json:"merchantCustomerId,omitempty"`
		Description        string                `json:"description,omitempty"`
		Email              string                `json:"email,omitempty"`
	} `json:"profile"`
	SubscriptionIds []string `json:"subscriptionIds"`
	MessagesResponse
}

type DeleteCustomerPaymentProfileRequest struct {
	DeleteCustomerPaymentProfile DeleteCustomerPaymentProfile `json:"deleteCustomerPaymentProfileRequest"`
}

type DeleteCustomerPaymentProfile struct {
	MerchantAuthentication   MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileID        string                 `json:"customerProfileId"`
	CustomerPaymentProfileID string                 `json:"customerPaymentProfileId"`
}

type DeleteCustomerShippingProfileRequest struct {
	DeleteCustomerShippingProfile DeleteCustomerShippingProfile `json:"deleteCustomerShippingAddressRequest"`
}

type DeleteCustomerShippingProfile struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileID      string                 `json:"customerProfileId"`
	CustomerShippingID     string                 `json:"customerAddressId"`
}

type GetShippingProfiles struct {
	CustomerAddressID string `json:"customerAddressId"`
	FirstName         string `json:"firstName,omitempty"`
	LastName          string `json:"lastName,omitempty"`
	Company           string `json:"company,omitempty"`
	Address           string `json:"address,omitempty"`
	City              string `json:"city,omitempty"`
	State             string `json:"state,omitempty"`
	Zip               string `json:"zip,omitempty"`
	Country           string `json:"country,omitempty"`
	PhoneNumber       string `json:"phoneNumber,omitempty"`
}

type GetPaymentProfiles struct {
	CustomerPaymentProfileID string `json:"customerPaymentProfileId"`
	Payment                  struct {
		CreditCard struct {
			CardNumber     string `json:"cardNumber"`
			ExpirationDate string `json:"expirationDate"`
		} `json:"creditCard"`
	} `json:"payment"`
	CustomerType string `json:"customerType"`
	BillTo       struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"billTo"`
}

type GetCustomerProfileIdsRequest struct {
	CustomerProfileIdsRequest CustomerProfileIdsRequest `json:"getCustomerProfileIdsRequest"`
}

type CustomerProfileIdsRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
}

type CustomerProfileIdsResponse struct {
	Ids      []string `json:"ids"`
	Messages struct {
		ResultCode string `json:"resultCode"`
		Message    []struct {
			Code string `json:"code"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"messages"`
}

type UpdateCustomerProfileRequest struct {
	UpdateCustomerProfile UpdateCustomerProfile `json:"updateCustomerProfileRequest"`
}

type UpdateCustomerProfile struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	Profile                Profile                `json:"profile"`
}

type DeleteCustomerProfileRequest struct {
	DeleteCustomerProfile DeleteCustomerProfile `json:"deleteCustomerProfileRequest"`
}

type DeleteCustomerProfile struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileID      string                 `json:"customerProfileId"`
}

type MessagesResponse struct {
	Messages struct {
		ResultCode string `json:"resultCode"`
		Message    []struct {
			Code string `json:"code"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"messages"`
}

type MessageResponse struct {
	ResultCode string `json:"resultCode"`
	Message    struct {
		Code string `json:"code"`
		Text string `json:"text"`
	} `json:"message"`
}

type CustomerPaymentProfile struct {
	CustomerProfileID string         `json:"customerProfileId"`
	PaymentProfile    PaymentProfile `json:"paymentProfile"`
}

type CreateCustomerPaymentProfile struct {
	CreateCustomerPaymentProfileRequest CreateCustomerPaymentProfileRequest `json:"createCustomerPaymentProfileRequest"`
}

type CreateCustomerPaymentProfileRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileID      string                 `json:"customerProfileId"`
	PaymentProfile         PaymentProfile         `json:"paymentProfile"`
}

type PaymentProfile struct {
	BillTo                BillTo  `json:"billTo"`
	Payment               Payment `json:"payment"`
	DefaultPaymentProfile string  `json:"defaultPaymentProfile"`
}

type CustomerPaymentProfileResponse struct {
	CustomerProfileId        string `json:"customerProfileId"`
	CustomerPaymentProfileID string `json:"customerPaymentProfileId"`
	ValidationDirectResponse string `json:"validationDirectResponse"`
	MessagesResponse
}

type GetCustomerPaymentProfileListRequest struct {
	GetCustomerPaymentProfileList GetCustomerPaymentProfileList `json:"getCustomerPaymentProfileListRequest"`
}

type GetCustomerPaymentProfileList struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	SearchType             string                 `json:"searchType"`
	Month                  string                 `json:"month"`
	Sorting                Sorting                `json:"sorting"`
	Paging                 Paging                 `json:"paging"`
}

type GetCustomerPaymentProfileListResponse struct {
	GetCustomerPaymentProfileList struct {
		MessagesResponse
		TotalNumInResultSet string `json:"totalNumInResultSet"`
		PaymentProfiles     struct {
			PaymentProfile []PaymentProfile `json:"paymentProfile"`
		} `json:"paymentProfiles"`
	} `json:"getCustomerPaymentProfileListResponse"`
}

type ValidateCustomerPaymentProfileRequest struct {
	ValidateCustomerPaymentProfile ValidateCustomerPaymentProfile `json:"validateCustomerPaymentProfileRequest"`
}

type ValidateCustomerPaymentProfile struct {
	MerchantAuthentication   MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileID        string                 `json:"customerProfileId"`
	CustomerPaymentProfileID string                 `json:"customerPaymentProfileId"`
	ValidationMode           string                 `json:"validationMode"`
}

type ValidateCustomerPaymentProfileResponse struct {
	DirectResponse string `json:"directResponse"`
	MessagesResponse
}

type CreateCustomerShippingAddressRequest struct {
	CreateCustomerShippingAddress CreateCustomerShippingAddress `json:"createCustomerShippingAddressRequest"`
}

type CreateCustomerShippingAddress struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileID      string                 `json:"customerProfileId,omitempty"`
	Address                *Address               `json:"address,omitempty"`
}

type CreateCustomerShippingAddressResponse struct {
	CustomerAddressID string `json:"customerAddressId,omitempty"`
	MessagesResponse
}
