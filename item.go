package tinker

// Item contains the description of a generated Item.
type Item struct {
	Name       string `json:"item"`
	Descriptor Attribute
	Components []Component
	Text       string
}
