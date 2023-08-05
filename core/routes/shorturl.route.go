package routes

import echo2 "github.com/labstack/echo/v4"

func AllRoute(e echo2.Echo) {
	GroupRoute("s", e)
}

func GroupRoute(path string, e echo2.Echo) {
	g := e.Group(path)

	g.GET("/:code", func(c echo2.Context) error {

	})
}
