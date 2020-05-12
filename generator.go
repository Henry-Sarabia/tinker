package tinker

import (
	"html/template"
	"math/rand"
	"strings"
)

// Generator generates items using the provided recipes.
type Generator struct {
	ItemBank []ItemRecipe
	AtbBank  map[string]AttributeRecipe
	VerbBank map[string]Verb
}

// New returns an unloaded Generator.
func New() *Generator {
	return &Generator{}
}

// LoadItemRecipes loads the provided files into the Generator.
func (g *Generator) LoadItemRecipes(filenames ...string) error {
	rcps, err := readItemRecipes(filenames...)
	if err != nil {
		return err
	}
	g.ItemBank = rcps
	return nil
}

// LoadAttributeRecipes loads the provided files into the Generator.
func (g *Generator) LoadAttributeRecipes(filenames ...string) error {
	rcps, err := loadAttributeRecipes(filenames...)
	if err != nil {
		return err
	}
	g.AtbBank = rcps
	return nil
}

// LoadVerbs loads the provided files into the Generator.
func (g *Generator) LoadVerbs(filenames ...string) error {
	verbs, err := loadVerbs(filenames...)
	if err != nil {
		return err
	}
	g.VerbBank = verbs
	return nil
}

// Item generates a random item corresponding to one of the loaded ItemRecipes.
func (g *Generator) Item() (Item, error) {
	i := rand.Intn(len(g.ItemBank))
	it, err := g.item(g.ItemBank[i])
	if err != nil {
		return Item{}, err
	}

	return it, err
}

// item generates an item according to the provided ItemRecipe.
func (g *Generator) item(rcp ItemRecipe) (Item, error) {
	comps, err := g.components(rcp.ComponentRecipes())
	if err != nil {
		return Item{}, err
	}
	t, err := template.New("item").Parse(itemTemplate)

	it := Item{
		Name:       rcp.Name,
		Prelude:    comps[0].RandProperty().Attribute,
		Components: comps,
	}

	wr := &strings.Builder{}
	t.Execute(wr, it)
	it.Description = wr.String()
	return it, nil
}

// components generates a slice of Components according to the provided ComponentRecipes.
func (g *Generator) components(rcps []ComponentRecipe) ([]Component, error) {
	comps := []Component{}
	for _, r := range rcps {
		comp, err := g.component(r)
		if err != nil {
			return nil, err
		}
		comps = append(comps, comp)
	}

	return comps, nil
}

// component generates a Component according to the provided ComponentRecipe.
func (g *Generator) component(rcp ComponentRecipe) (Component, error) {
	props := []Property{}
	for _, p := range rcp.PropertyRecipes() {
		prop, err := p.property(g.AtbBank, g.VerbBank)
		if err != nil {
			return Component{}, err
		}

		props = append(props, prop)
	}

	return Component{
		Name:       rcp.Name,
		Properties: props,
	}, nil
}
