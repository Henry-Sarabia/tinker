package tinker

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
)

// ItemRecipe defines how to generate a specific type of item.
type ItemRecipe struct {
	Name       string            `json:"name"`
	Components []ComponentRecipe `json:"components"`
}

// ComponentRecipes returns the Recipe's ComponentRecipes according to their respective
// frequencies.
func (ir *ItemRecipe) ComponentRecipes() []ComponentRecipe {
	rcps := []ComponentRecipe{}
	for _, c := range ir.Components {
		if c.Frequency >= rand.Float64() {
			rcps = append(rcps, c)
		}
	}
	return rcps
}

func readItemRecipes(filenames ...string) []ItemRecipe {
	var rcps []ItemRecipe

	for _, fn := range filenames {
		f, err := ioutil.ReadFile(fn)
		if err != nil {
			log.Fatal(err)
		}

		rcp := []ItemRecipe{}
		if err := json.Unmarshal(f, &rcp); err != nil {
			log.Fatal(err)
		}

		rcps = append(rcps, rcp...)
	}
	return rcps
}
