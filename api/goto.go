package api

const (
	GOTO ActionType = "goto"
)

type Goto struct {
	Url string
}

func NewGoto(url string) *Goto {
	return &Goto{
		Url: url,
	}
}

func (g *Goto) GetType() ActionType {
	return GOTO
}






func (g *Goto) IsValid() bool {
	return true
}
