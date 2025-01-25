package helperfuncs_test

import (
	"testing"

	"github.com/fengdotdev/golibs-staticpages/helperfuncs"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestGenerateIdDeterministic(t *testing.T) {
	id1 := helperfuncs.NewSomeUUID()
	id2 := helperfuncs.NewSomeUUID()

	assert.NotEqual(t, id1, id2)

	name1 := "test1"
	name2 := "test2"

	determ1 := helperfuncs.GenerateIdDeterministic(name1, id1)
	determ2 := helperfuncs.GenerateIdDeterministic(name1, id1)

	assert.Equal(t, determ1, determ2)

	determ3 := helperfuncs.GenerateIdDeterministic(name2, id1)
	determ4 := helperfuncs.GenerateIdDeterministic(name2, id2)

	assert.NotEqual(t, determ1, determ3)

	assert.NotEqual(t, determ3, determ4)



}
