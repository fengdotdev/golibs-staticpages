package agnostic

import "github.com/fengdotdev/golibs-staticpages/types"

type Page struct {
	Title      string
	Header     string
	Body       string
	Footer     string
	Companions []types.Companion
}

func NewPage(title, header, body, footer string) *Page {
	return &Page{
		Title:  title,
		Header: header,
		Body:   body,
		Footer: footer,
	}
}
