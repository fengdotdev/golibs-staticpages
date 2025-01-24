package webwriter

import "github.com/fengdotdev/golibs-staticpages/models/agnostic"

func NewSinglePage(pages []agnostic.Page) *WebPage {
	panic("not implemented")
	return &WebPage{
		pageType: SinglePageType,
		pages:    pages,
	}
}
