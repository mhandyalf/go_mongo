package routes

import (
	"go_mongo/controllers"

	"github.com/labstack/echo/v4"
)

func SetEmployeeRoutes(e *echo.Echo) {
	e.GET("/employees", controllers.GetEmployees)
	e.GET("/employees/:id", controllers.GetEmployee)
	e.POST("/employees", controllers.CreateEmployee)
	e.PUT("/employees/:id", controllers.UpdateEmployee)
	e.DELETE("/employees/:id", controllers.DeleteEmployee)
}
