package tinker

import (
	"html/template"
	"log"
	"math/rand"
	"strings"
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
	Items      []ItemRecipe
	Attributes map[string]AttributeRecipe
	Verbs      map[string]Verb
}

// New returns a properly configured Generator.
func New() *Generator {
	items := readItemRecipes(fileItems)
	atbs := loadAttributeRecipes(fileAdverbs, fileCreatures, fileDecorations, fileMaterials, fileQualities)
	verbs := loadVerbs(fileVerbs)
	g := &Generator{
		Items:      items,
		Attributes: atbs,
		Verbs:      verbs,
	}

	return g
}

// Item generates a random item corresponding to one of the loaded ItemRecipes.
func (g *Generator) Item() Item {
	i := rand.Intn(len(g.Items))
	return g.item(g.Items[i])
}

// item generates an item according to the provided ItemRecipe.
func (g *Generator) item(rcp ItemRecipe) Item {
	comps := g.components(rcp.ComponentRecipes())
	t, err := template.ParseFiles(fileItemTemplate)
	if err != nil {
		log.Fatal(err)
	}

	it := Item{
		Name:       rcp.Name,
		Prelude:    comps[0].RandProperty().Attribute,
		Components: comps,
	}

	wr := &strings.Builder{}
	t.Execute(wr, it)
	it.Description = wr.String()
	return it
}

// components generates a slice of Components according to the provided ComponentRecipes.
func (g *Generator) components(rcps []ComponentRecipe) []Component {
	comps := []Component{}
	for _, r := range rcps {
		comp := g.component(r)
		comps = append(comps, comp)
	}

	return comps
}

// component generates a Component according to the provided ComponentRecipe.
func (g *Generator) component(recipe ComponentRecipe) Component {
	props := []Property{}
	for _, p := range recipe.PropertyRecipes() {
		props = append(props, p.property(g.Attributes, g.Verbs))
	}

	return Component{
		Name:       recipe.Name,
		Properties: props,
	}
}
