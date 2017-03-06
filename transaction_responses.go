package AuthorizeCIM

func (transx TransactionResponse) Approved() bool {
	if transx.Response.ResponseCode == "1" {
		return true
	}
	return false
}

func (transx TransactionResponse) TransactionID() string {
	return transx.Response.TransID
}

func (transx TransactionResponse) ListMessages() []Message {
	return transx.Messages.Message
}

func (transx TransactionResponse) ErrorMessage() []AuthNetErrors {
	return transx.Response.Errors
}
