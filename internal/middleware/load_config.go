package middleware

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/labstack/echo"
	"gopkg.in/yaml.v2"
)

func LoadAppConfig(configFileName string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			path, err := filepath.Abs(configFileName)
			if err != nil {
				log.Println(err)
				return err
			}

			file, err := os.Open(path)
			if err != nil {
				log.Println(err)
				return err
			}

			content, err := ioutil.ReadAll(file)
			if err != nil {
				log.Println(err)
				return err
			}

			var cloudConfig domain.SpringCloudConfig
			if err := yaml.Unmarshal(content, &cloudConfig); err != nil {
				log.Println(err)
				return err
			}

			c.Set("cloudConfig", cloudConfig)
			return next(c)
		}
	}
}
