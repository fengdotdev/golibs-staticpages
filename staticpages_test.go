package golibsstaticpages_test

import (
	"testing"

	"github.com/fengdotdev/golibs-staticpages/models/agnostic"
	"github.com/fengdotdev/golibs-staticpages/webwriter"
	"github.com/fengdotdev/golibs-testing/assert"
)

const (
	TestOutputDir = "TestOutput"
)

func TestE2E(t *testing.T) {

	testName := "TestE2E"
	outputDir := TestOutputDir + testName

	home := agnostic.Page{
		Title:      "Home",
		Route:      "/",
		Header:     "Home",
		Body:       "Welcome to the home page",
		Footer:     "Home",
		Companions: nil,
	}

	about := agnostic.Page{
		Title:      "About",
		Route:      "/about",
		Header:     "About",
		Body:       "Welcome to the about page",
		Footer:     "About",
		Companions: nil,
	}

	contact := agnostic.Page{
		Title:      "Contact",
		Route:      "/contact",
		Header:     "Contact",
		Body:       "Welcome to the contact page",
		Footer:     "Contact",
		Companions: nil,
	}

	// Create a new multi-page

	webPage := webwriter.NewMultiPage([]agnostic.Page{home, about, contact})

	wr := webwriter.NewWriter(*webPage, "")
	err := wr.Write(outputDir)
	assert.Nil(t, err)
}
