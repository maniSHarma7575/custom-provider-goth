package main

import (
	"net/http"
	"oauth-go/providers/elitmus"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func main() {
	goth.UseProviders(
		elitmus.New("bllnH9z7WbTqEra1yebnIZspFxwOupDrnpZV-wA0skY", "B5W_qnkvwy_JLFfE7gjrLLCxLVVyIQ-U0KtuuchdVRM", "http://localhost:8080/auth/elitmus/callback", "localhost:3000"),
	)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/auth/:provider/callback", func(c echo.Context) error {
		user, err := gothic.CompleteUserAuth(c.Response().Writer, c.Request())

		c.Logger().Error(err)
		if err != nil {
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}
		return c.JSON(http.StatusOK, user)
	})

	e.GET("/auth/:provider", func(c echo.Context) error {
		provider := c.Param("provider")
		if provider == "" {
			return c.String(http.StatusBadRequest, "Provider not specified")
		}

		q := c.Request().URL.Query()
		q.Add("provider", c.Param("provider"))
		c.Request().URL.RawQuery = q.Encode()

		req := c.Request()
		res := c.Response().Writer

		gothic.BeginAuthHandler(res, req)
		return nil
	})

	e.Logger.Fatal(e.Start(":8080"))
}
