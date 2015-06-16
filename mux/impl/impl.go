package mux

import (
	"../adt"
	"github.com/gorilla/mux"
	"net/http"
)

func NewFactory() imux.Factory {
	return &Factory{}
}

type Factory struct {
}

func (f *Factory) Vars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func (f *Factory) CurrentRoute(r *http.Request) imux.Route {
	return &Route{mux.CurrentRoute(r)}
}

func (f *Factory) NewRouter() imux.Router {
	return &Router{mux.NewRouter()}
}

type Route struct {
	route *mux.Route
}

func (r *Route) BuildOnly() Route {

}

func (r *Route) BuildVarsFunc(f mux.BuildVarsFunc) Route {

}

func (r *Route) GetError() error {

}

func (r *Route) GetHandler() http.Handler {

}

func (r *Route) GetName() string {

}

func (r *Route) Handler(handler http.Handler) Route {

}

func (r *Route) HandlerFunc(f func(http.ResponseWriter, *http.Request)) Route {

}

func (r *Route) Headers(pairs ...string) Route {

}

func (r *Route) Host(tpl string) Route {

}

func (r *Route) Match(req *http.Request, match *mux.RouteMatch) bool {

}

func (r *Route) MatcherFunc(f mux.MatcherFunc) Route {

}

func (r *Route) Methods(methods ...string) Route {

}

func (r *Route) Name(name string) Route {

}

func (r *Route) Path(tpl string) Route {

}

func (r *Route) PathPrefix(tpl string) Route {

}

func (r *Route) Queries(pairs ...string) Route {

}

func (r *Route) Schemes(schemes ...string) Route {

}

func (r *Route) Subrouter() Router {

}

func (r *Route) URL(pairs ...string) (*url.URL, error) {

}

func (r *Route) URLHost(pairs ...string) (*url.URL, error) {

}

func (r *Route) URLPath(pairs ...string) (*url.URL, error) {

}

type Router struct {
	router *mux.Router
}

func (r *Router) BuildVarsFunc(f mux.BuildVarsFunc) Route {

}

func (r *Router) Get(name string) Route {

}

func (r *Router) GetRoute(name string) Route {

}

func (r *Router) Handle(path string, handler http.Handler) Route {

}

func (r *Router) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) Route {

}

func (r *Router) Headers(pairs ...string) Route {

}

func (r *Router) Host(tpl string) Route {

}

func (r *Router) Match(req *http.Request, match *mux.RouteMatch) bool {

}

func (r *Router) MatcherFunc(f mux.MatcherFunc) Route {

}

func (r *Router) Methods(methods ...string) Route {

}

func (r *Router) NewRoute() Route {

}

func (r *Router) Path(tpl string) Route {

}

func (r *Router) PathPrefix(tpl string) Route {

}

func (r *Router) Queries(pairs ...string) Route {

}

func (r *Router) Schemes(schemes ...string) Route {

}

func (r *Router) ServeHTTP(w http.ResponseWriter, req http.Request) {

}

func (r *Router) StrictSlash(value bool) Router {

}
