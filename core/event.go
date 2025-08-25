package core

import "net/http"

type EventRequest struct {
	App      App
	Request  *http.Request
	Response http.ResponseWriter
}
