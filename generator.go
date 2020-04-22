package tinker

import (
	"html/template"
	"log"
	"math/rand"
	"strings"
)

const (
	fileAdverbs      string = "test_data/adverbs.json"
	fileGroups       string = "test_data/attribute_groups.json"
	fileDecorations  string = "test_data/decorations.json"
	fileItemTemplate string = "test_data/item.tmpl"
	fileMaterials    string = "test_data/materials.json"
	fileQualities    string = "test_data/qualities.json"
	fileRecipes      string = "test_data/recipes.json"
)

// Generator builds items.
type Generator struct {
	Items      []ItemRecipe
	Attributes map[string]AttributeRecipe
}

// New returns a properly configured Generator.
func New() *Generator {

	items := readItemRecipes(fileRecipes)
	atbs := loadAttributeRecipes(fileAdverbs, fileDecorations, fileMaterials, fileQualities)
	g := &Generator{
		Items:      items,
		Attributes: atbs,
	}

	return g
}

// recipe returns a random recipe from the Generator's Items.
func (g *Generator) recipe() ItemRecipe {
	return g.Items[rand.Intn(len(g.Items))]
}

// components builds a slice of Components according to the provided ComponentRecipes.
func (g *Generator) components(recipes []ComponentRecipe) []Component {
	comps := []Component{}
	for _, r := range recipes {
		comp := g.component(r)
		comps = append(comps, comp)
	}

	return comps
}

// component builds a Component according to the provided ComponentRecipe.
func (g *Generator) component(recipe ComponentRecipe) Component {
	props := []Property{}
	for _, p := range recipe.PropertyRecipes() {
		props = append(props, p.property(g.Attributes))
	}

	return Component{
		Name:       recipe.Name,
		Properties: props,
	}
}

// item generates an item according to the provided ItemRecipe.
func (g *Generator) item(recipe ItemRecipe) Item {
	comps := g.components(recipe.ComponentRecipes())
	t, err := template.ParseFiles(fileItemTemplate)
	if err != nil {
		log.Fatal(err)
	}

	s := &strings.Builder{}
	i := Item{
		Name:       recipe.Name,
		Descriptor: comps[rand.Intn(len(comps))].Property().Attribute,
		Components: comps,
	}

	t.Execute(s, i)
	i.Text = s.String()
	return i
}

// Item generates a random item corresponding to one of the loaded ItemRecipes.
func (g *Generator) Item() Item {
	return g.item(g.recipe())
}
