package main

import (
	"maths"
	"fmt"
	"routes"
	"channels"
	"mariadb"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"status_route"
)

func main() {
	r := routes.Routes{
		Maths: maths.New(),
		Channels:   channels.New(),
		Database: mariadb.New(),
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

	api := createRoutes(r)
	http.ListenAndServe(":8001", api.MakeHandler())

	fmt.Println("COMPLETE")
}

func createRoutes(routes routes.Routes) *rest.Api {
	api := rest.NewApi()
	//api.Use(middleware.NewStack(VERSION, BUILD_NUMBER, false, false)...)
	//api.Use(auth.NewAuthMiddleware(cfg.Auth.PublicKey, true, false))

	statusInformation := status_route.New("v1.0")

	router, err := rest.MakeRouter(
		statusInformation.GetRoute(),
		rest.Get("/dbversion", routes.Database.GetDbVersion),
	)

	if err != nil {
		error := fmt.Sprintf("ERROR: %s", error.Error)
		fmt.Println(error)
	}

	api.SetApp(router)
	return api
}
