package tinker

// Property describes a specific physical property.
type Property struct {
	Name      string
	Attribute Attribute
}

// dont try to return strings, try to return a well-composed struct

// PropertyRecipe describes a physical property.
type PropertyRecipe struct {
	Name           string   `json:"name"`
	Frequency      float64  `json:"frequency"`
	AttributeNames []string `json:"attribute_names"`
}

func (p *PropertyRecipe) roll(bank map[string]AttributeRecipe) string {
	n := randomString(p.AttributeNames)
	return bank[n].resolve(bank)
}

func (p *PropertyRecipe) property(bank map[string]AttributeRecipe) Property {
	n := randomString(p.AttributeNames)
	return Property{
		Name:      p.Name,
		Attribute: bank[n].attribute(bank),
	}
}
