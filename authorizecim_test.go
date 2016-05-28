package AuthorizeCIM

import (
	"testing"
	"time"
	"math/rand"
)


var test_profile_id string
var test_payment_id string

func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func TestAPIAccess(t *testing.T) {

}

func TestUserCreation(t *testing.T) {
	customer_info := AuthUser{"54",RandomString(9)+"@random.com","Test Account"}
	newuser, err := CreateCustomerProfile(customer_info)
	if err!=nil{
		t.Fail()
	}
	test_profile_id = newuser
	t.Log("User was created "+test_profile_id+"\n")
}


func TestCreatePaymentProfile(t *testing.T){
	address := Address{FirstName: "Test", LastName: "User", Address: "1234 Road St", City: "City Name", State:" California", Zip: "93063", Country: "USA", PhoneNumber: "5555555555"}
	credit_card := CreditCard{CardNumber: "4111111111111111", ExpirationDate: "2020-12"}
	new_payment_id, err := CreateCustomerBillingProfile(test_profile_id, credit_card, address)
	if err!=nil{
		t.Fail()
	}
	test_payment_id = new_payment_id
	t.Log("User Payment Profile created "+test_payment_id+"\n")
}


func TestProfileTransaction(t *testing.T) {
	item := LineItem{ItemID: "S0897", Name: "New Product", Description: "brand new", Quantity: "1", UnitPrice: "5.50"}
	amount := "14.43"
	tranx, err := CreateTransaction(test_profile_id, test_payment_id, item, amount)
	if err!=nil{
		t.Fail()
	}
	var tranx_id string
	fixtransx, _ := tranx["transactionResponse"].(map[string]interface{})
	if tranx["transactionResponse"]==nil {
		tranx_id = "0"
		t.Fail()
		t.Log("Transaction has failed! "+tranx_id+"\n")
	} else {
		tranx_id = fixtransx["transId"].(string)
		t.Log("Received Transaction: "+tranx_id+"\n")
	}
}


func TestDeleteCustomerPaymentProfile(t *testing.T) {
	DeleteCustomerPaymentProfile(test_profile_id, test_payment_id)
	t.Log("User Payment Profile was deleted: "+test_payment_id+"\n")
}

func TestDeleteCustomerProfile(t *testing.T){
	DeleteCustomerProfile(test_profile_id)
	t.Log("Customer Profile was deleted: "+test_profile_id+"\n")
}
