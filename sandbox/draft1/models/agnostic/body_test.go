package agnostic_test

import (
	"testing"

	"github.com/fengdotdev/golibs-staticpages/models/agnostic"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestBody(t *testing.T) {

	var body agnostic.Agnostic = agnostic.NewBody(nil)
	//TODO Update this test to not nil test
	assert.TrueWithMessage(t, body != nil, "Body should not be nil")
}
