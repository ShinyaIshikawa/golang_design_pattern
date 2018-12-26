package abfactory

type factoryStrategy interface {
	createLink(caption string, url string) Link
	createTray(caption string) Tray
	createPage(title string, author string) Page
}

type Item interface {
	makeHtml() string
}

type Link interface {
	Item
}

type ListLink struct {
	url     string
	caption string
}

func NewListLink(url string, caption string) *ListLink {
	return &ListLink{url: url, caption: caption}
}

func (l ListLink) makeHtml() string {
	return "<li><a href=\"" + l.url + "\">" + l.caption + "</a></li>"
}

type Tray interface {
}

type ListTray struct {
	title  string
	author string
}

func NewListTray(title string, author string) *ListTray {
	return &ListTray{title: title, author: author}
}

func (l ListTray) makeHtml() string {
	var buf string
	buf = "<html><head><title>" + l.title + "</title></head>\n"
	buf += "<body>\n"
	buf += "<h1>" + l.title + "</h1>\n"
	buf += "<ul>\n"
}

type Page interface {
	Item
}
