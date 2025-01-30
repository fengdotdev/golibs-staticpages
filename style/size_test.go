package style_test

import (
	"testing"

	"github.com/fengdotdev/golibs-staticpages/interfaces"
	"github.com/fengdotdev/golibs-staticpages/style"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestSize_css(t *testing.T) {
	var s interfaces.RenderCSS = &style.Size{}

	assert.True(t, s != nil)
}
