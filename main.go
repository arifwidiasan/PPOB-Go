package main

import (
	conf "github.com/CapstoneProject31/backend_ppob_31/config"
	handler "github.com/CapstoneProject31/backend_ppob_31/controller/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	config := conf.InitConfiguration()
	e := echo.New()

	handler.RegisterGroupAPI(e, config)

	e.Logger.Fatal(e.Start((config.SERVER_ADDRESS)))
}
