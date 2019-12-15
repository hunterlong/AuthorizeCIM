package AuthorizeCIM

import (
	"testing"
)

func TestGetSettledBatchList(t *testing.T) {

	list := Range{
		Start: LastWeek(),
		End:   Now(),
	}

	batches, err := list.SettledBatch()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	batchList := batches.List()

	for _, v := range batchList {
		t.Log("Batch ID: ", v.BatchID, "\n")
		t.Log("Payment Method: ", v.PaymentMethod, "\n")
		t.Log("State: ", v.SettlementState, "\n")
	}

}

func TestGetTransactionList(t *testing.T) {

	list := Range{
		BatchId: "6933560",
	}

	batches, err := list.Transactions()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	batchList := batches.List()

	for _, v := range batchList {
		t.Log("Transaction ID: ", v.TransID, "\n")
		t.Log("Amount: ", v.Amount, "\n")
		t.Log("Account: ", v.AccountNumber, "\n")
	}

}

func TestGetTransactionDetails(t *testing.T) {

	newTransaction := PreviousTransaction{
		RefID: "60019493304",
	}
	response, err := newTransaction.Info()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	t.Log("Transaction Status: ", response.TransactionStatus, "\n")
}

func TestGetUnSettledBatchList(t *testing.T) {

	batches, err := UnSettledBatch()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	batchList := batches.List()

	for _, v := range batchList {
		t.Log("Status: ", v.TransactionStatus, "\n")
		t.Log("Amount: ", v.Amount, "\n")
		t.Log("Transaction ID: #", v.TransID, "\n")
	}

}

func TestGetBatchStatistics(t *testing.T) {

	list := Range{
		BatchId: "6933560",
	}

	batch, err := list.Statistics()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	t.Log("Refund Count: ", batch.RefundCount, "\n")
	t.Log("Charge Count: ", batch.ChargeCount, "\n")
	t.Log("Void Count: ", batch.VoidCount, "\n")
	t.Log("Charge Amount: ", batch.ChargeAmount, "\n")
	t.Log("Refund Amount: ", batch.RefundAmount, "\n")

}

func TestGetMerchantDetails(t *testing.T) {

	info, err := GetMerchantDetails()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	t.Log("Test Mode: ", info.IsTestMode, "\n")
	t.Log("Merchant Name: ", info.MerchantName, "\n")
	t.Log("Gateway ID: ", info.GatewayID, "\n")
}
