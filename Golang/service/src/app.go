package main

import (
	"fmt"
	"routes"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"status_route"
)

func main() {
	api := createRoutes(routes.New())
	http.ListenAndServe(":8001", api.MakeHandler())
	fmt.Println("COMPLETE")
}

func createRoutes(routes *routes.Routes) *rest.Api {
	api := rest.NewApi()
	//api.Use(middleware.NewStack(VERSION, BUILD_NUMBER, false, false)...)
	//api.Use(auth.NewAuthMiddleware(cfg.Auth.PublicKey, true, false))

	statusInformation := status_route.New("v1.0")

	router, err := rest.MakeRouter(
		statusInformation.GetRoute(),
		rest.Get("/dbversion", routes.GetDbVersion),
		rest.Get("/counter", routes.GetCounter),
	)

	if err != nil {
		e := fmt.Sprintf("ERROR: %s", error.Error)
		fmt.Println(e)
	}

	api.SetApp(router)
	return api
}
