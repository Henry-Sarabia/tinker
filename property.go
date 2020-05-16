package tinker

// Property describes a specific physical property such as material or shape.
type Property struct {
	Name      string
	Attribute Attribute
	Verb      string
	Countable bool
}

// PropertyRecipe describes a physical property such as material or shape.
type PropertyRecipe struct {
	Name           string   `json:"name"`
	Frequency      float64  `json:"frequency"`
	AttributeNames []string `json:"attribute_names"`
	VerbNames      []string `json:"verb_names"`
	Countable      bool     `json:"countable"`
}

// property creates a Property according to the PropertyRecipe receiver.
func (pr *PropertyRecipe) property(atbs map[string]AttributeRecipe, verbs map[string]Verb) (Property, error) {
	n := randString(pr.AttributeNames)
	atb, err := atbs[n].attribute(atbs)
	if err != nil {
		return Property{}, nil
	}

	v := randString(pr.VerbNames)
	return Property{
		Name:      pr.Name,
		Attribute: atb,
		Verb:      verbs[v].RandSynonym(),
		Countable: pr.Countable,
	}, nil
}
