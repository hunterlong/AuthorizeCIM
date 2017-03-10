package AuthorizeCIM

func (transx TransactionResponse) TransactionID() string {
	return transx.Response.TransID
}

func (transx TransactionResponse) Message() string {
	return transx.Response.Message.Messages.Message[0].Text
}

func (transx TransactionResponse) AVS() AVS {
	out := AVS{
		avsResultCode:  transx.Response.AvsResultCode,
		cvvResultCode:  transx.Response.CvvResultCode,
		cavvResultCode: transx.Response.CavvResultCode,
	}
	return out
}

type AVS struct {
	avsResultCode  string
	cvvResultCode  string
	cavvResultCode string
}

type TransxReponse interface {
	ErrorMessage()
	Approved()
}

func (r MessagesResponse) ErrorMessage() string {
	return r.Messages.Message[0].Text
}

func (r MessagesResponse) Approved() bool {
	if r.Messages.ResultCode == "Ok" {
		return true
	}
	return false
}
