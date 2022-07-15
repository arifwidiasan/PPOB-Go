package handler

import (
	"github.com/CapstoneProject31/backend_ppob_31/config"
	"github.com/CapstoneProject31/backend_ppob_31/controller"
	"github.com/CapstoneProject31/backend_ppob_31/database"

	m "github.com/CapstoneProject31/backend_ppob_31/middleware"
	"github.com/CapstoneProject31/backend_ppob_31/repository"
	"github.com/CapstoneProject31/backend_ppob_31/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	e.POST("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "your request awesome",
		})
	})

	api := e.Group("/v1", middleware.CORS())

	m.LogMiddleware(e)
	api.POST("/admin/login", cont.LoginAdminController)

	api.GET("/product_types", cont.GetAllProductTypeController)
	api.POST("/product_types", cont.CreateProductTypeController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.GET("/product_types/:id", cont.GetOneProductTypeController)
	api.PUT("/product_types/:id", cont.UpdateProductTypeController, middleware.JWT([]byte(conf.JWT_KEY)))
	api.DELETE("/product_types/:id", cont.DeleteProductTypeController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.GET("/operators", cont.GetAllOperatorController)
	api.POST("/operators", cont.CreateOperatorController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.GET("/operators/:id", cont.GetOneOperatorController)
	api.PUT("/operators/:id", cont.UpdateOperatorController, middleware.JWT([]byte(conf.JWT_KEY)))
	api.DELETE("/operators/:id", cont.DeleteOperatorController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.GET("/products", cont.GetAllProductController)
	api.POST("/products", cont.CreateProductController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.GET("/products/:id", cont.GetOneProductController)
	api.PUT("/products/:id", cont.UpdateProductController, middleware.JWT([]byte(conf.JWT_KEY)))
	api.DELETE("/products/:id", cont.DeleteProductController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.GET("/products/types/:product_type_id", cont.GetProductByProductTypeController)
	api.GET("/products/operators/:operator_id", cont.GetProductByOperatorController)

	api.GET("/users", cont.GetAllUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	api.POST("/users", cont.CreateUserController)

	api.GET("/users/:id", cont.GetOneUserController)
	api.PUT("/users/:id", cont.UpdateUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	api.DELETE("/users/:id", cont.DeleteUserController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.POST("/users/login", cont.LoginUserController)

	api.GET("/payment_methods", cont.GetAllPaymentMethodController)
	api.POST("/payment_methods", cont.CreatePaymentMethodController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.GET("/payment_methods/:id", cont.GetOnePaymentMethodController)
	api.PUT("/payment_methods/:id", cont.UpdatePaymentMethodController, middleware.JWT([]byte(conf.JWT_KEY)))
	api.DELETE("/payment_methods/:id", cont.DeletePaymentMethodController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.GET("/transactions", cont.GetAllTransactionController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.GET("/transactions/:id", cont.GetOneTransactionController, middleware.JWT([]byte(conf.JWT_KEY)))
	api.PUT("/transactions/:id", cont.UpdateTransactionController, middleware.JWT([]byte(conf.JWT_KEY)))
	api.DELETE("/transactions/:id", cont.DeleteTransactionController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.GET("/users/transactions", cont.GetAllUserTransactionController, middleware.JWT([]byte(conf.JWT_KEY)))
	api.POST("/users/transactions", cont.CreateTransactionController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.GET("/users/transactions/:code_transaction", cont.GetOneUserTransactionController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.GET("/xendith_callback_payments", cont.GetAllCallbackPaymentController)
	api.POST("/xendith_callback_payments", cont.CreateCallbackPaymentController)

	api.GET("/payments/:code_transaction/:price", cont.SimulatePaymentController)

	api.GET("/transactions/:id/emails", cont.SendEmailController)
}
