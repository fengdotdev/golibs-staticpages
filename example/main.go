package main

import (
	"github.com/fengdotdev/golibs-staticpages/models/agnostic"
	"github.com/fengdotdev/golibs-staticpages/server"
	"github.com/fengdotdev/golibs-staticpages/webwriter"
	"github.com/fengdotdev/golibs-staticpages/widgets"
)

func main() {

	webpage := webwriter.NewMultiPageEmpty()

	text := widgets.NewTextWithClass("Welcome to the home page", "text-lg font-semibold text-gray-900")

	pageHome := agnostic.Page{
		Title:      *agnostic.NewTitle("Home"),
		Route:      *agnostic.NewRoute("/"),
		Header:     *agnostic.NewHeader(text),
		Body:       *agnostic.NewBody(text),
		Footer:     *agnostic.NewFooter(text),
		Companions: nil,
	}

	webpage.AddPage(pageHome)

	pageAboutUs := agnostic.Page{
		Title:      *agnostic.NewTitle("About Us"),
		Route:      *agnostic.NewRoute("/AboutUs"),
		Header:     *agnostic.NewHeader(text),
		Body:       *agnostic.NewBody(text),
		Footer:     *agnostic.NewFooter(text),
		Companions: nil,
	}

	webpage.AddPage(pageAboutUs)

	webwr := webwriter.NewWriter(*webpage, "")

	output := "output"

	err := webwr.Write(output)
	if err != nil {
		panic(err)
	}

	// Render the page
	// Start the server
	server.NewDevServer(8080, output).Start()
}
