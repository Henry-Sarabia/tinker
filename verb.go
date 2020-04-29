package tinker

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
)

// Verb contains a verb and its synonyms.
type Verb struct {
	Name     string   `json:"name"`
	Synonyms []string `json:"variants"`
}

// RandVariant returns a random variant from the Verb's Variants slice.
func (v Verb) RandVariant() string {
	i := rand.Intn(len(v.Synonyms))
	return v.Synonyms[i]
}

func loadVerbs(filenames ...string) map[string]Verb {
	loaded := make(map[string]Verb)

	verbs := readVerbs(filenames...)
	for _, v := range verbs {
		loaded[v.Name] = v
	}

	return loaded
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
