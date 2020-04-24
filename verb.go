package tinker

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
)

// Verb describes a verb and its equivalent variants.
type Verb struct {
	Name     string   `json:"name"`
	Variants []string `json:"variants"`
}

// RandVariant returns a random variant from the Verb's Variants slice.
func (v Verb) RandVariant() string {
	return v.Variants[rand.Intn(len(v.Variants))]
}

func loadVerbs(filenames ...string) map[string]Verb {
	out := make(map[string]Verb)

	verbs := readVerbs(filenames...)
	for _, v := range verbs {
		out[v.Name] = v
	}

	return out
}

func readVerbs(filenames ...string) []Verb {
	var verbs []Verb

	for _, fn := range filenames {
		f, err := ioutil.ReadFile(fn)
		if err != nil {
			log.Fatal(err)
		}

		v := []Verb{}
		if err := json.Unmarshal(f, &v); err != nil {
			log.Fatal(err)
		}

		verbs = append(verbs, v...)
	}

	return verbs
}
