package builder

import "os"

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
	makeItems(items []string)
	close()
}

// TextBuilder implements Builder.
type TextBuilder struct {
	str string
}

// NewTextBuilder TextBuilder constructor
func NewTextBuilder() *TextBuilder {
	return &TextBuilder{}
}

func (t *TextBuilder) makeTitle(s string) {
	t.str = "=============================¥n"
	t.str = t.str + "「" + s + "」"
}

func (t *TextBuilder) makeString(s string) {
	t.str = "□" + s + "¥n"
	t.str = t.str + "¥n"
}

func (t *TextBuilder) makeItems(items []string) {
	for i := 0; i < len(items); i++ {
		t.str = t.str + items[i] + "¥n"
	}
	t.str = t.str + "¥n"
}

func (t *TextBuilder) close() {
	t.str = "=============================¥n"
}

// HTMLBuilder implement Builder.
type HTMLBuilder struct {
}

// NewHTMLBuilder HTMLBuilder constructor.
func NewHTMLBuilder() *HTMLBuilder {
	return &HTMLBuilder{}
}

func (h HTMLBuilder) makeTitle(s string) {
	fileNm := s + ".html"
	file, err := os.Create(fileNm)
	if err != nil {
		return
	}
	defer file.Close()
	output := "<html><head><title>" + s + "</head></title></html>"
	file.Write(([]byte)(output))
}

func (h HTMLBuilder) makeString(s string) {
}

func (h HTMLBuilder) makeItems(s []string) {
}

func (h HTMLBuilder) close() {
}
