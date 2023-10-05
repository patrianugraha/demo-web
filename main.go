package main

import (
	"demo-devops/demo-web/utils"
	"fmt"
	"net/http"

	_ "github.com/joho/godotenv/autoload"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var htmlBody = `
<div style="background-color:%s; padding:10px;">
	<h1>
		%s
	</h1>
	<h1>
		%s
	</h1>
	<h3>
		Current time : %s
	</h3>
</div>
`

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":8888"))
}

// Handler
func hello(c echo.Context) error {
	currentTime := utils.GetCurrentTime()
	greetingText := utils.GetGreetingText()
	bgColor := utils.GetBgColor()
	const appVersion = "CANARY VERSION 4"
	responeBody := fmt.Sprintf(htmlBody, bgColor, greetingText, appVersion, currentTime)

	return c.HTML(http.StatusOK, responeBody)
}
