package typesdef

import "github.com/fengdotdev/golibs-staticpages/validators"

type JS struct {
	content string
}

func NewJS(content string) *JS {
	return &JS{
		content: content,
	}
}

func (j *JS) GetType() CompanionType {
	return JSCompanion
}

func (j *JS) GetContent() string {
	return j.content
}

func (j *JS) SetContent(content string) {
	j.content = content
}

func (j *JS) IsValid() bool {
	return validators.JSIsValid(j.content)
}
