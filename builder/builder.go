package builder

// Director directe to builder.
type Director struct {
	build Builder
}

// NewDirector constructor.
func NewDirector(b Builder) *Director {
	return &Director{b}
}

// Construct instruct the builder instance.
func (d *Director) Construct() {
	d.build.makeTitle("Greeting")
	d.build.makeString("朝から昼にかけて")
	d.build.makeItems([]string{"おはようございます", "こんにちは"})
	d.build.makeString("夜に")
	d.build.makeItems([]string{"こんばんは", "おやすみなさい", "さようなら"})
}

// Builder interface define some function.
type Builder interface {
	makeTitle(s string)
	makeString(s string)
	makeItems(s []string)
	close()
}

// TextBuilder implements Builder.
type TextBuilder struct {
}

// NewTextBuilder TextBuilder constructor
func NewTextBuilder() *TextBuilder {
	return &TextBuilder{}
}

func (t TextBuilder) makeTitle(s string) {
}

func (t TextBuilder) makeString(s string) {
}

func (t TextBuilder) makeItems(s []string) {
}

func (t TextBuilder) close() {
}

// HTMLBuilder implement Builder.
type HTMLBuilder struct {
}

// NewHTMLBuilder HTMLBuilder constructor.
func NewHTMLBuilder() *HTMLBuilder {
	return &HTMLBuilder{}
}

func (h HTMLBuilder) makeTitle(s string) {
}

func (h HTMLBuilder) makeString(s string) {
}

func (h HTMLBuilder) makeItems(s []string) {
}

func (h HTMLBuilder) close() {
}
