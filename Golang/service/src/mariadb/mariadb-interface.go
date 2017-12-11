package mariadb

import "github.com/ant0ine/go-json-rest/rest"

type DatabaseInterface interface {
	GetDbVersion(w rest.ResponseWriter, r *rest.Request)
}
