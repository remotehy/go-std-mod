package data

import (
	"io"
	"os"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewData,
	wire.InterfaceValue(new(io.Reader), os.Stdin),
	wire.Value(Baz{MyStr: "hello"}),
	wire.FieldsOf(new(Baz), "MyStr"))

type DataAccess interface {
	GetData() string
}

func NewData(s string, r io.Reader) *DataAccessImpl {
	return &DataAccessImpl{s: s, r: r}
}

type DataAccessImpl struct {
	s string
	r io.Reader
}

type Baz struct {
	MyStr    string
	MyReader io.Reader
}

func (d *DataAccessImpl) GetData() string {
	return "demo-data"
}
