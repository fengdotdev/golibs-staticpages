package webwriter

import (
	"log"
	"strings"
)

type WebWriter struct {
	web  WebPage
	root string
}

func NewWriter(web WebPage, root string) *WebWriter {
	return &WebWriter{
		web:  web,
		root: root,
	}
}

// Write writes the web page to the output folder
func (w *WebWriter) Write(outputFolder string) error {
	if w.web.IsSPA() {
		return w.writeSinglePage(outputFolder)
	}
	return w.writeMultiPage(outputFolder)
}

func (w *WebWriter) writeSinglePage(outputfolder string) error {
	panic("not implemented") // TODO: Implement
}

func (w *WebWriter) writeMultiPage(outputfolder string) error {
	// clean output folder
	err := RemoveAll(outputfolder)
	if err != nil {
		return err
	}

	pages := w.web.GetPages()

	var outerErr error = nil
	for _, page := range pages {
		bundle, innerErr := page.RenderBundle()
		if innerErr != nil {
			outerErr = innerErr
			break
		}

		route := outputfolder + w.root + bundle.Route

		// making folders
		err := MkdirAll(route)
		if err != nil {
			outerErr = err
			break
		}

		if bundle.IsHTML() {

			err := WriteHTML(routeFix(route), bundle.HTML)
			if err != nil {
				outerErr = err
				break
			}
		}

		if bundle.IsCss() {
			err := WriteCSS(routeFix(route), bundle.CSS)
			if err != nil {
				outerErr = err
				break
			}
		}

		if bundle.IsJs() {
			err := WriteJS(routeFix(route), bundle.JS)
			if err != nil {
				outerErr = err
				break
			}
		}

		if bundle.AreAssets() {
			panic(" writeMultiplage bundle assets not implemented") // TODO: Implement
		}

	}
	if outerErr != nil {
		// clear all files
		log.Println("Error writing files, cleaning up "+outputfolder, outerErr)
		return outerErr
	}

	return nil
}

func routeFix(route string) string {
	if strings.HasSuffix(route, "/") {
		return route + "index"
	}
	return route + "/index"
}
