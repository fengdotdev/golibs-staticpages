package webwriter_test

import (
	"testing"

	"github.com/fengdotdev/golibs-staticpages/webwriter"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestWriteToFile(t *testing.T) {

	err := webwriter.WriteToFileInWorkingDir("foo", "content")
	t.Log(err)
	assert.NoError(t, err)
}

func TestMkdirAll(t *testing.T) {
	err := webwriter.MkdirAllInWorkingDir("test")
	t.Log(err)
	assert.NoError(t, err)
}
