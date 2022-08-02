package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title SunPathMicroService/DSM API 1.0
// @version 1.0
// @description Service for elevation data and coordinate tranformation.
// @termsOfService http://www.EagleView.com/terms/

// @contact.name EagleView Support
// @contact.url http://www.EagleView.com/support
// @contact.email support@EagleView.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /solar/sunpath-dsm-service

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) (err error) {
		return c.HTML(http.StatusOK, "<head><meta http-equiv=\"refresh\" content=\"0; url=/swagger/index.html\" /></head><body>p><a href=\"/swagger/index.html\">Redirect</a></p></body>")
	})

	apiBasePath := ""

	e.GET(apiBasePath+"/swagger/*", echoSwagger.WrapHandler)

	fmt.Print("check 1")
	e.Logger.Fatal(e.Start(":8000"))
}
