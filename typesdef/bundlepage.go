package typesdef

type BundlePage struct {
	Title  string
	Route  string
	HTML   string
	JS     string
	CSS    string
	Assets []string
}

func NewBundlePage(title, route, html, js, css string, assets []string) *BundlePage {
	return &BundlePage{
		Title:  title,
		Route:  route,
		HTML:   html,
		JS:     js,
		CSS:    css,
		Assets: assets,
	}
}

func (b *BundlePage) IsJs() bool {
	return b.JS != ""
}

func (b *BundlePage) IsCss() bool {
	return b.CSS != ""
}

func (b *BundlePage) IsHTML() bool {
	return b.HTML != ""
}

func (b *BundlePage) AreAssets() bool {
	return len(b.Assets) > 0
}
