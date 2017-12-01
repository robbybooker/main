package routes

import (
	"controller"
	"channels"
)

type Routes struct {
	Controller controller.ControllerInterface
	Channels channels.ChannelsInterface
}
