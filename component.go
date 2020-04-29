package tinker

import (
	"log"
	"math/rand"
	"strings"
	"text/template"
)

const templateFile string = "test_data/item.tmpl"

// Component holds specific data generated from a Recipe.
type Component struct {
	Name       string
	Properties []Property //[material]slightly burned ashwood; might also need to be a slice for ordering
}

// RandProperty returns a randomly selected RandProperty from the Component's Properties.
func (c Component) RandProperty() Property {
	return c.Properties[rand.Intn(len(c.Properties))]
}

// FirstProperty returns the Component's first Property.
func (c Component) FirstProperty() Property {
	return c.Properties[0]
}

// IsMultiProperty returns true if the Component has multiple Properties.
func (c Component) IsMultiProperty() bool {
	if len(c.Properties) > 1 {
		return true
	}

	return false
}

// compToSentence generates a sentence for the provided Component.
func compToSentence(comp Component) string {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal(err)
	}

	var b strings.Builder
	err = tmpl.Execute(&b, comp)
	if err != nil {
		log.Fatal(err)
	}

	return b.String()

}
