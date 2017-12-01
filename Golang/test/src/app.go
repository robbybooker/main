package main

import (
	"controller"
	"fmt"
	"routes"
	"channels"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"status_route"
)

func main() {
	r := routes.Routes{
		Controller: controller.New(),
		Channels:   channels.New(),
	}

	//number := 101
	//r.Controller.IncrementNumber(&number)

	//routes.Controller.Error()

	// r.Controller.GetDbVersion()

	//fmt.Println(number)

	//r.Controller.IncrementCounter();
	//r.Controller.ShowCounter();

	//c := make(chan int)
	//go r.Channels.Moo(c)
	//channelValue := <-c
	//fmt.Println(channelValue)

	api := createRoutes(r.Controller)
	http.ListenAndServe(":8000", api.MakeHandler())

	fmt.Println("COMPLETE")
}

func createRoutes(cc controller.ControllerInterface) *rest.Api {
	api := rest.NewApi()
	//api.Use(middleware.NewStack(VERSION, BUILD_NUMBER, false, false)...)
	//api.Use(auth.NewAuthMiddleware(cfg.Auth.PublicKey, true, false))

	statusInformation := status_route.New("v1.0")

	router, err := rest.MakeRouter(
		statusInformation.GetRoute(),
		rest.Get("/dbversion", cc.GetDbVersion),
	)

	if err != nil {
		error := fmt.Sprintf("ERROR: %s", error.Error)
		fmt.Println(error)
	}

	api.SetApp(router)
	return api
}
