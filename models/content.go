package models

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
