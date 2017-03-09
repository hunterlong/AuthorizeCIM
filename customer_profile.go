package AuthorizeCIM

import (
	"encoding/json"
	"fmt"
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

func (response MessageResponse) Approved() bool {
	if response.Messages.ResultCode == "Ok" {
		return true
	}
	return false
}

func (response CustomerPaymentProfileResponse) Approved() bool {
	if response.Messages.ResultCode == "Ok" {
		return true
	}
	return false
}

func (response CreateCustomerShippingAddressResponse) Approved() bool {
	if response.Messages.ResultCode == "Ok" {
		return true
	}
	return false
}

func (response CustomerPaymentProfileResponse) ErrorMessage() string {
	return response.Messages.Message[0].Text
}

func (response MessageResponse) ErrorMessage() string {
	return response.Messages.Message[0].Text
}

func (response CustomProfileResponse) Approved() bool {
	if response.Messages.ResultCode == "Ok" {
		return true
	}
	return false
}

func (response ValidateCustomerPaymentProfileResponse) Approved() bool {
	if response.Messages.ResultCode == "Ok" {
		return true
	}
	return false
}

func (response ValidateCustomerPaymentProfileResponse) ErrorMessage() string {
	return response.Messages.Message[0].Text
}

func (response CustomProfileResponse) ErrorMessage() string {
	return response.Messages.Message[0].Text
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

func (customer Customer) Delete() MessageResponse {
	response, _ := DeleteProfile(customer)
	return response
}

func (payment CustomerPaymentProfile) Add() CustomerPaymentProfileResponse {
	response, _ := CreatePaymentProfile(payment)
	return response
}

func (response GetCustomerProfileResponse) PaymentProfiles() []GetPaymentProfiles {
	return response.Profile.PaymentProfiles
}

func (profile Profile) UpdateProfile() MessageResponse {
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
	fmt.Println(string(jsoned))

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

func UpdateProfile(profile Profile) (MessageResponse, interface{}) {
	action := UpdateCustomerProfileRequest{
		UpdateCustomerProfile: UpdateCustomerProfile{
			MerchantAuthentication: GetAuthentication(),
			Profile:                profile,
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsoned))
	response := SendRequest(jsoned)
	var dat MessageResponse
	err = json.Unmarshal(response, &dat)
	if err != nil {
		panic(err)
	}
	return dat, err
}

func DeleteProfile(customer Customer) (MessageResponse, interface{}) {
	action := DeleteCustomerProfileRequest{
		DeleteCustomerProfile: DeleteCustomerProfile{
			MerchantAuthentication: GetAuthentication(),
			CustomerProfileID:      customer.ID,
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}
	response := SendRequest(jsoned)
	var dat MessageResponse
	err = json.Unmarshal(response, &dat)
	if err != nil {
		panic(err)
	}
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
	fmt.Println(string(jsoned))
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
	CustomerProfileID             string          `json:"customerProfileId"`
	CustomerPaymentProfileIDList  []string        `json:"customerPaymentProfileIdList"`
	CustomerShippingAddressIDList []interface{}   `json:"customerShippingAddressIdList"`
	ValidationDirectResponseList  []string        `json:"validationDirectResponseList"`
	Messages                      ProfileMessages `json:"messages"`
}

type ProfileMessages struct {
	ResultCode string `json:"resultCode"`
	Message    []struct {
		Code string `json:"code"`
		Text string `json:"text"`
	} `json:"message"`
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
		PaymentProfiles    []GetPaymentProfiles `json:"paymentProfiles"`
		CustomerProfileID  string               `json:"customerProfileId"`
		MerchantCustomerID string               `json:"merchantCustomerId"`
		Description        string               `json:"description"`
		Email              string               `json:"email"`
	} `json:"profile"`
	SubscriptionIds []string `json:"subscriptionIds"`
	Messages        struct {
		ResultCode string `json:"resultCode"`
		Message    []struct {
			Code string `json:"code"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"messages"`
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

type MessageResponse struct {
	Messages struct {
		ResultCode string `json:"resultCode"`
		Message    []struct {
			Code string `json:"code"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"messages"`
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
	Messages                 struct {
		ResultCode string `json:"resultCode"`
		Message    []struct {
			Code string `json:"code"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"messages"`
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
		Messages struct {
			ResultCode string `json:"resultCode"`
			Message    struct {
				Code string `json:"code"`
				Text string `json:"text"`
			} `json:"message"`
		} `json:"messages"`
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
	Messages       struct {
		ResultCode string `json:"resultCode"`
		Message    []struct {
			Code string `json:"code"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"messages"`
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
	Messages          struct {
		ResultCode string `json:"resultCode"`
		Message    []struct {
			Code string `json:"code"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"messages"`
}
