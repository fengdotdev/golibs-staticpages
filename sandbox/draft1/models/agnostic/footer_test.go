package agnostic_test

import (
	"testing"

	"github.com/fengdotdev/golibs-staticpages/models/agnostic"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestFooter(t *testing.T) {

	var footer agnostic.Agnostic = agnostic.NewFooter(nil)
	//TODO Update this test to not nil test
	assert.TrueWithMessage(t, footer != nil, "Footer should not be nil")
}
