package AuthorizeCIM

import (
	"testing"
)

func TestGetUnsettledTransactions(t *testing.T) {
	transactions, err := UnsettledBatchList()
	if err != nil {
		t.Fail()
	}

	t.Log("Count Unsettled: ", transactions.Count)
	t.Log(transactions.List())
}

func TestApproveTransaction(t *testing.T) {
	oldTransaction := PreviousTransaction{
		Amount: "49.99",
		RefId:  "39824723983",
	}

	response, err := oldTransaction.Approve()
	if err != nil {
		t.Fail()
	}

	if response.Approved() {
		t.Log(response.ErrorMessage())
	} else {
		t.Log(response.ErrorMessage())
	}
}

func TestDeclineTransaction2(t *testing.T) {
	oldTransaction := PreviousTransaction{
		Amount: "49.99",
		RefId:  "39824723983",
	}

	response, err := oldTransaction.Decline()
	if err != nil {
		t.Fail()
	}

	if response.Approved() {
		t.Log(response.ErrorMessage())
	} else {
		t.Log(response.ErrorMessage())
	}
}
