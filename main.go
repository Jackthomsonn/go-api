package main

import (
	userController "jackthomson/go-api/controllers/user"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Url         string
	Method      string
	Handler     func(*gin.Context)
}

func setupRoutes(routes []Route) {
	r := gin.Default()

	for _, route := range routes {
		switch route.Method {
			case "GET":
				r.GET(route.Url, route.Handler)
			case "POST":
				r.POST(route.Url, route.Handler)
			}
	}

	r.Run()
}

func main() {
	routes := []Route{
		{
			Url: "/users",
			Method: "GET",
			Handler: userController.GetUsers,
		},
		{
			Url: "/users",
			Method: "POST",
			Handler: userController.CreateUser,
		},
		{
			Url: "/users/:name",
			Method: "GET",
			Handler: userController.GetUser,
		},
	}

	setupRoutes(routes)
}