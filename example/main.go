package main

import (
	"github.com/fengdotdev/golibs-staticpages/models/agnostic"
	"github.com/fengdotdev/golibs-staticpages/server"
	"github.com/fengdotdev/golibs-staticpages/webwriter"
)

func main() {

	webpage := webwriter.NewMultiPageEmpty()

	pageHome := agnostic.Page{
		Title:      "Home",
		Route:      "/",
		Body:       "Welcome to the homepage",
		Header:     "header",
		Footer:     "footer",
		Companions: nil,
	}

	webpage.AddPage(pageHome)

	pageAboutUs := agnostic.Page{
		Title:      "About Us",
		Route:      "/about",
		Body:       "About Us",
		Header:     "Header",
		Footer:     "Footer",
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
