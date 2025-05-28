package server

import (
	"fmt"
	"net/http"
)

type DevServer struct {
	servefolder string
	port        int
}

func NewDevServer(port int, serveFolder string) *DevServer {
	return &DevServer{
		servefolder: serveFolder,
		port:        port,
	}
}

func (d *DevServer) Start() {

	d.serveFolder()
	d.displayServerRunningAt()

	http.ListenAndServe(d.getPort(), nil)
}

func (d *DevServer) Stop() {
	// TODO implement
}

func (d *DevServer) getPort() string {
	return ":" + fmt.Sprint(d.port)
}

func (d *DevServer) displayServerRunningAt() {
	fmt.Println("Server running at http://localhost" + d.getPort())
}

func (d *DevServer) noContent() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
}

func (d *DevServer) serveFolder() {
	http.Handle("/", http.FileServer(http.Dir(d.servefolder)))
}
