package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pianrspn/go-rest/app/controllers"
)

// Init is ..
func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	e.GET("/employee", controllers.FetchAllEmployee)

	return e
}
