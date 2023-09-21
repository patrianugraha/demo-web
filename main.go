// ============================================================================
// Copyright 2021 PT ALTO NETWORK, All Rights Reserved
// This source code is protected by Indonesian and International copyright laws.
// Any reproduction, modification, disclosure and/or distribution of the source
// code in any form is strictly prohibited and may be unlawful without
// PT ALTO Network's written consent.
// All other copyright or ALTO trademark, including but not limited to this
// source code, is PT ALTO NETWORK's property.
// ============================================================================

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
	responeBody := fmt.Sprintf(htmlBody, bgColor, greetingText, currentTime)

	return c.HTML(http.StatusOK, responeBody)
}
