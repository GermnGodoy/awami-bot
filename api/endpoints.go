package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type posResponse struct {
	PrevPosition string `json:"prevPosition"`
	MovePosition string `json:"movePosition"`
}

func createEndpoints(e *echo.Echo) {

	// TESTING
	test := e.Group("/test")

	test.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "The API is working as expected.")
	})

	// RANDOM BOY
	randomBoy := e.Group("/random-boy/v1")

	randomBoy.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello I am Random Boy, use the url to ask me any chess position you want")
	})

	randomBoy.GET("/:pos", func(c echo.Context) error {
		pos := c.Param("pos")

		res := &posResponse{
			PrevPosition: pos,
			MovePosition: "1245jfdjalfjkfds",
		}

		return c.JSON(http.StatusOK, res)
	})
}
