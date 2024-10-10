package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"fserver/config"
	"fserver/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

const DefaultPort = "3333"

func main() {
	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	var c config.Config
	err = yaml.NewDecoder(f).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	if c.CorsEnabled {
		r.Use(cors.Default())
	}
	for url, route := range c.Routes {
		func(url string, route routes.Route) {
			err := route.Validate()
			if err != nil {
				log.Fatal(fmt.Errorf("validating route error: url: %s: %w", url, err))
			}
			method := strings.ToUpper(route.Method)
			body, err := route.Content.GetBody()
			if err != nil {
				log.Fatal(fmt.Errorf("getting body error: url: %s: %w", url, err))
			}
			err = body.Validate()
			if err != nil {
				log.Fatal(fmt.Errorf("validation body error: url: %s: %w", url, err))
			}
			err = body.Prepare()
			if err != nil {
				log.Fatal(fmt.Errorf("preparing error: url: %s: %w", url, err))
			}
			r.Handle(method, url, func(c *gin.Context) {
				p := c.Params
				_ = p
				for key, val := range route.Headers {
					c.Header(key, val)
				}
				err := body.Write(c)
				if err != nil {
					log.Println(err)
				}

				c.Next()
			})
		}(url, route)
	}
	port := c.Port
	if port == "" {
		port = DefaultPort
	}
	r.Run(":" + port)
}
