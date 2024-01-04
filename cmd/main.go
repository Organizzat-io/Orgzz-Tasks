package main

import (
	"Orgzz-Tasks/config"
	"Orgzz-Tasks/internal/web"
)

func main() {
	config.OpenDbConnection()

	routeLoader := web.RouteLoader{}
	routeLoader.LoadRoutes()
}
