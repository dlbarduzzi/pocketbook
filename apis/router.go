package apis

import (
	"fmt"
	"net/http"

	"github.com/dlbarduzzi/pocketbook/core"
)

type httpMethod string

const (
	methodGet httpMethod = http.MethodGet
)

type route struct {
	path   string
	event  func(*core.EventRequest)
	method httpMethod
}

type router struct {
	app    core.App
	prefix string
	routes []route
}

func newRouter(app core.App) *router {
	return &router{app: app}
}

func (r *router) add(path string, event func(*core.EventRequest), method httpMethod) *route {
	newRoute := &route{
		path:   path,
		event:  event,
		method: method,
	}

	r.routes = append(r.routes, *newRoute)

	return newRoute
}

func (r *router) get(path string, event func(*core.EventRequest)) *route {
	return r.add(path, event, methodGet)
}

func (r *router) buildHandler() (http.Handler, error) {
	mux := http.NewServeMux()

	for _, route := range r.routes {
		pattern := fmt.Sprintf("%s %s", route.method, route.path)
		mux.HandleFunc(pattern, func(res http.ResponseWriter, req *http.Request) {
			route.event(&core.EventRequest{
				App:      r.app,
				Request:  req,
				Response: res,
			})
		})
	}

	return mux, nil
}
