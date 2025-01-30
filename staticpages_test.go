package golibsstaticpages_test

import (
	"testing"

	"github.com/fengdotdev/golibs-staticpages/models/agnostic"
	"github.com/fengdotdev/golibs-staticpages/webwriter"
	"github.com/fengdotdev/golibs-staticpages/widgets"
	"github.com/fengdotdev/golibs-testing/assert"
)

const (
	TestOutputDir = "TestOutput"
)

func TestE2E(t *testing.T) {

	testName := "TestE2E"
	outputDir := TestOutputDir + testName

	text := widgets.NewTextWithClass("Welcome to the home page", "text-lg font-semibold text-gray-900")

	container := widgets.NewContainer(text, widgets.ContainerOptions{Id: "container"}, "bg-gray-100")

	companion := webwriter.NewTailwindCompanion()
	home := agnostic.Page{
		Title:      *agnostic.NewTitle("Home"),
		Route:      *agnostic.NewRoute("/"),
		Header:     *agnostic.NewHeader(text),
		Body:       *agnostic.NewBody(container),
		Footer:     *agnostic.NewFooter(text),
		Companions: agnostic.Companions(companion),
	}

	// Create a new multi-page

	webPage := webwriter.NewMultiPage([]agnostic.Page{home})

	wr := webwriter.NewWriter(*webPage, "")
	undo, err := wr.Write(outputDir)
	assert.Nil(t, err)

	// undo the write
	undo()
}
