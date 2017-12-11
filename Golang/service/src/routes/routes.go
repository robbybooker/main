package routes

import (
	"maths"
	"channels"
	"mariadb"
)

type Routes struct {
	Maths maths.MathsInterface
	Channels channels.ChannelsInterface
	Database mariadb.DatabaseInterface
}
