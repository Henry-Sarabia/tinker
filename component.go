package tinker

import (
	"log"
	"strings"
	"text/template"
)

const templateFile string = "test_data/item.tmpl"

// Component holds specific data generated from a Recipe.
type Component struct {
	Name       string
	Properties []Property //[material]slightly burned ashwood; might also need to be a slice for ordering
	// Property Property
}

// Property returns the first of a Component's Properties.
func (c *Component) Property() Property {
	return c.Properties[0]
}

// // Description generates the Component's description.
// func (c *Component) Description() string {
// 	// The {blade/name} is {silver/property} and {foo/property} and...

// }

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
