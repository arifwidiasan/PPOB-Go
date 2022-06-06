package handler

import (
	"github.com/CapstoneProject31/backend_ppob_31/config"
	"github.com/CapstoneProject31/backend_ppob_31/controller"
	"github.com/CapstoneProject31/backend_ppob_31/database"

	m "github.com/CapstoneProject31/backend_ppob_31/middleware"
	"github.com/CapstoneProject31/backend_ppob_31/repository"
	"github.com/CapstoneProject31/backend_ppob_31/service"

	"github.com/labstack/echo/v4"
)

func RegisterGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)
	repo := repository.NewMysqlRepository(db)

	svc := service.NewService(repo, conf)

	cont := controller.EchoController{
		Svc: svc,
	}

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "your request awesome",
		})
	})

	api := e.Group("/v1")

	m.LogMiddleware(e)
	api.POST("/admin/login", cont.LoginAdminController)

	api.POST("/product_types", cont.CreateProductTypeController)
}
