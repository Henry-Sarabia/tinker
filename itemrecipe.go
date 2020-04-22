package tinker

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
)

// ItemRecipe outlines how to generate a specific type of item.
type ItemRecipe struct {
	Name       string            `json:"name"`
	Components []ComponentRecipe `json:"components"`
}

// ComponentRecipes returns the Recipe's ComponentRecipes according to their respective
// frequencies.
func (r *ItemRecipe) ComponentRecipes() []ComponentRecipe {
	recipes := []ComponentRecipe{}
	for _, c := range r.Components {
		if c.Frequency >= rand.Float64() {
			recipes = append(recipes, c)
		}
	}
	return recipes
}

func readItemRecipes(filenames ...string) []ItemRecipe {
	var recs []ItemRecipe

	for _, fn := range filenames {
		f, err := ioutil.ReadFile(fn)
		if err != nil {
			log.Fatal(err)
		}

		rec := []ItemRecipe{}
		if err := json.Unmarshal(f, &rec); err != nil {
			log.Fatal(err)
		}

		recs = append(recs, rec...)
	}
	return recs
}
