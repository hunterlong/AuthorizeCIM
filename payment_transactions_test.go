package AuthorizeCIM

import (
	"testing"
)

var previousAuth string
var previousCharged string

func TestChargeCard(t *testing.T) {
	newTransaction := NewTransaction{
		Amount: RandomNumber(5, 99) + ".90",
		CreditCard: CreditCard{
			CardNumber:     "4007000000027",
			ExpirationDate: "10/23",
		},
		BillTo: &BillTo{
			FirstName:   "okokk",
			LastName:    "okok",
			Address:     "1111 white ct",
			City:        "los angeles",
			Country:     "USA",
			Zip:         "29292",
			PhoneNumber: "8885555555",
		},
	}
	response := newTransaction.Charge()
	if response.Approved() {
		previousCharged = response.TransactionID()
		t.Log("#", response.TransactionID(), "Transaction was CHARGED $", newTransaction.Amount, "\n")
		t.Log("AVS Result Code: ", response.AVS().avsResultCode+"\n")
		t.Log("AVS ACVV Result Code: ", response.AVS().cavvResultCode+"\n")
		t.Log("AVS CVV Result Code: ", response.AVS().cvvResultCode+"\n")
	} else {
		t.Log(response.ErrorMessage(), "\n")
		t.SkipNow()
	}
}

func TestDeclinedChargeCard(t *testing.T) {
	newTransaction := NewTransaction{
		Amount: RandomNumber(5, 99) + ".90",
		CreditCard: CreditCard{
			CardNumber:     "4007000000027",
			ExpirationDate: "10/23",
		},
		BillTo: &BillTo{
			FirstName:   "Fraud",
			LastName:    "User",
			Address:     "1337 Yolo Ln.",
			City:        "Beverly Hills",
			State:       "CA",
			Country:     "USA",
			Zip:         "46282",
			PhoneNumber: "8885555555",
		},
	}
	response := newTransaction.Charge()
	if response.Approved() {
		t.Fail()
	} else {
		previousCharged = response.TransactionID()
		t.Log("#", response.TransactionID(), "Transaction was CHARGED $", newTransaction.Amount, "\n")
		t.Log("AVS Result Code: ", response.AVS().avsResultCode+"\n")
		t.Log("AVS ACVV Result Code: ", response.AVS().cavvResultCode+"\n")
		t.Log("AVS CVV Result Code: ", response.AVS().cvvResultCode+"\n")
	}
}

func TestAuthOnlyCard(t *testing.T) {
	newTransaction := NewTransaction{
		Amount: RandomNumber(5, 99) + ".00",
		CreditCard: CreditCard{
			CardNumber:     "4012888818888",
			ExpirationDate: "10/27",
		},
	}
	response := newTransaction.AuthOnly()
	if response.Approved() {
		previousAuth = response.TransactionID()
		t.Log("#", response.TransactionID(), "Transaction was AUTHORIZED $", newTransaction.Amount, "\n")
	} else {
		t.Log(response.ErrorMessage(), "\n")
		t.Log(response.Message(), "\n")
		t.SkipNow()
	}
}

func TestCaptureAuth(t *testing.T) {
	oldTransaction := PreviousTransaction{
		Amount: RandomNumber(5, 99) + ".99",
		RefId:  previousAuth,
	}
	response := oldTransaction.Capture()
	if response.Approved() {
		t.Log("#", response.TransactionID(), "Transaction was CAPTURED $", oldTransaction.Amount, "\n")
	} else {
		t.Log(response.ErrorMessage(), "\n")
		t.Log(response.Message(), "\n")
		t.SkipNow()
	}
}

func TestChargeCardChannel(t *testing.T) {
	newTransaction := NewTransaction{
		Amount: RandomNumber(5, 99) + ".00",
		CreditCard: CreditCard{
			CardNumber:     "4012888818888",
			ExpirationDate: "10/24",
		},
		AuthCode: "RANDOMAUTHCODE",
	}
	response := newTransaction.Charge()

	if response.Approved() {
		previousAuth = response.TransactionID()
		t.Log("#", response.TransactionID(), "Transaction was Charged Through Channel (AuthCode) $", newTransaction.Amount, "\n")
	} else {
		t.Log(response.ErrorMessage(), "\n")
		t.Log(response.Message(), "\n")
		t.SkipNow()
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
	response := newTransaction.Refund()
	if response.Approved() {
		t.Log("#", response.TransactionID(), "Transaction was REFUNDED $", newTransaction.Amount, "\n")
	} else {
		t.Log(response.ErrorMessage(), "\n")
		t.Log(response.Message(), "\n")
		t.SkipNow()
	}
}

func TestVoidCard(t *testing.T) {
	newTransaction := PreviousTransaction{
		RefId: previousCharged,
	}
	response := newTransaction.Void()
	if response.Approved() {
		t.Log("#", response.TransactionID(), "Transaction was VOIDED $", newTransaction.Amount, "\n")
	} else {
		t.Log(response.ErrorMessage(), "\n")
		t.Log(response.Message(), "\n")
		t.SkipNow()
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
		Amount: RandomNumber(5, 99) + ".00",
	}

	response := newTransaction.ChargeProfile(customer)

	if response.Approved() {
		t.Log("#", response.TransactionID(), "Customer was Charged $", newTransaction.Amount, "\n")
	} else {
		t.Log(response.ErrorMessage(), "\n")
		t.Log(response.Message(), "\n")
		t.SkipNow()
	}
}
