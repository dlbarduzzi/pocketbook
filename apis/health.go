package apis

import (
	"net/http"

	"github.com/dlbarduzzi/pocketbook/tools/router"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello, health!"))
}

func bindHealthApi(rg *router.RouterGroup) {
	sub := rg.Group("/health")
	sub.GET("", healthHandler)
}
