package independentreserveclient

//Client is the default independent reserve client
type Client struct {
	apikey  string
	private string
	Domain string
}


//New generates a new independent reserve client
func New(apikey, private string) (Client, error) {
	c := Client{apikey, private, Domain, []string{}, []string{}}
	var err error
	c.primaryCurrencies, err = c.getCurrencies(true)
	if err != nil {
		return c, err
	}
	c.secondaryCurrencies, err = c.getCurrencies(false)
	return c, err
}
