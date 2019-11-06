package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"

	"fserver/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

const DefaultPort = "3333"

func main() {
	rand.Seed(42)
	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	var c models.Config
	err = yaml.NewDecoder(f).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(c)

	r := gin.Default()
	if c.CorsEnabled {
		r.Use(cors.Default())
	}
	for url, route := range c.Routes {
		func(url string, route models.Route) {
			method := strings.ToUpper(route.Method)
			var content []byte
			var contents [][]byte
			switch route.Content.Type {
			case models.ContentTypeInline:
				content = []byte(route.Content.Body)
			case models.ContentTypeLink:
				content, err = ioutil.ReadFile(route.Content.Link)
				if err != nil {
					log.Println("cant load content file for route " + url)
					log.Fatal(err)
				}
			case models.ContentTypeRandom:
				// do something
				for _, val := range route.Content.Random {
					contents = append(contents, []byte(val.Body))
				}
			default:
				log.Fatal("Content type is wrong for route " + url)
			}
			r.Handle(method, url, func(c *gin.Context) {
				for key, val := range route.Headers {
					c.Header(key, val)
				}
				if route.Content.Type == models.ContentTypeRandom {
					content = getRandomContent(contents)
				}

				c.Writer.Write(content)

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

func getRandomContent(contents [][]byte) []byte {
	return contents[rand.Intn(len(contents))]
}
