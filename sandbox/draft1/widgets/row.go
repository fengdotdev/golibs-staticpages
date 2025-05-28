package widgets

import (
	"github.com/fengdotdev/golibs-staticpages/typesdef"
)

type Row struct {
	children     []Widget
	rowOptions   RowOptions
	classBaseOpt typesdef.ClassBaseOpt
}

func NewRow(children []Widget, rowOptions RowOptions, classBaseOpt typesdef.ClassBaseOpt) *Row {
	return &Row{
		children:     children,
		rowOptions:   rowOptions,
		classBaseOpt: classBaseOpt,
	}
}

type RowOptions struct {

}
