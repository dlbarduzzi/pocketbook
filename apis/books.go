package apis

import (
	"fmt"

	"github.com/dlbarduzzi/pocketbook/core"
)

func booksHandler(e *core.EventRequest) {
	e.App.Logger().Info("calling db to get all books")
	_, _ = e.Response.Write([]byte("Getting all books..."))
}

func booksOneHandler(e *core.EventRequest) {
	e.App.Logger().Info("calling db to get one book")
	_, _ = e.Response.Write([]byte("Getting one book..."))
}

func bindBooksAPI(r *router) {
	sub := fmt.Sprintf("%s/v1/books", r.prefix)
	r.get(sub, booksHandler)
	r.get(sub+"/get-one", booksOneHandler)
}
