package webwriter_test

import (
	"testing"

	"github.com/fengdotdev/golibs-staticpages/typesdef"
	"github.com/fengdotdev/golibs-staticpages/webwriter"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestTailwindCompanionInterface(t *testing.T) {

	var companion typesdef.Companion = webwriter.NewTailwindCompanion()

	assert.TrueWithMessage(t, companion != nil, "companion should not be nil")

}
