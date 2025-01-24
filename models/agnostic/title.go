package agnostic

type Title struct {
	content string
}

func NewTitle(content string) *Title {
	return &Title{content: content}
}
