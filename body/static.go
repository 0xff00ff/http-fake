package body

import (
	"errors"
	"fserver/models"

	"github.com/gin-gonic/gin"
)

func NewStaticBody(data any, preparer models.Preparer) models.BodyInterface {

	return &StaticBody{
		rawData:  data,
		preparer: preparer,
	}
}

type StaticBody struct {
	preparer models.Preparer
	data     string
	rawData  any
}

func (s StaticBody) Write(c *gin.Context) error {
	d := s.preparer.ModyfyContent(s.data)
	_, err := c.Writer.Write(d)
	return err
}

func (s *StaticBody) Prepare() error {
	s.data = s.rawData.(string)
	return s.preparer.Prepare(s.data)
}

func (s StaticBody) Validate() error {
	_, ok := s.rawData.(string)
	if !ok {
		return errors.New("data should be a string")
	}
	return nil
}
