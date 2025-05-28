package widgets

import "github.com/fengdotdev/golibs-staticpages/typesdef"

type Column struct {
	children      []Widget
	columnOptions ColumnOptions
	classBaseOpt  typesdef.ClassBaseOpt
}

func NewColumn(children []Widget, columnOptions ColumnOptions, classBaseOpt typesdef.ClassBaseOpt) *Column {
	return &Column{
		children:      children,
		columnOptions: columnOptions,
		classBaseOpt:  classBaseOpt,
	}
}

type ColumnOptions struct {
}
