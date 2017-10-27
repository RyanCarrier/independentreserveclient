package independentreserveclient

//Client is the default independent reserve client
type Client struct {
	apikey  string
	private string

	Domain string
}

//New generates a new independent reserve client
func New(apikey, private string) Client {
	return Client{apikey, private, Domain}
}
