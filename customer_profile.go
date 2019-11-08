package AuthorizeCIM

import (
	"encoding/json"
)

func GetPaymentProfileIds(month string, method string) (*GetCustomerPaymentProfileListResponse, error) {
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
		return nil, err
	}
	response, err := SendRequest(jsoned)
	var dat GetCustomerPaymentProfileListResponse
	err = json.Unmarshal(response, &dat)
	if err != nil {
		return nil, err
	}
	return &dat, err
}

func (profile Profile) CreateProfile() (*CustomProfileResponse, error) {
	response, err := CreateProfile(profile)
	return response, err
}

func (profile Profile) CreateShipping() (*CreateCustomerShippingAddressResponse, error) {
	response, err := CreateShipping(profile)
	return response, err
}

func (customer Customer) Info() (*GetCustomerProfileResponse, error) {
	response, err := GetProfile(customer)
	return response, err
}

func (customer Customer) Validate() (*ValidateCustomerPaymentProfileResponse, error) {
	response, err := ValidatePaymentProfile(customer)
	return response, err
}

func (customer Customer) DeleteProfile() (*MessagesResponse, error) {
	response, err := DeleteProfile(customer)
	return response, err
}

func (customer Customer) DeletePaymentProfile() (*MessagesResponse, error) {
	response, err := DeletePaymentProfile(customer)
	return response, err
}

func (customer Customer) DeleteShippingProfile() (*MessagesResponse, error) {
	response, err := DeleteShippingProfile(customer)
	return response, err
}

