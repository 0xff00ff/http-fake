package routes

import (
	"errors"
	"fmt"

	"fserver/body"
	"fserver/models"
	"fserver/preparers"
)

type ContentType string
type ContentAlgorithm string

var (
	ContentTypeInline ContentType = "inline"
	ContentTypeFile   ContentType = "file"

	ContentAlgorithmStatic    ContentAlgorithm = "static"
	ContentAlgorithmRandom    ContentAlgorithm = "random"
	ContentAlgorithmCondition ContentAlgorithm = "condition"
)

type Content struct {
	Type      ContentType
	Algorithm ContentAlgorithm
	Body      any
}

func (c *Content) validate() error {
	if c.Type == "" {
		c.Type = ContentTypeInline
	}
	if c.Algorithm == "" {
		c.Algorithm = ContentAlgorithmStatic
	}
	if c.Body == nil {
		return errors.New("field body is missing")
	}
	return nil
}

func (c Content) GetBody() (models.BodyInterface, error) {
	var m models.BodyInterface
	var p models.Preparer

	switch c.Type {
	case ContentTypeInline:
		p = preparers.InlinePreparer{}
	case ContentTypeFile:
		p = preparers.FilePreparer{}
	default:
		return m, fmt.Errorf("type %s is unknown", c.Type)
	}

	if c.Algorithm == "" {
		c.Algorithm = ContentAlgorithmStatic
	}

	switch c.Algorithm {
	case ContentAlgorithmStatic:
		m = body.NewStaticBody(c.Body, p)
	case ContentAlgorithmRandom:
		m = body.NewRandomBody(c.Body, p)
	case ContentAlgorithmCondition:
		m = body.NewConditionBody(c.Body, p)
	default:
		return m, fmt.Errorf("algorithm '%s' is unknown", c.Algorithm)
	}
	return m, nil
}
