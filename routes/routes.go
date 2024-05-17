package routes

import (
	"crud-api/controllers"

	"github.com/labstack/echo/v4"
)

func Routes() (e *echo.Echo) {
	e = echo.New()
	// Route Product
	productRoutes := e.Group("/product")
	productRoutes.GET("", controllers.ReadAllProducts)
	productRoutes.POST("/create", controllers.CreateProduct)
	productRoutes.GET("/:id", controllers.ReadDetailProducts)
	productRoutes.DELETE("/:id", controllers.DeleteProduct)
	productRoutes.PUT("/update", controllers.UpdateProduct)

	//Route Category
	categoryRoutes := e.Group("/category")
	categoryRoutes.GET("", controllers.ReadAllCategorys)
	categoryRoutes.GET("/:id", controllers.ReadDetailCategorys)
	categoryRoutes.POST("/create", controllers.CreateCategory)
	categoryRoutes.PUT("/update", controllers.UpdateCategory)
	categoryRoutes.DELETE("/delete/:id", controllers.DeleteCategory)
	return
}
