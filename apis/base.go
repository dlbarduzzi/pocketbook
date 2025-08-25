package apis

import (
	"github.com/dlbarduzzi/pocketbook/tools/router"
)

func NewRouter() *router.Router {
	router := router.NewRouter()
	apiGroup := router.Group("/api")

	// bindBooksApi(apiGroup)
	bindHealthApi(apiGroup)

	return router
}
