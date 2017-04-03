package independentreserveclient

import "github.com/RyanCarrier/cryptoclientgo"

//Domain is the default domain for the independentreserve api
const Domain = "https://api.independentreserve.com"

//Multiplier is the amoutn each decimal is multiplied by to avoid floating point
// calculation errors
const Multiplier = cryptoclientgo.Multiplier
