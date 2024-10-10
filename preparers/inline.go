package preparers

func NewInlinePreparer(name string) InlinePreparer {
	return InlinePreparer{
		name: name,
	}
}

type InlinePreparer struct {
	name string
}

func (i InlinePreparer) Prepare(data string) error {
	return nil
}
func (i InlinePreparer) Name() string {
	return i.name
}

func (i InlinePreparer) ModyfyContent(data string) []byte {
	return []byte(data)
}
