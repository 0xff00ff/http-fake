package models

type Route struct {
	Method  string
	Content Content
	Headers map[string]string
}
