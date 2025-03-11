package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(webServerPort string) *WebServer {
	ws := &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]map[string]http.HandlerFunc),
		WebServerPort: webServerPort,
	}
	return ws
}

func (s *WebServer) addPingRoute() {
	s.Router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
}

func (s *WebServer) AddRoute(method, path string, handler http.HandlerFunc) {
	if s.Handlers[method] == nil {
		s.Handlers[method] = make(map[string]http.HandlerFunc)
	}
	s.Handlers[method][path] = handler
}

func (s *WebServer) Start() error {
	s.Router.Use(middleware.Logger)
	s.addPingRoute()
	for method, routes := range s.Handlers {
		for path, handler := range routes {
			s.Router.MethodFunc(method, path, handler)
		}
	}

	return http.ListenAndServe(s.WebServerPort, s.Router)
}
