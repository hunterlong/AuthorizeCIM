package AuthorizeCIM

import (
	"testing"
)

func TestGetUnsettledTransactions(t *testing.T) {
	transactions := UnsettledBatchList()

	t.Log("Count Unsettled: ", transactions.Count)
	t.Log(transactions.List())
}
