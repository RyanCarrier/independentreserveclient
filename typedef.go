package independentreserveclient

//IndependentReserveRequest  is a request interface used to aid in signing the
// request
type IndependentReserveRequest interface {
	SignFormat() (string, error)
	SetNonce(int64)
	GetURL() string
	SetSignature(string)
}

//RequestTemplate is a template for all independent reserve requests
type RequestTemplate struct {
	Nonce     int64
	Signature string
	URL       string
}

//SetNonce sets the nonce, this needs to be done as late as possible when
// setting up the request
func (rt *RequestTemplate) SetNonce(i int64) {
	rt.Nonce = i
}

//SetSignature sets the signature field
func (rt *RequestTemplate) SetSignature(s string) {
	rt.Signature = s
}

//GetURL returns the URL destination for the request
func (rt *RequestTemplate) GetURL() string {
	return rt.URL
}
