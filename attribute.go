package tinker

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
)

const (
	// Must add up to 1.
	chanceCommon   float64 = 0.75
	chanceUncommon float64 = 0.20
	chanceRare     float64 = 0.5

	chancePrefixChain float64 = 0.5
)

// AttributeRecipe describes a physical attribute.
type AttributeRecipe struct {
	Name        string   `json:"name"`
	Common      []string `json:"common"`
	Uncommon    []string `json:"uncommon"`
	Rare        []string `json:"rare"`
	PrefixNames []string `json:"prefix_names"`
}

// Attribute describes a specific physical attribute.
type Attribute struct {
	Name        string
	Description string
}

func (a AttributeRecipe) resolve(bank map[string]AttributeRecipe) string {
	var s string

	r := rand.Float64()
	switch {
	case r < chanceCommon:
		s = randomString(a.Common)
	case r < chanceUncommon+chanceCommon:
		s = randomString(a.Uncommon)
	case r < chanceRare+chanceUncommon+chanceCommon:
		s = randomString(a.Rare)
	}

	if len(a.PrefixNames) > 0 && rand.Float64() > chancePrefixChain {
		n := randomString(a.PrefixNames)
		s = bank[n].resolve(bank) + " " + s
	}

	return s
}

// attribute returns an Attribute according to the AttributeRecipe.
func (a AttributeRecipe) attribute(bank map[string]AttributeRecipe) Attribute {
	return Attribute{
		Name:        a.Name,
		Description: a.resolve(bank),
	}
}

func loadAttributeRecipes(filenames ...string) map[string]AttributeRecipe {
	out := make(map[string]AttributeRecipe)

	recs := readAttributeRecipes(filenames...)
	for _, r := range recs {
		out[r.Name] = r
	}

	return out
}

func readAttributeRecipes(filenames ...string) []AttributeRecipe {
	var recs []AttributeRecipe

	for _, fn := range filenames {
		f, err := ioutil.ReadFile(fn)
		if err != nil {
			log.Fatal(err)
		}

		rec := []AttributeRecipe{}
		if err := json.Unmarshal(f, &rec); err != nil {
			log.Fatal(err)
		}

		recs = append(recs, rec...)
	}
	return recs
}
