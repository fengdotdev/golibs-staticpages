package typesdef_test

import (
	"testing"

	"github.com/fengdotdev/golibs-staticpages/typesdef"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestHTMLValidate(t *testing.T) {

	someCorrectHTML := "<html><head></head><body></body></html>"

	// Test correct HTML
	html := typesdef.NewHTML(someCorrectHTML)
	assert.TrueWithMessage(t, html.IsValid(), "Correct HTML should be valid")

	someIncorrectHTML := "<html><head></head><body></body>"

	// Test incorrect HTML
	html = typesdef.NewHTML(someIncorrectHTML)
	assert.FalseWithMessage(t, html.IsValid(), "Incorrect HTML should be invalid")
}
