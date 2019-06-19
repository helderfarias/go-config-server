package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/helderfarias/go-config-server/internal/endpoint"
	mwi "github.com/helderfarias/go-config-server/internal/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/yaml.v2"
)

func main() {
	httpAddr := flag.String("addr", "", "HTTP Listen Address")
	httpPort := flag.String("port", "3005", "HTTP Listen Port")
	appConfig := flag.String("config", "./configs/application.yml", "Aplication configs")
	flag.Parse()

	cfg, err := load(*appConfig)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(mwi.SetCloudConfig(cfg))
	e.Use(mwi.AuthBasic(cfg))
	e.Use(mwi.AuthApiKey(cfg))

	endpoint.Register(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", *httpAddr, *httpPort)))
}

func load(configFileName string) (domain.SpringCloudConfig, error) {
	path, err := filepath.Abs(configFileName)
	if err != nil {
		return domain.SpringCloudConfig{}, err
	}

	file, err := os.Open(path)
	if err != nil {
		return domain.SpringCloudConfig{}, err
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return domain.SpringCloudConfig{}, err
	}

	var cloudConfig domain.SpringCloudConfig
	if err := yaml.Unmarshal(content, &cloudConfig); err != nil {
		return domain.SpringCloudConfig{}, err
	}

	return cloudConfig, nil
}
