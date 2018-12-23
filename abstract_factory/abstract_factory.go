package abstractFactory

// Director directe to builder.
// Director has Builder instance.
type Factory interface {
	createLink(caption string, url string)
	createTray(caption string)
	createPage(title string, author string)
}

// NewDirector constructor.
func NewDirector(b Builder) *Director {
	return &Director{b}
}
