package apis

import (
	"fmt"

	"github.com/dlbarduzzi/pocketbook/core"
)

func healthHandler(e *core.EventRequest) {
	e.App.Logger().Info("health API called")
	_, _ = e.Response.Write([]byte("API is healthy!"))
}

func bindHealthAPI(r *router) {
	sub := fmt.Sprintf("%s/v1/health", r.prefix)
	r.get(sub, healthHandler)
}
