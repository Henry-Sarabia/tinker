package tinker

import "github.com/Henry-Sarabia/article"

// Property describes a specific physical property such as material or shape.
type Property struct {
	Name      string
	Attribute Attribute
	Verb      string
	Article   string
}

// PropertyRecipe describes a physical property such as material or shape.
type PropertyRecipe struct {
	Name           string   `json:"name"`
	Frequency      float64  `json:"frequency"`
	AttributeNames []string `json:"attribute_names"`
	VerbNames      []string `json:"verb_names"`
	Countable      bool     `json:"countable"`
}

// property creates a Property according the the PropertyRecipe.
func (p *PropertyRecipe) property(atbs map[string]AttributeRecipe, verbs map[string]Verb) (Property, error) {
	n := randString(p.AttributeNames)
	atb, err := atbs[n].attribute(atbs)
	if err != nil {
		return Property{}, nil
	}

	v := randString(p.VerbNames)
	return Property{
		Name:      p.Name,
		Attribute: atb,
		Verb:      verbs[v].RandVariant(),
		Article:   checkArticle(atb.Description, p.Countable),
	}, nil
}

// checkArticle returns the appropriate article.
func checkArticle(word string, countable bool) string {
	if !countable {
		return ""
	}

	return article.Indefinite(word)
}
