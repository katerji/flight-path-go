package main

import (
	"github.com/katerji/flight-path-go/handler"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()

	e.POST("/find-flight-path", PostFindFlightPath)

	// Start the web server
	err := e.Start(":8080")
	if err != nil {
		panic(err)
	}
}

// PostFindFlightPath handles the HTTP request using the generic function
func PostFindFlightPath(c echo.Context) error {
	h := handler.FindFlightPathHandler{}
	return HandleRequest(c, h.Handle)
}

// HandleRequest Generic handler function to process JSON requests
func HandleRequest[Req any, Res any](c echo.Context, handler func(Req) Res) error {
	var request Req

	// Bind JSON request to struct
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "Invalid request payload",
		})
	}

	// Process the request using the provided handler function
	response := handler(request)

	// Return JSON response
	return c.JSON(http.StatusOK, response)
}
