package models

import "github.com/gin-gonic/gin"

type BodyInterface interface {
	// Fill(url string, body []byte)
	Write(c *gin.Context) error
	Prepare() error
	Validate() error
}
