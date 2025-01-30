package widgets_test

import (
	"testing"

	"github.com/fengdotdev/golibs-staticpages/widgets"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestContainter_widgetInterface(t *testing.T) {
	var container widgets.Widget = widgets.NewEmptyContainer()

	assert.TrueWithMessage(t, container != nil, "container should not be nil")

}
