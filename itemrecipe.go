package tinker

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"

	"github.com/pkg/errors"
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

func readItemRecipes(filenames ...string) ([]ItemRecipe, error) {
	var rcps []ItemRecipe

	for _, fn := range filenames {
		f, err := ioutil.ReadFile(fn)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot read file '%s'", fn)
		}

		rcp := []ItemRecipe{}
		if err := json.Unmarshal(f, &rcp); err != nil {
			return nil, errors.Wrapf(err, "cannot unmarshal ItemRecipes from file '%s'", fn)
		}

		rcps = append(rcps, rcp...)
	}
	return rcps, nil
}
