package agnostic

import (
	"strings"
	"text/template"

	"github.com/fengdotdev/golibs-staticpages/typesdef"
)


func Companions(companions ...typesdef.Companion) []typesdef.Companion {
	if len(companions) == 0 || companions == nil {
		return nil
	}
	return companions
}

type Page struct {
	Title      Title
	Route      Route
	Header     Header
	Body       Body
	Footer     Footer
	Companions []typesdef.Companion
}

func NewPage(title Title, route Route, header Header, body Body, footer Footer, companions []typesdef.Companion) *Page {
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
func (p *Page) RenderBundle() (*typesdef.BundlePage, error) {

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

	return typesdef.NewBundlePage(p.Title.title, p.Route.route, html, js, css, nil), nil

}

func (p *Page) RenderHTML() (string, error) {

	// get the companion
	companion := pageCompanion{}
	for _, c := range p.Companions {
		companion.AppendCompanion(c.ForHTML(), c.GetPosition())
	}

	return htmlTemplate(p.Title, p.Header, p.Body, p.Footer, companion)
}

func (p *Page) RenderCSS() (string, error) {
	return "foo css", nil
}

func (p *Page) RenderJS() (string, error) {
	return "foo js", nil
}

func htmlTemplate(title Title, header Header, body Body, footer Footer, companion pageCompanion) (string, error) {

	tmpl := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		{{ .HeadCompanion }}
		<title>{{ .Title }}</title>
	</head>
	<body>
		 {{ .Header }}
			{{ .Body }}
		 {{ .Footer }}
		  {{ .BodyCompanion }}
	</body>
	{{ .FooterCompanion }}
	</html>`

	// use text/template to render the html because html/template escapes the html code
	// and we want to render the html code as it is
	intermediary, err := template.New("html").Parse(tmpl)
	if err != nil {
		return "", err
	}
	buffer := new(strings.Builder)

	headerRender := header.GetChild().RenderHTML()
	bodyRender := body.GetChild().RenderHTML()
	footerRender := footer.GetChild().RenderHTML()

	err = intermediary.Execute(buffer, htmlTemplateData{
		Title:           title.GetTitle(),
		Header:          headerRender.GetHTML(),
		Body:            bodyRender.GetHTML(),
		HeadCompanion:   companion.HeadCompanion,
		BodyCompanion:   companion.BodyCompanion,
		FooterCompanion: companion.FooterCompanion,
		Footer:          footerRender.GetHTML(),
	})

	return buffer.String(), err
}

type htmlTemplateData struct {
	Title           string
	Header          string
	HeadCompanion   string
	BodyCompanion   string
	FooterCompanion string
	Body            string
	Footer          string
}

type pageCompanion struct {
	HeadCompanion   string
	BodyCompanion   string
	FooterCompanion string
}

func (p *pageCompanion)AppendCompanion (companion string, position typesdef.CompanionPosition) {
	switch position {
	case typesdef.Head:
		p.AppendHeadCompanion(companion)
	case typesdef.Body:
		p.AppendBodyCompanion(companion)
	case typesdef.Footer:
		p.AppendFooterCompanion(companion)
	}
}

func (p *pageCompanion) AppendHeadCompanion(companion string) {
	p.HeadCompanion += companion + "\n"
}

func (p *pageCompanion) AppendBodyCompanion(companion string) {
	p.BodyCompanion += companion + "\n"
}

func (p *pageCompanion) AppendFooterCompanion(companion string) {
	p.FooterCompanion += companion + "\n"
}
