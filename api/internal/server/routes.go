package api

import (
	"github.com/labstack/echo/v4"
	"github.com/lucas-simao/kubernetes/api/internal/domain/customers"
	"github.com/lucas-simao/kubernetes/api/internal/server/handlers"
)

func addRoutes(e *echo.Echo, cs customers.Service) {
	v1 := e.Group("/v1")

	v1.POST("/customers", handlers.PostCustomer(cs))
	v1.GET("/customers", handlers.GetCustomers(cs))
	v1.GET("/customers/:id", handlers.GetCustomerById(cs))
	v1.PATCH("/customers/:id", handlers.PatchCustomerById(cs))
	v1.DELETE("/customers/:id", handlers.DeleteCustomerById(cs))
}
