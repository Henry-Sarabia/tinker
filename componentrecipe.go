package tinker

import "math/rand"

// ComponentRecipe describes how to generate a physical part of an Item.
type ComponentRecipe struct {
	Name       string           `json:"name"`
	Frequency  float64          `json:"frequency"`
	Properties []PropertyRecipe `json:"properties"`
}

// PropertyRecipes returns the Component's PropertyRecipes according to their respective
// frequencies.
func (c *ComponentRecipe) PropertyRecipes() []PropertyRecipe {
	recipes := []PropertyRecipe{}
	for _, p := range c.Properties {
		if p.Frequency >= rand.Float64() {
			recipes = append(recipes, p)
		}
	}
	return recipes
}
