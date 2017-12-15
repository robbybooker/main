package routes

import (
	"maths"
	"channels"
	"mariadb"
	"github.com/ant0ine/go-json-rest/rest"
)

type Routes struct {
	Maths    maths.MathsInterface
	Channels channels.ChannelsInterface
	Database mariadb.DatabaseInterface
}

func New() *Routes {
	return &Routes{
		Maths:    maths.New(),
		Channels: channels.New(),
		Database: mariadb.New(),
	}
}

func (this *Routes) GetDbVersion(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(this.Database.GetDbVersion())
}

func (this *Routes) GetCounter(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(this.Maths.GetCounter())
}
