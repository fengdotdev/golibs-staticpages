package webwriter_test

import (
	"testing"

	"github.com/fengdotdev/golibs-staticpages/webwriter"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestWriteToFile(t *testing.T) {

	err := webwriter.WriteToFile("foo", "content")
	t.Log(err)
	assert.NoError(t, err)
}

func TestMkdirAll(t *testing.T) {
	err := webwriter.MkdirAll("test")
	t.Log(err)
	assert.NoError(t, err)
}
