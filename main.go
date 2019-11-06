package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

func main() {
	rand.Seed(42)
	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	var c Config
	yaml.NewDecoder(f).Decode(&c)
	log.Println(c)

	r := gin.Default()
	if c.CorsEnabled {
		r.Use(cors.Default())
	}
	for url, route := range c.Routes {
		func(url string, route Route) {
			method := strings.ToUpper(route.Method)
			var content []byte
			var contents [][]byte
			switch route.Content.Type {
			case ContentTypeInline:
				content = []byte(route.Content.Body)
			case ContentTypeLink:
				content, err = ioutil.ReadFile(route.Content.Link)
				if err != nil {
					log.Println("cant load content file for route " + url)
					log.Fatal(err)
				}
			case ContentTypeRandom:
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
				if route.Content.Type == ContentTypeRandom {
					content = contents[rand.Intn(len(contents))]
				}

				c.Writer.Write(content)

				c.Next()
			})
		}(url, route)
	}
	r.Run(":3333")
}

type Config struct {
	CorsEnabled bool `yaml:"corsEnabled"`
	Routes      map[string]Route
}

type Route struct {
	Method  string
	Content Content
	Headers map[string]string
}

type ContentType string

var (
	ContentTypeInline ContentType = "inline"
	ContentTypeLink   ContentType = "link"
	ContentTypeRandom ContentType = "random"
)

type Content struct {
	Type   ContentType
	Body   string
	Link   string
	Random []RandomContent
}

type RandomContent struct {
	Body string
	Link string
}
