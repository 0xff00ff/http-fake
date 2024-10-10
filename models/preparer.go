package models

type Preparer interface {
	Prepare(data string) error
	ModyfyContent(data string) []byte
	Name() string
}
