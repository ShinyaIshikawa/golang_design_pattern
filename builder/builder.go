package builder

type Builder interface {
	makeTitle()
}

// Define clone function.
type TextBuilder struct {
}

func NewTextBuilder() *TextBuilder {
	return &TextBuilder{}
}

// Define clone function.
type HTMLBuilder struct {
}

func NewHTMLBuilder() *HTMLBuilder {
	return &HTMLBuilder{}
}
