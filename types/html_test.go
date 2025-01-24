package types_test

import (
	"testing"

	"github.com/fengdotdev/golibs-staticpages/types"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestHTMLValidate(t *testing.T) {

	someCorrectHTML := "<html><head></head><body></body></html>"

	// Test correct HTML
	html := types.NewHTML(someCorrectHTML)
	assert.TrueWithMessage(t, html.IsValid(), "Correct HTML should be valid")

	someIncorrectHTML := "<html><head></head><body></body>"

	// Test incorrect HTML
	html = types.NewHTML(someIncorrectHTML)
	assert.FalseWithMessage(t, html.IsValid(), "Incorrect HTML should be invalid")
}
