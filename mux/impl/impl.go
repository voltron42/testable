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

func (r *Route) BuildOnly() imux.Route {
	return &Route{r.route.BuildOnly()}
}

func (r *Route) BuildVarsFunc(f mux.BuildVarsFunc) imux.Route {
	return &Route{r.route.BuildVarsFunc(f)}
}

func (r *Route) GetError() error {
	return r.route.GetError()
}

func (r *Route) GetHandler() http.Handler {
	return r.route.GetHandler()
}

func (r *Route) GetName() string {
	return r.route.GetName()
}

func (r *Route) Handler(handler http.Handler) imux.Route {
	return &Route{r.route.Handler(handler)}
}

func (r *Route) HandlerFunc(f func(http.ResponseWriter, *http.Request)) imux.Route {
	return &Route{r.route.HandlerFunc(f)}
}

func (r *Route) Headers(pairs ...string) imux.Route {
	return &Route{r.route.Headers(pairs...)}
}

func (r *Route) Host(tpl string) imux.Route {
	return &Route{r.route.Host(tpl)}
}

func (r *Route) Match(req *http.Request, match *mux.RouteMatch) bool {
	return r.route.Match(req, match)
}

func (r *Route) MatcherFunc(f mux.MatcherFunc) imux.Route {
	return &Route{r.route.MatcherFunc(f)}
}

func (r *Route) Methods(methods ...string) imux.Route {
	return &Route{r.route.Methods(methods...)}
}

func (r *Route) Name(name string) imux.Route {
	return &Route{r.route.Name(name)}
}

func (r *Route) Path(tpl string) imux.Route {
	return &Route{r.route.Path(tpl)}
}

func (r *Route) PathPrefix(tpl string) imux.Route {
	return &Route{r.route.PathPrefix(tpl)}
}

func (r *Route) Queries(pairs ...string) imux.Route {
	return &Route{r.route.Queries(pairs...)}
}

func (r *Route) Schemes(schemes ...string) imux.Route {
	return &Route{r.route.Schemes(schemes...)}
}

func (r *Route) Subrouter() imux.Router {
	return &Router{r.route.Subrouter()}
}

func (r *Route) URL(pairs ...string) (*url.URL, error) {
	return r.route.URL(pairs...)
}

func (r *Route) URLHost(pairs ...string) (*url.URL, error) {
	return r.route.URLHost(pairs...)
}

func (r *Route) URLPath(pairs ...string) (*url.URL, error) {
	return r.route.URLPath(pairs...)
}

type Router struct {
	router *mux.Router
}

func (r *Router) BuildVarsFunc(f mux.BuildVarsFunc) imux.Route {
	return &Route{r.router.BuildVarsFunc(f)}
}

func (r *Router) Get(name string) imux.Route {
	return &Route{r.router.Get(name)}
}

func (r *Router) GetRoute(name string) imux.Route {
	return &Route{r.router.GetRoute(name)}
}

func (r *Router) Handle(path string, handler http.Handler) imux.Route {
	return &Route{r.router.Handle(path, handler)}
}

func (r *Router) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) imux.Route {
	return &Route{r.router.HandleFunc(path, f)}
}

func (r *Router) Headers(pairs ...string) imux.Route {
	return &Route{r.router.Headers(pairs...)}
}

func (r *Router) Host(tpl string) imux.Route {
	return &Route{r.router.Host(tpl)}
}

func (r *Router) Match(req *http.Request, match *mux.RouteMatch) bool {
	return r.router.Match(req, match)
}

func (r *Router) MatcherFunc(f mux.MatcherFunc) imux.Route {
	return &Route{r.router.MatcherFunc(f)}
}

func (r *Router) Methods(methods ...string) imux.Route {
	return &Route{r.router.Methods(methods...)}
}

func (r *Router) NewRoute() imux.Route {
	return &Route{r.router.NewRoute()}
}

func (r *Router) Path(tpl string) imux.Route {
	return &Route{r.router.Path(tpl)}
}

func (r *Router) PathPrefix(tpl string) imux.Route {
	return &Route{r.router.PathPrefix(tpl)}
}

func (r *Router) Queries(pairs ...string) imux.Route {
	return &Route{r.router.Queries(pairs...)}
}

func (r *Router) Schemes(schemes ...string) imux.Route {
	return &Route{r.router.Schemes(schemes...)}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req http.Request) {
	r.route.ServeHTTP(w, req)
}

func (r *Router) StrictSlash(value bool) imux.Router {
	return &Router{r.router.StrictSlash(value)}
}
