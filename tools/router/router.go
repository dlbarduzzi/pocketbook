package router

import (
	"errors"
	"net/http"
)

type HttpMethod string

const (
	MethodGet  HttpMethod = http.MethodGet
	MethodPost HttpMethod = http.MethodPost
)

type handlerFunc func(w http.ResponseWriter, r *http.Request)

type Route struct {
	Path    string
	Method  HttpMethod
	Handler handlerFunc
}

type RouterGroup struct {
	Prefix   string
	children []any // Route or RouterGroup itself.
}

type Router struct {
	*RouterGroup
}

func NewRouter() *Router {
	return &Router{
		RouterGroup: &RouterGroup{},
	}
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	newGroup := &RouterGroup{Prefix: prefix}
	group.children = append(group.children, newGroup)
	return newGroup
}

func (rg *RouterGroup) Route(path string, method HttpMethod, handler handlerFunc) *Route {
	route := &Route{
		Path:    path,
		Method:  method,
		Handler: handler,
	}

	rg.children = append(rg.children, *route)

	return route
}

func (rg *RouterGroup) GET(path string, handler handlerFunc) *Route {
	return rg.Route(path, http.MethodGet, handler)
}

func (r *Router) BuildHandler() (http.Handler, error) {
	mux := http.NewServeMux()

	if err := r.loadRoutes(mux, r.RouterGroup, nil); err != nil {
		return nil, err
	}

	return mux, nil
}

func (r *Router) loadRoutes(mux *http.ServeMux, group *RouterGroup, parents []*RouterGroup) error {
	for _, child := range group.children {
		switch v := child.(type) {
		case Route:
			var pattern string

			if v.Method != "" {
				pattern = string(v.Method) + " "
			}

			for _, parent := range parents {
				pattern += parent.Prefix
			}

			pattern += group.Prefix
			pattern += v.Path

			mux.HandleFunc(pattern, v.Handler)
		case *RouterGroup:
			if err := r.loadRoutes(mux, v, append(parents, group)); err != nil {
				return err
			}
		default:
			return errors.New("invalid route group type")
		}
	}

	return nil
}
