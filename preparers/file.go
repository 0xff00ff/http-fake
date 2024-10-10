package preparers

import (
	"errors"
	"log"
	"os"
)

func NewPreparer(name string) FilePreparer {
	return FilePreparer{
		name: name,
	}
}

type FilePreparer struct {
	name string
}

func (f FilePreparer) Prepare(data string) error {
	return checkFile(data)
}

func (f FilePreparer) Name() string {
	return f.name
}

func (f FilePreparer) ModyfyContent(data string) []byte {
	return getFileContents(data)
}

func getFileContents(file string) []byte {
	content, err := os.ReadFile(string(file))
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func checkFile(file string) error {
	info, err := os.Lstat(file)
	if err != nil {
		return err
	}
	if !info.Mode().IsRegular() {
		return errors.New("not a file")
	}
	return nil
}
