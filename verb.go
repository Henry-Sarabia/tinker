package tinker

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"

	"github.com/pkg/errors"
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

func loadVerbs(filenames ...string) (map[string]Verb, error) {
	loaded := make(map[string]Verb)

	verbs, err := readVerbs(filenames...)
	if err != nil {
		return nil, err
	}
	for _, v := range verbs {
		loaded[v.Name] = v
	}

	return loaded, nil
}

func readVerbs(filenames ...string) ([]Verb, error) {
	var verbs []Verb

	for _, fn := range filenames {
		f, err := ioutil.ReadFile(fn)
		if err != nil {
			errors.Wrapf(err, "cannot read file '%s'", fn)
		}

		v := []Verb{}
		if err := json.Unmarshal(f, &v); err != nil {
			errors.Wrapf(err, "cannot unmarshal Verbs from file '%s'", fn)
		}

		verbs = append(verbs, v...)
	}

	return verbs, nil
}
