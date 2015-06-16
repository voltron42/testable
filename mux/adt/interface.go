package imux

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
)

type Factory interface {
	Vars(r *http.Request) map[string]string
	CurrentRoute(r *http.Request) Route
	NewRouter() Router
}
type Route interface {
	BuildOnly() Route
	BuildVarsFunc(f mux.BuildVarsFunc) Route
	GetError() error
	GetHandler() http.Handler
	GetName() string
	Handler(handler http.Handler) Route
	HandlerFunc(f func(http.ResponseWriter, *http.Request)) Route
	Headers(pairs ...string) Route
	Host(tpl string) Route
	Match(req *http.Request, match *mux.RouteMatch) bool
	MatcherFunc(f mux.MatcherFunc) Route
	Methods(methods ...string) Route
	Name(name string) Route
	Path(tpl string) Route
	PathPrefix(tpl string) Route
	Queries(pairs ...string) Route
	Schemes(schemes ...string) Route
	Subrouter() Router
	URL(pairs ...string) (*url.URL, error)
	URLHost(pairs ...string) (*url.URL, error)
	URLPath(pairs ...string) (*url.URL, error)
}
type Router interface {
	BuildVarsFunc(f mux.BuildVarsFunc) Route
	Get(name string) Route
	GetRoute(name string) Route
	Handle(path string, handler http.Handler) Route
	HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) Route
	Headers(pairs ...string) Route
	Host(tpl string) Route
	Match(req *http.Request, match *mux.RouteMatch) bool
	MatcherFunc(f mux.MatcherFunc) Route
	Methods(methods ...string) Route
	NewRoute() Route
	Path(tpl string) Route
	PathPrefix(tpl string) Route
	Queries(pairs ...string) Route
	Schemes(schemes ...string) Route
	ServeHTTP(w http.ResponseWriter, req http.Request)
	StrictSlash(value bool) Router
}
