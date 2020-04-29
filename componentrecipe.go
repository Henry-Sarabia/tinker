package tinker

import "math/rand"

// ComponentRecipe defines how to generate a piece of an Item.
type ComponentRecipe struct {
	Name       string           `json:"name"`
	Frequency  float64          `json:"frequency"`
	Properties []PropertyRecipe `json:"properties"`
}

// PropertyRecipes returns the Component's PropertyRecipes according to their respective
// frequencies.
func (cr *ComponentRecipe) PropertyRecipes() []PropertyRecipe {
	rcps := []PropertyRecipe{}
	for _, p := range cr.Properties {
		if p.Frequency >= rand.Float64() {
			rcps = append(rcps, p)
		}
	}
	return rcps
}
