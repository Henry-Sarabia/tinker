package tinker

// Item contains the description of a generated Item.
type Item struct {
	Name       string `json:"item"`
	Descriptor Attribute
	Components []Component
	Text       string
}

// IsMultiComponent returns true if the Item has multiple components.
func (i Item) IsMultiComponent() bool {
	if len(i.Components) > 1 {
		return true
	}

	return false
}

// FirstComponent returns the item's first Component.
func (i Item) FirstComponent() Component {
	return i.Components[0]
}
