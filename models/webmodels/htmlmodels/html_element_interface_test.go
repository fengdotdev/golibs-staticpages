package htmlmodels_test

import (
	"testing"

	"github.com/fengdotdev/golibs-staticpages/models/webmodels/htmlmodels"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestHTMLelement_interface(t *testing.T) {

	var attributes htmlmodels.HTMlAttributes = &htmlmodels.HTMLElement{}
	assert.TrueWithMessage(t, attributes != nil, "HTMLelement_interface should not be nil")

	var htmlrenders htmlmodels.HTML_Renders = &htmlmodels.HTMLElement{}

	assert.TrueWithMessage(t, htmlrenders != nil, "HTMLelement_interface should not be nil")

	var htmlcontent htmlmodels.HTML_Content = &htmlmodels.HTMLElement{}

	assert.TrueWithMessage(t, htmlcontent != nil, "HTMLelement_interface should not be nil")

	var htmlidclassstyles htmlmodels.HTML_ID_CLass_Styles = &htmlmodels.HTMLElement{}

	assert.TrueWithMessage(t, htmlidclassstyles != nil, "HTMLelement_interface should not be nil")

	var a htmlmodels.HTMLElementInterface = &htmlmodels.HTMLElement{}
	assert.TrueWithMessage(t, a != nil, "HTMLelement_interface should not be nil")
}
