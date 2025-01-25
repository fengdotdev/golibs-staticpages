package agnostic

type Header struct {
	content string
}

func NewHeader(content string) *Header {
	return &Header{content: content}
}

func (h *Header) GetContent() string {
	return h.content
}
