package controller

import "github.com/ant0ine/go-json-rest/rest"

type ControllerInterface interface {
	IncrementNumber(*int)
	Error()
	GetDbVersion(w rest.ResponseWriter, r *rest.Request)
	IncrementCounter()
	ShowCounter()
}
