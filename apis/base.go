package apis

import "github.com/dlbarduzzi/pocketbook/core"

func baseRouter(app core.App) *router {
	router := newRouter(app)
	router.prefix = "/api"

	bindBooksAPI(router)
	bindHealthAPI(router)

	return router
}
