package htmlmodels_test

import (
	"testing"

	"github.com/fengdotdev/golibs-staticpages/models/webmodels/htmlmodels"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestAttributes_interface(t *testing.T) {

	var a htmlmodels.AttributesInterface = &htmlmodels.Attributes{}

	assert.TrueWithMessage(t, a != nil, "AttributesInterface should not be nil")
}
