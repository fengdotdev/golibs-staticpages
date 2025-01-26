package agnostic_test

import (
	"testing"

	"github.com/fengdotdev/golibs-staticpages/models/agnostic"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestHeader(t *testing.T) {

	var header agnostic.Agnostic = agnostic.NewHeader(nil)
	//TODO Update this test to not nil test
	assert.TrueWithMessage(t, header != nil, "Header should not be nil")
}
