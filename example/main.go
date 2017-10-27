package main

import (
	"github.com/RyanCarrier/cryptoclientgo"
	"github.com/RyanCarrier/independentreserveclient"
)

func main() {
	c := independentreserveclient.New("a", "b")
	ccc := cryptoclientgo.New(c)
}
