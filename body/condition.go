package body

import (
	"encoding/json"
	"errors"
	"fmt"
	"fserver/models"
	"io"
	"log"

	"github.com/0xff00ff/ask"
	"github.com/gin-gonic/gin"
)

type ConditionSource string

const (
	ConditionSourceBody ConditionSource = "body"
	ConditionSourcePath ConditionSource = "path"
)

func NewConditionBody(data any, preparer models.Preparer) models.BodyInterface {

	return &ConditionBody{
		rawData: data,
		// data:     d,
		preparer: preparer,
	}
}

type ConditionBody struct {
	rawData  any
	data     []ConditionData
	preparer models.Preparer
}

func (r ConditionBody) Write(c *gin.Context) error {
	// data
	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	var body any
	err = json.Unmarshal(b, &body)
	if err != nil {
		return err
	}
	log.Println("body is:", string(b))
	for k, v := range r.data {
		var val any
		if v.Source == ConditionSourceBody {
			val = ask.For(body, v.Query).Value()
		} else if v.Source == ConditionSourcePath {
			val = c.Param(v.Query)
		} else {
			return fmt.Errorf("source %s on index %d not found", v.Source, k)
		}
		if val == v.Value {
			b := r.preparer.ModyfyContent(v.Body)
			c.Writer.Write(b)
			return nil
		}
	}
	return errors.New("no condition found")
}

func (r *ConditionBody) Prepare() error {
	var d []ConditionData
	for _, v := range r.rawData.([]any) {
		d = append(d, ConditionDataFromMap(v.(map[any]any)))
	}
	r.data = d
	for k, v := range r.data {
		err := r.preparer.Prepare(v.Body)
		if err != nil {
			return fmt.Errorf("preparer %s returned an error on key %d: %w", r.preparer.Name(), k, err)
		}
	}
	return nil
}

func (r ConditionBody) Validate() error {
	d, ok := r.rawData.([]any)
	if !ok {
		return errors.New("wrong data format")
	}
	for k, v := range d {
		i, ok := v.(map[any]any)
		if !ok {
			return fmt.Errorf("index: %d, wrong data format, expecting map", k)
		}
		if i["source"] == nil {
			return fmt.Errorf("index: %d, field 'source' is absent", k)
		}
		if i["source"] != "body" && i["source"] != "path" {
			return fmt.Errorf("index: %d, field 'source' should one of: 'body', 'path'", k)
		}
		if i["query"] == nil {
			return fmt.Errorf("index: %d, field 'query' is absent", k)
		}
		if i["value"] == nil {
			return fmt.Errorf("index: %d, field 'value' is absent", k)
		}
		if i["body"] == nil {
			return fmt.Errorf("index: %d, field 'body' is absent", k)
		}
	}
	return nil
}

func MapValue(data map[string]any, key string) any {
	v := data[key]
	return v
}

type ConditionData struct {
	Query  string
	Value  any
	Body   string
	Source ConditionSource
}

func ConditionDataFromMap(m map[any]any) ConditionData {
	return ConditionData{
		Query:  m["query"].(string),
		Value:  m["value"],
		Body:   m["body"].(string),
		Source: ConditionSource(m["source"].(string)),
	}
}
