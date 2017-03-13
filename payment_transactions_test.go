package AuthorizeCIM

import (
	"testing"
)

var previousAuth string
var previousCharged string
var heldTransactionId string

func TestChargeCard(t *testing.T) {
	newTransaction := NewTransaction{
		Amount: RandomNumber(5, 99) + ".90",
		CreditCard: CreditCard{
			CardNumber:     "4007000000027",
			ExpirationDate: "08/" + RandomNumber(20, 27),
		},
	}
	response := newTransaction.Charge()
	if response.Approved() {
		previousCharged = response.TransactionID()
		t.Log("#", response.TransactionID(), "Transaction was CHARGED $", newTransaction.Amount, "\n")
	} else {
		t.Log(response.ErrorMessage(), "\n")
		t.Log(response.Message(), "\n")
		t.SkipNow()
	}
}

func TestAVSDeclinedChargeCard(t *testing.T) {
	newTransaction := NewTransaction{
		Amount: RandomNumber(5, 99) + ".75",
		CreditCard: CreditCard{
			CardNumber:     "5424000000000015",
			ExpirationDate: "08/" + RandomNumber(20, 27),
		},
		BillTo: &BillTo{
			FirstName:   RandomString(7),
			LastName:    RandomString(9),
			Address:     "1111 white ct",
			City:        "los angeles",
			State:       "CA",
			Zip:         "46205",
			Country:     "USA",
			PhoneNumber: "8885555555",
		},
	}
	response := newTransaction.Charge()

	if response.AVS().avsResultCode == "N" {
		t.Log("#", response.TransactionID(), "AVS Transaction was DECLINED due to AVS Code. $", newTransaction.Amount, "\n")
		t.Log("AVS Result Text: ", response.AVS().Text(), "\n")
		t.Log("AVS Result Code: ", response.AVS().avsResultCode, "\n")
		t.Log("AVS ACVV Result Code: ", response.AVS().cavvResultCode, "\n")
		t.Log("AVS CVV Result Code: ", response.AVS().cvvResultCode, "\n")
	} else {
		t.Log(response.ErrorMessage(), "\n")
		t.Log(response.Message(), "\n")
		t.Fail()
	}
}

func TestAVSChargeCard(t *testing.T) {
	newTransaction := NewTransaction{
		Amount: RandomNumber(5, 99) + ".75",
		CreditCard: CreditCard{
			CardNumber:     "4012888818888",
			ExpirationDate: "08/" + RandomNumber(20, 27),
		},
		BillTo: &BillTo{
			FirstName:   RandomString(7),
			LastName:    RandomString(9),
			Address:     "1111 green ct",
			City:        "los angeles",
			State:       "CA",
			Zip:         "46203",
			Country:     "USA",
			PhoneNumber: "8885555555",
		},
	}
	response := newTransaction.Charge()

	if response.Approved() {
		heldTransactionId = response.TransactionID()
	}

	if response.Held() {
		t.Log("Transaction is being Held for Review", "\n")
	}

	if response.AVS().avsResultCode == "E" {
		t.Log("#", response.TransactionID(), "AVS Transaction was CHARGED is now on HOLD$", newTransaction.Amount, "\n")
		t.Log("AVS Result Text: ", response.AVS().Text(), "\n")
		t.Log("AVS Result Code: ", response.AVS().avsResultCode, "\n")
		t.Log("AVS ACVV Result Code: ", response.AVS().cavvResultCode, "\n")
		t.Log("AVS CVV Result Code: ", response.AVS().cvvResultCode, "\n")
	} else {
		t.Log(response.ErrorMessage(), "\n")
		t.Log(response.Message(), "\n")
		t.Fail()
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
			FirstName:   "Declined",
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
		t.Log("#", response.TransactionID(), "Transaction was DECLINED!!!", "\n")
		t.Log(response.Message(), "\n")
		t.Log("AVS Result Text: ", response.AVS().Text(), "\n")
		t.Log("AVS Result Code: ", response.AVS().avsResultCode, "\n")
		t.Log("AVS ACVV Result Code: ", response.AVS().cavvResultCode, "\n")
		t.Log("AVS CVV Result Code: ", response.AVS().cvvResultCode, "\n")
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
		t.Log("#", response.TransactionID(), "Transaction was VOIDED", "\n")
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
