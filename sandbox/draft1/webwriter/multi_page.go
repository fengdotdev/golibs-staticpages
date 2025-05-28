package webwriter

import "github.com/fengdotdev/golibs-staticpages/models/agnostic"

func NewMultiPage(pages []agnostic.Page) *WebPage {
	return &WebPage{
		pageType: MultiPageType,
		pages:    pages,
	}
}

func NewMultiPageEmpty() *WebPage {
	return &WebPage{
		pageType: MultiPageType,
		pages:    []agnostic.Page{},
	}
}


