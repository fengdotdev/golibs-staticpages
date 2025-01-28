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

	// TODO TEST
	var err error
	//asuming we are in working dir
	writedFiles := make([]string, 0)

	defer func() {
		if err != nil {
			// clean all files
			log.Println("Error writing files, cleaning up "+outputfolder, err)
			for _, file := range writedFiles {
				err := RemoveAllInWorkingDir(file)
				if err != nil {
					log.Println("Error cleaning up "+file, err)
				}
			}
		}
	}()


	// clean output folder
	err = RemoveAllInWorkingDir(outputfolder)
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
		err := MkdirAllInWorkingDir(route)
		writedFiles = append(writedFiles, route)
		if err != nil {
			outerErr = err
			break
		}

		if bundle.IsHTML() {
			htmlpath := routeFix(route) + ".html"
			err := WriteToFileInWorkingDir(htmlpath, bundle.HTML)
			writedFiles = append(writedFiles, htmlpath)
			if err != nil {
				outerErr = err
				break
			}
		}

		if bundle.IsCss() {
			csspath := routeFix(route) + ".css"
			err := WriteToFileInWorkingDir(csspath, bundle.CSS)
			writedFiles = append(writedFiles, csspath)
			if err != nil {
				outerErr = err
				break
			}
		}

		if bundle.IsJs() {
			jspath := routeFix(route) + ".js"
			err := WriteToFileInWorkingDir(jspath, bundle.JS)
			writedFiles = append(writedFiles, jspath)
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

// routeFix adds index to the end of the route if it ends with a slash
func routeFix(route string) string {
	if strings.HasSuffix(route, "/") {
		return route + "index"
	}
	return route + "/index"
}
