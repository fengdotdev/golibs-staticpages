package agnostic

type Body struct {
	content string
}

func NewBody(content string) *Body {
	return &Body{
		content: content,
	}
}

func (b *Body) GetContent() string {
	return b.content
}
