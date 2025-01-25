package agnostic

type Title struct {
	title string
}

func NewTitle(content string) *Title {
	return &Title{title: content}
}

func (t *Title) GetTitle() string {
	return t.title
}
