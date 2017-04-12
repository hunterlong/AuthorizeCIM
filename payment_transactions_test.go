package AuthorizeCIM

import (
	"testing"
)

var previousAuth string
var previousCharged string

func TestChargeCard(t *testing.T) {
	newTransaction := NewTransaction{
		Amount: "15.90",
		CreditCard: CreditCard{
			CardNumber:     "4007000000027",
			ExpirationDate: "10/23",
		},
	}
	response, err := newTransaction.Charge()
	if err != nil {
		t.Fail()
	}
	if response.Approved() {
		previousCharged = response.TransactionID()
		t.Log("#", response.TransactionID(), "Transaction was CHARGED $", newTransaction.Amount, "\n")
		t.Log("AVS Result Code: ", response.AVS().avsResultCode+"\n")
		t.Log("AVS ACVV Result Code: ", response.AVS().cavvResultCode+"\n")
		t.Log("AVS CVV Result Code: ", response.AVS().cvvResultCode+"\n")
	} else {
		t.Log(response.ErrorMessage(), "\n")
	}
}

func TestAuthOnlyCard(t *testing.T) {
	newTransaction := NewTransaction{
		Amount: "100.00",
		CreditCard: CreditCard{
			CardNumber:     "4012888818888",
			ExpirationDate: "10/27",
		},
	}
	response, err := newTransaction.AuthOnly()
	if err != nil {
		t.Fail()
	}

	if response.Approved() {
		previousAuth = response.TransactionID()
		t.Log("#", response.TransactionID(), "Transaction was AUTHORIZED $", newTransaction.Amount, "\n")
	} else {
		t.Log(response.ErrorMessage(), "\n")
	}
}

func TestCaptureAuth(t *testing.T) {
	oldTransaction := PreviousTransaction{
		Amount: "49.99",
		RefId:  previousAuth,
	}
	response, err := oldTransaction.Capture()
	if err != nil {
		t.Fail()
	}
	if response.Approved() {
		t.Log("#", response.TransactionID(), "Transaction was CAPTURED $", oldTransaction.Amount, "\n")
	} else {
		t.Log(response.ErrorMessage(), "\n")
	}
}

func TestChargeCardChannel(t *testing.T) {
	newTransaction := NewTransaction{
		Amount: "38.00",
		CreditCard: CreditCard{
			CardNumber:     "4012888818888",
			ExpirationDate: "10/24",
		},
		AuthCode: "RANDOMAUTHCODE",
	}
	response, err := newTransaction.Charge()
	if err != nil {
		t.Fail()
	}

	if response.Approved() {
		previousAuth = response.TransactionID()
		t.Log("#", response.TransactionID(), "Transaction was Charged Through Channel (AuthCode) $", newTransaction.Amount, "\n")
	} else {
		t.Log(response.ErrorMessage(), "\n")
	}
}

func TestRefundCard(t *testing.T) {
	newTransaction := NewTransaction{
		Amount: "15.00",
		CreditCard: CreditCard{
			CardNumber:     "4012888818888",
			ExpirationDate: "10/24",
		},
		RefTransId: "0392482938402",
	}
	response, err := newTransaction.Refund()
	if err != nil {
		t.Fail()
	}
	if response.Approved() {
		t.Log("#", response.TransactionID(), "Transaction was REFUNDED $", newTransaction.Amount, "\n")
	} else {
		t.Log(response.ErrorMessage(), "\n")
	}
}

func TestVoidCard(t *testing.T) {
	newTransaction := PreviousTransaction{
		RefId: previousCharged,
	}
	response, err := newTransaction.Void()
	if err != nil {
		t.Fail()
	}
	if response.Approved() {
		t.Log("#", response.TransactionID(), "Transaction was VOIDED $", newTransaction.Amount, "\n")
	} else {
		t.Log(response.ErrorMessage(), "\n")
	}
}

func TestChargeCustomerProfile(t *testing.T) {

	oldProfileId := "1810921101"
	oldPaymentId := "1805617738"

	customer := Customer{
		ID:        oldProfileId,
		PaymentID: oldPaymentId,
	}

	newTransaction := NewTransaction{
		Amount: "35.00",
	}

	response, err := newTransaction.ChargeProfile(customer)
	if err != nil {
		t.Fail()
	}

	if response.Approved() {
		t.Log("#", response.TransactionID(), "Customer was Charged $", newTransaction.Amount, "\n")
	} else {
		t.Log(response.ErrorMessage(), "\n")
	}
}
