package tinker

// Item contains the description of a generated Item.
type Item struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Prelude     Attribute   `json:"-"`
	Components  []Component `json:"-"`
}