func (payment CustomerPaymentProfile) Add() (*CustomerPaymentProfileResponse, error) {
	response, err := CreatePaymentProfile(payment)
	return response, err
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

func (profile Profile) UpdateProfile() (*MessagesResponse, error) {
	response, err := UpdateProfile(profile)
	return response, err
}

func (profile Profile) UpdatePaymentProfile() (*MessagesResponse, error) {
	response, err := UpdatePaymentProfile(profile)
	return response, err
}

func (profile Profile) UpdateShippingProfile() (*MessagesResponse, error) {
	response, err := UpdateShippingProfile(profile)
	return response, err
}

func GetProfileIds() ([]string, error) {
	action := GetCustomerProfileIdsRequest{
		CustomerProfileIdsRequest: CustomerProfileIdsRequest{
			MerchantAuthentication: GetAuthentication(),
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		return []string{}, err
	}
	response, err := SendRequest(jsoned)
	var dat CustomerProfileIdsResponse
	err = json.Unmarshal(response, &dat)
	if err != nil {
		return []string{}, err
	}
	return dat.Ids, err
}

func ValidatePaymentProfile(customer Customer) (*ValidateCustomerPaymentProfileResponse, error) {
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
		return nil, err
	}
	response, err := SendRequest(jsoned)
	var dat ValidateCustomerPaymentProfileResponse
	err = json.Unmarshal(response, &dat)
	if err != nil {
		return nil, err
	}
	return &dat, err
}

func GetProfile(customer Customer) (*GetCustomerProfileResponse, error) {
	action := CustomerProfileRequest{
		GetCustomerProfile: GetCustomerProfile{
			MerchantAuthentication: GetAuthentication(),
			CustomerProfileID:      customer.ID,
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		return nil, err
	}
	response, err := SendRequest(jsoned)
	var dat GetCustomerProfileResponse
	err = json.Unmarshal(response, &dat)
	if err != nil {
		return nil, err
	}
	return &dat, err
}

func CreateProfile(profile Profile) (*CustomProfileResponse, error) {
	action := CreateCustomerProfileRequest{
		CreateCustomerProfile: CreateCustomerProfile{
			MerchantAuthentication: GetAuthentication(),
			Profile:                profile,
			ValidationMode:         "testMode",
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		return nil, err
	}

	response, err := SendRequest(jsoned)
	var dat CustomProfileResponse
	err = json.Unmarshal(response, &dat)
	if err != nil {
		return nil, err
	}
	return &dat, err
}

func CreateShipping(profile Profile) (*CreateCustomerShippingAddressResponse, error) {
	action := CreateCustomerShippingAddressRequest{
		CreateCustomerShippingAddress: CreateCustomerShippingAddress{
			MerchantAuthentication: GetAuthentication(),
			Address:                profile.Shipping,
			CustomerProfileID:      profile.CustomerProfileId,
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		return nil, err
	}
	response, err := SendRequest(jsoned)
	var dat CreateCustomerShippingAddressResponse
	err = json.Unmarshal(response, &dat)
	if err != nil {
		return nil, err
	}
	return &dat, err
}

func UpdateProfile(profile Profile) (*MessagesResponse, error) {
	action := UpdateCustomerProfileRequest{
		UpdateCustomerProfile: UpdateCustomerProfile{
			MerchantAuthentication: GetAuthentication(),
			Profile:                profile,
		},
	}
	dat, err := MessageResponder(action)
	return dat, err
}

func UpdatePaymentProfile(profile Profile) (*MessagesResponse, error) {
	action := UpdateCustomerPaymentProfileRequest{
		UpdateCustomerPaymentProfile: UpdateCustomerPaymentProfile{
			CustomerProfileID:      profile.CustomerProfileId,
			MerchantAuthentication: GetAuthentication(),
			UpPaymentProfile: UpPaymentProfile{
				BillTo:                   profile.PaymentProfiles.BillTo,
				Payment:                  profile.PaymentProfiles.Payment,
				CustomerPaymentProfileID: profile.PaymentProfileId,
			},
			ValidationMode: testMode,
		},
	}
	dat, err := MessageResponder(action)
	return dat, err
}

func UpdateShippingProfile(profile Profile) (*MessagesResponse, error) {
	action := UpdateCustomerShippingAddressRequest{
		UpdateCustomerShippingAddress: UpdateCustomerShippingAddress{
			CustomerProfileID:      profile.CustomerProfileId,
			MerchantAuthentication: GetAuthentication(),
			Address: Address{
				FirstName:         "My",
				LastName:          "Name",
				Address:           "38485 New Road ave.",
				City:              "Los Angeles",
				State:             "CA",
				Zip:               "283848",
				Country:           "USA",
				PhoneNumber:       "8885555555",
				CustomerAddressID: profile.CustomerAddressId,
			},
		},
	}
	dat, err := MessageResponder(action)
	return dat, err
}

func DeleteProfile(customer Customer) (*MessagesResponse, error) {
	action := DeleteCustomerProfileRequest{
		DeleteCustomerProfile: DeleteCustomerProfile{
			MerchantAuthentication: GetAuthentication(),
			CustomerProfileID:      customer.ID,
		},
	}
	dat, err := MessageResponder(action)
	return dat, err
}

func MessageResponder(d interface{}) (*MessagesResponse, error) {
	jsoned, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	response, err := SendRequest(jsoned)
	var dat MessagesResponse
	err = json.Unmarshal(response, &dat)
	if err != nil {
		return nil, err
	}
	return &dat, err
}

func DeletePaymentProfile(customer Customer) (*MessagesResponse, error) {
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

func DeleteShippingProfile(customer Customer) (*MessagesResponse, error) {
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

func CreatePaymentProfile(profile CustomerPaymentProfile) (*CustomerPaymentProfileResponse, error) {
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
		return nil, err
	}
	response, err := SendRequest(jsoned)
	var dat CustomerPaymentProfileResponse
	err = json.Unmarshal(response, &dat)
	if err != nil {
		return nil, err
	}
	return &dat, err
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
	PaymentProfileId   string           `json:"customerPaymentProfileId,omitempty"`
	Shipping           *Address         `json:"address,omitempty"`
	CustomerAddressId  string           `json:"customerAddressId,omitempty"`
	PaymentProfile     *PaymentProfile  `json:"paymentProfile,omitempty"`
}

type PaymentProfiles struct {
	CustomerType string  `json:"customerType,omitempty"`
	Payment      Payment `json:"payment,omitempty"`
	BillTo       *BillTo `json:"billTo,omitempty"`
	PaymentId    string  `json:"paymentProfileId,omitempty"`
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
	BillTo                *BillTo  `json:"billTo,omitempty"`
	Payment               *Payment `json:"payment,omitempty"`
	DefaultPaymentProfile string   `json:"defaultPaymentProfile,omitempty"`
	PaymentProfileId      string   `json:"paymentProfileId,omitempty"`
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

type UpdateCustomerPaymentProfileRequest struct {
	UpdateCustomerPaymentProfile UpdateCustomerPaymentProfile `json:"updateCustomerPaymentProfileRequest"`
}

type UpdateCustomerPaymentProfile struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileID      string                 `json:"customerProfileId"`
	UpPaymentProfile       UpPaymentProfile       `json:"paymentProfile"`
	ValidationMode         string                 `json:"validationMode"`
}

type UpPaymentProfile struct {
	BillTo                   *BillTo `json:"billTo"`
	Payment                  Payment `json:"payment"`
	CustomerPaymentProfileID string  `json:"customerPaymentProfileId"`
}

type UpdateCustomerShippingAddressRequest struct {
	UpdateCustomerShippingAddress UpdateCustomerShippingAddress `json:"updateCustomerShippingAddressRequest"`
}

type UpdateCustomerShippingAddress struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileID      string                 `json:"customerProfileId"`
	Address                Address                `json:"address"`
}
