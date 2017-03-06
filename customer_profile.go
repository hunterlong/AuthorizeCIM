package AuthorizeCIM

import (
	"encoding/json"
	"fmt"
)

func (response MessageResponse) Approved() bool {
	if response.Messages.ResultCode == "Ok" {
		return true
	}
	return false
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

func (response CustomProfileResponse) ErrorMessage() string {
	return response.Messages.Message[0].Text
}

func (profile Profile) Create() CustomProfileResponse {
	response, _ := CreateProfile(profile)
	return response
}

func (customer Customer) Info() GetCustomerProfileResponse {
	response, _ := GetProfile(customer)
	return response
}

func (customer Customer) Delete() MessageResponse {
	response, _ := DeleteProfile(customer)
	return response
}

func (response GetCustomerProfileResponse) PaymentProfiles() []GetPaymentProfiles {
	return response.Profile.PaymentProfiles
}

func (profile Profile) Update() MessageResponse {
	response, _ := UpdateProfile(profile)
	return response
}

func GetProfileIds() ([]string, interface{}) {
	action := GetCustomerProfileIdsRequest{
		CustomerProfileIdsRequest: CustomerProfileIdsRequest{
			MerchantAuthentication: MerchantAuthentication{
				Name:           "8v25DGQq9kf",
				TransactionKey: "5KDX8Vz3mx334aJm",
			},
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsoned))

	response := SendRequest(jsoned)
	var dat CustomerProfileIdsResponse
	fmt.Println(string(response))
	err = json.Unmarshal(response, &dat)
	if err != nil {
		panic(err)
	}
	return dat.Ids, err
}

func GetProfile(customer Customer) (GetCustomerProfileResponse, interface{}) {
	action := CustomerProfileRequest{
		GetCustomerProfile: GetCustomerProfile{
			MerchantAuthentication: GetAuthentication(),
			CustomerProfileID: customer.ID,
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsoned))

	response := SendRequest(jsoned)
	var dat GetCustomerProfileResponse
	fmt.Println(string(response))
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
			Profile:        profile,
			ValidationMode: "testMode",
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsoned))

	response := SendRequest(jsoned)
	var dat CustomProfileResponse
	fmt.Println(string(response))
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
			Profile:        profile,
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsoned))

	response := SendRequest(jsoned)
	var dat MessageResponse
	fmt.Println(string(response))
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
			CustomerProfileID: customer.ID,
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsoned))

	response := SendRequest(jsoned)
	var dat MessageResponse
	fmt.Println(string(response))
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

type Profile struct {
	MerchantCustomerID string          `json:"merchantCustomerId,omitempty"`
	Description        string          `json:"description,omitempty"`
	Email              string          `json:"email,omitempty"`
	CustomerProfileId  string          `json:"customerProfileId,omitempty"`
	PaymentProfiles    *PaymentProfiles `json:"paymentProfiles,omitempty"`
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
	Messages                      ProfileMessages `json:"messages"`
}

type ProfileMessages  struct {
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
	MerchantAuthentication MerchantAuthentication`json:"merchantAuthentication"`
	Profile Profile `json:"profile"`
}

type DeleteCustomerProfileRequest struct {
	DeleteCustomerProfile DeleteCustomerProfile `json:"deleteCustomerProfileRequest"`
}

type DeleteCustomerProfile struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	CustomerProfileID string `json:"customerProfileId"`
}

type MessageResponse struct {
	Messages struct {
			 ResultCode string `json:"resultCode"`
			 Message []struct {
				 Code string `json:"code"`
				 Text string `json:"text"`
			 } `json:"message"`
		 } `json:"messages"`
}