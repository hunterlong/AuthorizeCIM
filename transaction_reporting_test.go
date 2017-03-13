package AuthorizeCIM

import (
	"testing"
)

func TestGetSettledBatchList(t *testing.T) {

	list := Range{
		Start: LastWeek(),
		End:   Now(),
	}

	batches := list.SettledBatch().List()

	for _, v := range batches {
		t.Log("Batch ID: ", v.BatchID, "\n")
		t.Log("Payment Method: ", v.PaymentMethod, "\n")
		t.Log("State: ", v.SettlementState, "\n")
	}

}

func TestGetTransactionList(t *testing.T) {

	list := Range{
		BatchId: "6933560",
	}

	batches := list.Transactions().List()

	for _, v := range batches {
		t.Log("Transaction ID: ", v.TransID, "\n")
		t.Log("Amount: ", v.Amount, "\n")
		t.Log("Account: ", v.AccountNumber, "\n")
	}

}

func TestGetTransactionDetails(t *testing.T) {

	newTransaction := PreviousTransaction{
		RefId: "60019493304",
	}
	response := newTransaction.Info()

	t.Log("Transaction Status: ", response.TransactionStatus, "\n")
}

func TestGetUnSettledBatchList(t *testing.T) {

	batches := UnSettledBatch().List()

	for _, v := range batches {
		t.Log("Status: ", v.TransactionStatus, "\n")
		t.Log("Amount: ", v.Amount, "\n")
		t.Log("Transaction ID: #", v.TransID, "\n")
	}

	if len(batches) == 0 {
		t.Fail()
	}

}

func TestGetBatchStatistics(t *testing.T) {

	list := Range{
		BatchId: "6933560",
	}

	batch := list.Statistics()

	t.Log("Refund Count: ", batch.RefundCount, "\n")
	t.Log("Charge Count: ", batch.ChargeCount, "\n")
	t.Log("Void Count: ", batch.VoidCount, "\n")
	t.Log("Charge Amount: ", batch.ChargeAmount, "\n")
	t.Log("Refund Amount: ", batch.RefundAmount, "\n")

}

func TestGetMerchantDetails(t *testing.T) {

	info := GetMerchantDetails()
	t.Log("Test Mode: ", info.IsTestMode, "\n")
	t.Log("Merchant Name: ", info.MerchantName, "\n")
	t.Log("Gateway ID: ", info.GatewayID, "\n")
}
