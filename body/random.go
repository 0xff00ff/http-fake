package body

import (
	"errors"
	"fmt"
	"fserver/models"
	"math/rand"

	"github.com/gin-gonic/gin"
)

func NewRandomBody(data any, preparer models.Preparer) models.BodyInterface {
	return &RandomBody{
		rawData: data,
		// data:     s,
		preparer: preparer,
	}
}

type RandomBody struct {
	rawData  any
	data     []string
	preparer models.Preparer
}

func (r RandomBody) Write(c *gin.Context) error {
	s := getRandomContent(r.data)
	b := r.preparer.ModyfyContent(s)
	_, err := c.Writer.Write(b)
	return err
}

func (r *RandomBody) Prepare() error {
	d := r.rawData.([]any)
	s := []string{}
	for _, v := range d {
		s = append(s, v.(string))
	}
	r.data = s
	for k, v := range r.data {
		err := r.preparer.Prepare(v)
		if err != nil {
			return fmt.Errorf("preparer \"%s\" returned an error on key %d: %w", r.preparer.Name(), k, err)
		}
	}
	return nil
}

func (r RandomBody) Validate() error {
	d, ok := r.rawData.([]any)
	if !ok {
		return errors.New("data should be a list")
	}
	for k, v := range d {
		_, ok = v.(string)
		if !ok {
			return fmt.Errorf("index %d, data should be a string", k)
		}
	}
	return nil
}

func getRandomContent(contents []string) string {
	return contents[rand.Intn(len(contents))]
}
