package agnostic

type Footer struct {
	content string
}

func NewFooter(content string) *Footer {
	return &Footer{content: content}
}

func (f *Footer) GetContent() string {
	return f.content
}
