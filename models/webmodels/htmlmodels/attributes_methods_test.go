package htmlmodels_test

import (
	"strings"
	"testing"

	"github.com/fengdotdev/golibs-staticpages/models/webmodels/htmlmodels"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestAttributes_ToString(t *testing.T) {

	var attributes htmlmodels.Attributes = htmlmodels.NewAttributesEmpty()

	tailwindcss := "text-red-500"

	id := "foo"
	attributes.AddAttribute("id", id)
	attributes.AddAttribute("class", tailwindcss)

	result := attributes.ToString()

	expected := `id="foo" class="text-red-500"`
	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(result))

}
