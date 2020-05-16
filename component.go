package tinker

import (
	"math/rand"
)

// Component holds specific data generated from a Recipe.
type Component struct {
	Name       string
	Properties []Property
}

// RandProperty returns a randomly selected Property from the Component's Properties.
func (c Component) RandProperty() Property {
	i := rand.Intn(len(c.Properties))
	return c.Properties[i]
}
