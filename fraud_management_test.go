package AuthorizeCIM

import (
	"testing"
)

func TestGetUnsettledTransactions(t *testing.T) {
	transactions := UnsettledBatchList()

	t.Log("Count Unsettled: ", transactions.Count())
	t.Log(transactions.List())
}

func TestApproveTransaction(t *testing.T) {
	oldTransaction := PreviousTransaction{
		Amount: "49.99",
		RefId:  "39824723983",
	}

	response := oldTransaction.Approve()

	if response.Approved() {
		t.Log(response.ErrorMessage())
	} else {
		t.Log(response.ErrorMessage())
	}
}
