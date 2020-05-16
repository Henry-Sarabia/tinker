package tinker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"

	"github.com/Henry-Sarabia/article"
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
func (ar AttributeRecipe) description(bank map[string]AttributeRecipe) ([]string, error) {
	desc := []string{ar.randomBase()}

	// base case
	if pbChainPrefix < rand.Float64() {
		return desc, nil
	}

	pfx, err := ar.randomPrefix(bank)
	if err != nil {
		return nil, err
	}

	p, err := pfx.description(bank)
	if err != nil {
		return nil, err
	}

	desc = append(p, desc...)
	return desc, nil
}

// randomBase returns a random base attribute based on the constant rarity probabilities.
func (ar AttributeRecipe) randomBase() string {
	pb := rand.Float64()

	switch {
	case pb < pbCommon:
		return randString(ar.Common)
	case pb < pbUncommon+pbCommon:
		return randString(ar.Uncommon)
	case pb < pbRare+pbUncommon+pbCommon:
		return randString(ar.Rare)
	default:
		return "you should never see this"
	}
}

// randomPrefix returns a random prefix so long as it can be found in the provided AttributeRecipe bank.
func (ar AttributeRecipe) randomPrefix(bank map[string]AttributeRecipe) (AttributeRecipe, error) {
	if len(ar.PrefixNames) <= 0 {
		return AttributeRecipe{}, nil
	}

	n := randString(ar.PrefixNames)
	pfx, ok := bank[n]
	if !ok {
		return AttributeRecipe{}, fmt.Errorf("cannot find prefix AttributeRecipe '%s'", n)
	}

	return pfx, nil
}

// attribute returns an Attribute according to the AttributeRecipe.
func (ar AttributeRecipe) attribute(bank map[string]AttributeRecipe) (Attribute, error) {
	d, err := ar.description(bank)
	if err != nil {
		return Attribute{}, err
	}

	return Attribute{
		Name:        d[len(d)-1],
		Description: strings.TrimSpace(strings.Join(d, " ")),
		Article:     article.Indefinite(d[0]),
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
