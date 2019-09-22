package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/helderfarias/go-config-server/internal/endpoint"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tylerb/graceful"
	"gopkg.in/yaml.v2"
)

func main() {
	appConfig := flag.String("config", "./configs/application.yml", "Aplication configs")
	flag.Parse()

	cfg, err := load(*appConfig)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Server.Addr = fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	logrus.Infof("⇨ http server started on %s\n", e.Server.Addr)

	setLoggerLevel(cfg.Logging.Level.Root)

	endpoint.Register(e, cfg)

	logrus.Fatal(graceful.ListenAndServe(e.Server, 5*time.Second))
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
