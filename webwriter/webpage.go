package webwriter

import "github.com/fengdotdev/golibs-staticpages/models/agnostic"

type WebPage struct {
	pageType PageType
	pages    []agnostic.Page
}

func (w *WebPage) IsSPA() bool {
	return w.pageType == SinglePageType
}

func (w *WebPage) GetPages() []agnostic.Page {
	return w.pages
}

func (w *WebPage) AddPage(page agnostic.Page) {
	w.pages = append(w.pages, page)
}
