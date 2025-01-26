package api

const (
	GOTO ActionKind = "goto"
)

type Goto struct {
	Url string
}

func NewGoto(url string) *Goto {
	return &Goto{
		Url: url,
	}
}

func (g *Goto) GetType() ActionKind {
	return GOTO
}

func (g *Goto) IsValid() bool {
	return true
}
