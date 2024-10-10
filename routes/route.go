package routes

type Route struct {
	Method  string
	Content Content
	Headers map[string]string
}

func (r Route) Validate() error {
	return r.Content.validate()
}
