package main

import (
	"flag"
	"fmt"

	"github.com/helderfarias/go-config-server/internal/endpoint"
	"github.com/helderfarias/go-config-server/internal/middleware"
	"github.com/labstack/echo"
)

func main() {
	httpAddr := flag.String("addr", "", "HTTP Listen Address")
	httpPort := flag.String("port", "3005", "HTTP Listen Port")
	appConfig := flag.String("config", "./configs/application.yml", "Aplication configs")
	flag.Parse()

	e := echo.New()
	e.Use(middleware.LoadAppConfig(*appConfig))
	endpoint.Register(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", *httpAddr, *httpPort)))
}
