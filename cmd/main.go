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
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func main() {
	appConfig := flag.String("config", "./configs/application.yml", "Aplication configs")
	flag.Parse()

	cfg, err := load(*appConfig)
	if err != nil {
		log.Fatal(err)
	}

	setLoggerLevel(cfg.Logging.Level.Root)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(mwi.SetCloudConfig(cfg))
	e.Use(mwi.AuthBasic(cfg))
	e.Use(mwi.AuthApiKey(cfg))

	endpoint.Register(e, cfg)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)))
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

func setLoggerLevel(src string) {
	if level, err := logrus.ParseLevel(src); err == nil {
		logrus.SetLevel(level)
	} else {
		logrus.SetLevel(logrus.ErrorLevel)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000",
		FullTimestamp:   true,
	})
}
