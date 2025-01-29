package htmlmodels_test

import (
	"testing"

	htmlmodels "github.com/fengdotdev/golibs-staticpages/models/html_models"
)

func TestAttributes_interface(t *testing.T) {

	var _ htmlmodels.AttributesInterface = htmlmodels.NewAttributesEmpty()
}
