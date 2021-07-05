package routes

import (
	"github.com/akhidnukhlis/controllers"
	"github.com/akhidnukhlis/middleware"
	"github.com/labstack/echo"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is application inventory!!")
	})

	// api data product
	e.GET("/api/product", controllers.FetchAllProduct, middleware.IsAuthenticated)
	e.POST("/api/product", controllers.StoreProduct, middleware.IsAuthenticated)
	e.PUT("/api/product", controllers.UpdateProduct, middleware.IsAuthenticated)
	e.DELETE("/api/product", controllers.DeleteProduct, middleware.IsAuthenticated)

	// api data warehouse
	e.GET("/api/warehouse", controllers.FetchAllWarehouse, middleware.IsAuthenticated)
	e.POST("/api/warehouse", controllers.StoreWarehouse, middleware.IsAuthenticated)
	e.PUT("/api/warehouse", controllers.UpdateWarehouse, middleware.IsAuthenticated)
	e.DELETE("/api/warehouse", controllers.DeleteWarehouse, middleware.IsAuthenticated)

	// api data supplier
	e.GET("/api/supplier", controllers.FetchAllSupplier, middleware.IsAuthenticated)
	e.POST("/api/supplier", controllers.StoreSupplier, middleware.IsAuthenticated)
	e.PUT("/api/supplier", controllers.UpdateSupplier, middleware.IsAuthenticated)
	e.DELETE("/api/supplier", controllers.DeleteSupplier, middleware.IsAuthenticated)

	// api data source
	e.GET("/api/source", controllers.FetchAllSource, middleware.IsAuthenticated)
	e.POST("/api/source", controllers.StoreSource, middleware.IsAuthenticated)
	e.PUT("/api/source", controllers.UpdateSource, middleware.IsAuthenticated)
	e.DELETE("/api/source", controllers.DeleteSource, middleware.IsAuthenticated)

	// api data so
	e.GET("/api/so", controllers.FetchAllSo, middleware.IsAuthenticated)
	e.POST("/api/so", controllers.StoreSo, middleware.IsAuthenticated)
	e.PUT("/api/so", controllers.UpdateSo, middleware.IsAuthenticated)
	e.DELETE("/api/so", controllers.DeleteSo, middleware.IsAuthenticated)

	// api data role
	e.GET("/api/role", controllers.FetchAllRole, middleware.IsAuthenticated)
	e.POST("/api/role", controllers.StoreRole, middleware.IsAuthenticated)
	e.PUT("/api/role", controllers.UpdateRole, middleware.IsAuthenticated)
	e.DELETE("/api/role", controllers.DeleteRole, middleware.IsAuthenticated)

	// api data po
	e.GET("/api/po", controllers.FetchAllPo, middleware.IsAuthenticated)
	e.POST("/api/po", controllers.StorePo, middleware.IsAuthenticated)
	e.PUT("/api/po", controllers.UpdatePo, middleware.IsAuthenticated)
	e.DELETE("/api/po", controllers.DeletePo, middleware.IsAuthenticated)

	// api data goods
	e.GET("/api/goods", controllers.FetchAllGoods, middleware.IsAuthenticated)
	e.POST("/api/goods", controllers.StoreGoods, middleware.IsAuthenticated)
	e.PUT("/api/goods", controllers.UpdateGoods, middleware.IsAuthenticated)
	e.DELETE("/api/goods", controllers.DeleteGoods, middleware.IsAuthenticated)

	// api data category
	e.GET("/api/category", controllers.FetchAllCategory, middleware.IsAuthenticated)
	e.POST("/api/category", controllers.StoreCategory, middleware.IsAuthenticated)
	e.PUT("/api/category", controllers.UpdateCategory, middleware.IsAuthenticated)
	e.DELETE("/api/category", controllers.DeleteCategory, middleware.IsAuthenticated)

	// api data audit info
	e.GET("/api/auditing", controllers.FetchAllAuditing, middleware.IsAuthenticated)
	e.GET("/api/auditing", controllers.StoreAuditing, middleware.IsAuthenticated)

	// api data audit
	e.GET("/api/audit", controllers.FetchAllAudit, middleware.IsAuthenticated)
	e.GET("/api/audit", controllers.StoreAudit, middleware.IsAuthenticated)

	// api hash & login
	e.GET("/api/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/api/login", controllers.CheckLogin)

	// api unit tests
	e.GET("/api/tests-struct-validation", controllers.TestStructValidation)
	e.GET("/api/tests-variable-validation", controllers.TestVariableValidation)

	// api unit tests product

	return e
}