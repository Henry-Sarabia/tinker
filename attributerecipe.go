package tinker

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"

	"github.com/pkg/errors"
)

const (
	// Attribute rarity probabilities; must add up to 1.
	pbCommon   float64 = 0.75
	pbUncommon float64 = 0.20
	pbRare     float64 = 0.5

	// Probability of chaining an additional prefix to any given prefix.
	pbChainPrefix float64 = 0.5
)

// AttributeRecipe describes how to generate a physical attribute.
type AttributeRecipe struct {
	Name        string   `json:"name"`
	Common      []string `json:"common"`
	Uncommon    []string `json:"uncommon"`
	Rare        []string `json:"rare"`
	PrefixNames []string `json:"prefix_names"`
}

// description returns a description of the AttributeRecipe.
func (a AttributeRecipe) description(bank map[string]AttributeRecipe) string {
	var d string
	pb := rand.Float64()

	switch {
	case pb < pbCommon:
		d = randomString(a.Common)
	case pb < pbUncommon+pbCommon:
		d = randomString(a.Uncommon)
	case pb < pbRare+pbUncommon+pbCommon:
		d = randomString(a.Rare)
	}

	if len(a.PrefixNames) > 0 && rand.Float64() > pbChainPrefix {
		n := randomString(a.PrefixNames)
		d = bank[n].description(bank) + " " + d
	}

	return d
}

// attribute returns an Attribute according to the AttributeRecipe.
func (a AttributeRecipe) attribute(bank map[string]AttributeRecipe) Attribute {
	return Attribute{
		Name:        a.Name,
		Description: a.description(bank),
	}
}

func loadAttributeRecipes(filenames ...string) (map[string]AttributeRecipe, error) {
	loaded := make(map[string]AttributeRecipe)

	rcps, err := readAttributeRecipes(filenames...)
	if err != nil {
		return nil, err
	}
	for _, r := range rcps {
		loaded[r.Name] = r
	}

	return loaded, nil
}

func readAttributeRecipes(filenames ...string) ([]AttributeRecipe, error) {
	var rcps []AttributeRecipe

	for _, fn := range filenames {
		f, err := ioutil.ReadFile(fn)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot read file '%s'", fn)
		}

		rcp := []AttributeRecipe{}
		if err := json.Unmarshal(f, &rcp); err != nil {
			return nil, errors.Wrapf(err, "cannot unmarshal AttributeRecipes from file '%s'", fn)
		}

		rcps = append(rcps, rcp...)
	}
	return rcps, nil
}
