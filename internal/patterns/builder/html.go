package builder

import (
	"fmt"
	"strings"
)

// IHtmlBuilder is the builder interface
type IHtmlBuilder interface {
	SetTitle(title string) IHtmlBuilder
	AddParagraph(text string) IHtmlBuilder
	AddList(items []string) IHtmlBuilder
	Build() string
}

// HtmlBuilder is a concrete builder
type HtmlBuilder struct {
	title      string
	paragraphs []string
	lists      [][]string
}

// NewHtmlBuilder creates a new instance of HtmlBuilder
func NewHtmlBuilder() *HtmlBuilder {
	return &HtmlBuilder{}
}

func (hb *HtmlBuilder) SetTitle(title string) *HtmlBuilder {
	hb.title = title
	return hb
}

func (hb *HtmlBuilder) AddParagraph(text string) *HtmlBuilder {
	hb.paragraphs = append(hb.paragraphs, text)
	return hb
}

func (hb *HtmlBuilder) AddList(items []string) *HtmlBuilder {
	hb.lists = append(hb.lists, items)
	return hb
}

func (hb *HtmlBuilder) Build() string {
	var sb strings.Builder
	sb.WriteString("<html>\n")
	sb.WriteString("<head>\n")
	sb.WriteString(fmt.Sprintf("<title>%s</title>\n", hb.title))
	sb.WriteString("</head>\n")
	sb.WriteString("<body>\n")
	for _, p := range hb.paragraphs {
		sb.WriteString(fmt.Sprintf("<p>%s</p>\n", p))
	}
	for _, list := range hb.lists {
		sb.WriteString("<ul>\n")
		for _, item := range list {
			sb.WriteString(fmt.Sprintf("<li>%s</li>\n", item))
		}
		sb.WriteString("</ul>\n")
	}
	sb.WriteString("</body>\n")
	sb.WriteString("</html>\n")
	return sb.String()
}
