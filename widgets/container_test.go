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

func TestContainerIdentifier(t *testing.T) {

	container := "container"

	id1 := widgets.ContainerIdentifier(container)
	id2 := widgets.ContainerIdentifier(container)

	assert.Equal(t, id1, id2)

	container2 := "container"

	id3 := widgets.ContainerIdentifier(container2)
	id4 := widgets.ContainerIdentifier(container)

	assert.Equal(t, id3, id4)

	container3 := "supercontainer"

	id5 := widgets.ContainerIdentifier(container3)
	id6 := widgets.ContainerIdentifier(container)

	assert.NotEqual(t, id5, id6)

}
