package tinker

import (
	"encoding/json"
	"fmt"
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

// description returns a written description of an AttributeRecipe.
func (a AttributeRecipe) description(bank map[string]AttributeRecipe) (string, error) {
	desc := a.randomBase()

	// base case
	if pbChainPrefix < rand.Float64() {
		return desc, nil
	}

	pfx, err := a.randomPrefix(bank)
	if err != nil {
		return "", err
	}

	p, err := pfx.description(bank)
	if err != nil {
		return "", err
	}

	desc = p + " " + desc
	return desc, nil
}

// randomBase returns a random base attribute based on the constant rarity probabilities.
func (a AttributeRecipe) randomBase() string {
	pb := rand.Float64()

	switch {
	case pb < pbCommon:
		return randString(a.Common)
	case pb < pbUncommon+pbCommon:
		return randString(a.Uncommon)
	case pb < pbRare+pbUncommon+pbCommon:
		return randString(a.Rare)
	default:
		return "you should never see this"
	}
}

// randomPrefix returns a random prefix so long as it can be found in the provided AttributeRecipe bank.
func (a AttributeRecipe) randomPrefix(bank map[string]AttributeRecipe) (AttributeRecipe, error) {
	if len(a.PrefixNames) <= 0 {
		return AttributeRecipe{}, nil
	}

	n := randString(a.PrefixNames)
	pfx, ok := bank[n]
	if !ok {
		return AttributeRecipe{}, fmt.Errorf("cannot find prefix AttributeRecipe '%s'", n)
	}

	return pfx, nil
}

// attribute returns an Attribute according to the AttributeRecipe.
func (a AttributeRecipe) attribute(bank map[string]AttributeRecipe) (Attribute, error) {
	d, err := a.description(bank)
	if err != nil {
		return Attribute{}, err
	}

	return Attribute{
		Name:        a.Name,
		Description: d,
	}, nil
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
