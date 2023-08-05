package controllers

import echo2 "github.com/labstack/echo/v4"

func GetUrl(c echo2.Context) error {
	var code string = c.Param("code")
}
