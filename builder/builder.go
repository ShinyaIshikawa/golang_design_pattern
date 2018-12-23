package builder

import (
	"log"
	"os"
)

// Director directe to builder.
// Director has Builder instance.
type Director struct {
	build Builder
}

// NewDirector constructor.
func NewDirector(b Builder) *Director {
	return &Director{b}
}

// Construct instruct the builder instance.
// Construct build sentence.
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
//
type TextBuilder struct {
	str string
}

// NewTextBuilder TextBuilder constructor
func NewTextBuilder() *TextBuilder {
	return &TextBuilder{}
}

func (t *TextBuilder) makeTitle(s string) {
	t.str = t.str + "=============================\n"
	t.str = t.str + "「" + s + "」"
}

func (t *TextBuilder) makeString(s string) {
	t.str = t.str + "□" + s + "\n"
	t.str = t.str + "\n"
}

func (t *TextBuilder) makeItems(items []string) {
	for i := 0; i < len(items); i++ {
		t.str = t.str + items[i] + "\n"
	}
	t.str = t.str + "\n"
}

func (t *TextBuilder) close() {
	t.str = t.str + "=============================\n"
}

// GetResult return decorate strings.
func (t *TextBuilder) GetResult() string {
	return t.str
}

// HTMLBuilder implement Builder.
type HTMLBuilder struct {
	fileNm string
	file   *os.File
}

// NewHTMLBuilder HTMLBuilder constructor.
func NewHTMLBuilder() *HTMLBuilder {
	return &HTMLBuilder{}
}

func (h *HTMLBuilder) makeTitle(s string) {
	h.fileNm = s + ".html"
	file, err := os.Create(h.fileNm)
	if err != nil {
		log.Fatal(err)
		return
	}
	output := "<html><head><title>" + s + "</head></title></html>"
	_, wErr := file.Write(([]byte)(output))
	if wErr != nil {
		log.Fatal(wErr)
	}
	h.file = file
}

func (h *HTMLBuilder) makeString(s string) {
	str := "<p>" + s + "</p>"
	_, err := h.file.Write([]byte(str))
	if err != nil {
		log.Fatal(err)
	}
}

func (h *HTMLBuilder) makeItems(items []string) {
	_, err := h.file.Write([]byte("<ul>"))
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(items); i++ {
		_, err := h.file.Write([]byte(items[i]))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (h *HTMLBuilder) close() {
	_, err := h.file.Write([]byte("</body></html>"))
	if err != nil {
		log.Fatal(err)
	}
	h.file.Close()
}

// GetResult return HTML File name.
func (h *HTMLBuilder) GetResult() string {
	return h.fileNm
}
