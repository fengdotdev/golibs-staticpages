package agnostic

import (
	"github.com/fengdotdev/golibs-staticpages/types"
)

type Page struct {
	Title      string
	Route      string
	Header     string
	Body       string
	Footer     string
	Companions []types.Companion
}

func NewPage(title, route, header, body, footer string, companions []types.Companion) *Page {
	return &Page{
		Title:      title,
		Route:      route,
		Header:     header,
		Body:       body,
		Footer:     footer,
		Companions: companions,
	}
}

// make a bundle of the page with html, js, css strings separated
func (p *Page) RenderBundle() (*types.BundlePage, error) {

	html, err := p.RenderHTML()
	if err != nil {
		return nil, err
	}

	css, err := p.RenderCSS()
	if err != nil {
		return nil, err
	}

	js, err := p.RenderJS()
	if err != nil {
		return nil, err
	}

	return types.NewBundlePage(p.Title, p.Route, html, js, css, nil), nil

}

func (p *Page) RenderHTML() (string, error) {
	return "foo html", nil
}

func (p *Page) RenderCSS() (string, error) {
	return "foo css", nil
}

func (p *Page) RenderJS() (string, error) {
	return "foo js", nil
}
