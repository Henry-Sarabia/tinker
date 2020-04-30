package tinker

import (
	"html/template"
	"math/rand"
	"strings"

	"github.com/pkg/errors"
)

const (
	fileItemTemplate string = "test_data/item.tmpl"
	fileItems        string = "test_data/items.json"
	fileAdverbs      string = "test_data/adverbs.json"
	fileCreatures    string = "test_data/creatures.json"
	fileDecorations  string = "test_data/decorations.json"
	fileMaterials    string = "test_data/materials.json"
	fileQualities    string = "test_data/qualities.json"
	fileVerbs        string = "test_data/verbs.json"
)

// Generator generates items using the provided recipes.
type Generator struct {
	ItemBank []ItemRecipe
	AtbBank  map[string]AttributeRecipe
	VerbBank map[string]Verb
}

// New returns a properly configured Generator.
func New() (*Generator, error) {
	items, err := readItemRecipes(fileItems)
	if err != nil {
		return nil, err
	}

	atbs, err := loadAttributeRecipes(fileAdverbs, fileCreatures, fileDecorations, fileMaterials, fileQualities)
	if err != nil {
		return nil, err
	}

	verbs, err := loadVerbs(fileVerbs)
	if err != nil {
		return nil, err
	}

	g := &Generator{
		ItemBank: items,
		AtbBank:  atbs,
		VerbBank: verbs,
	}

	return g, nil
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
	t, err := template.ParseFiles(fileItemTemplate)
	if err != nil {
		return Item{}, errors.Wrap(err, "cannot parse template")
	}

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
