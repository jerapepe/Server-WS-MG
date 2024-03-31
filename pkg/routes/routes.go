package routes

import (
	"Server-WS/pkg/websocket"
	"html/template"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/gorilla/mux"
)

const (
	endpointChat = "/chat"
	endpointRoom = "/room"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("../templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func SetRoutes(router *mux.Router) {
	r := websocket.NewRoom()
	staticDir := "../static/files"
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	router.Handle(endpointChat, &templateHandler{filename: "chat.html"}).Methods(http.MethodGet, http.MethodPost, http.MethodOptions)
	router.Handle(endpointRoom, r).Methods(http.MethodGet, http.MethodPost, http.MethodOptions)
	go r.Run()
}
