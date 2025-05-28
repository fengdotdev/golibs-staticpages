package server

type ServerConfig struct {
	Port            int
	AssetsPath      string
	StaticPagesPath string
	StaticPages     map[string]string
}
