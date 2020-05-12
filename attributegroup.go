//+build
package tinker

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// AttrGroupRecipe contains a list of related groups.
type AttrGroupRecipe struct {
	Name           string   `json:"name"`
	AttributeNames []string `json:"attribute_names"`
}

func loadAttrGroupRecipes(filenames ...string) map[string]AttrGroupRecipe {
	var out map[string]AttrGroupRecipe

	recs := readAttrGroupRecipes(filenames...)
	for _, r := range recs {
		out[r.Name] = r
	}

	return out
}

func readAttrGroupRecipes(filenames ...string) []AttrGroupRecipe {
	var recs []AttrGroupRecipe

	for _, fn := range filenames {
		f, err := ioutil.ReadFile(fn)
		if err != nil {
			log.Fatal(err)
		}

		rec := []AttrGroupRecipe{}
		if err := json.Unmarshal(f, &rec); err != nil {
			log.Fatal(err)
		}

		recs = append(recs, rec...)
	}
	return recs
}
