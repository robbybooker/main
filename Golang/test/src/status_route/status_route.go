package status_route

import (
	"github.com/ant0ine/go-json-rest/rest"
)

type Status map[string]interface{}

type status struct {
	version string
}

func New(version string) status {
	return status{version: version}
}

func (this *status) GetStatus() Status {
	return Status{
		"version": this.version,
	}
}

func (this *status) GetRoute() *rest.Route {
	return &rest.Route{
		HttpMethod: "GET",
		PathExp:    "/status",
		Func: func(w rest.ResponseWriter, r *rest.Request) {
			w.WriteJson(this.GetStatus())
		},
	}
}
