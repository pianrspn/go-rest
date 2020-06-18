package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pianrspn/go-rest/app/models"
)

// FetchAllEmployee is ..
func FetchAllEmployee(c echo.Context) error {
	result, err := models.FetchAllEmployee()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
