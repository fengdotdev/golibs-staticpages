package widgets

type Row struct {
	children []Widget
	rowOptions options.rowOptions
}



func NewRow(children []Widget) *Row {
	return &Row{
		children: children,
	}
}