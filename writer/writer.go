package writer

const (
	JSExtension   = ".js"
	CSSExtension  = ".css"
	HTMLExtension = ".html"
)

type Writer struct {
}

func NewWriter() *Writer {
	return &Writer{}
}

func (w *Writer) Write() error {
	panic("not implemented") // TODO: Implement
}

func writeJSToFile() error {
	panic("not implemented") // TODO: Implement
}

func writeCSSToFile() error {
	panic("not implemented") // TODO: Implement
}

func writeHTMLToFile() error {
	panic("not implemented") // TODO: Implement
}
