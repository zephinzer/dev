package prompt

import (
	"io"
)

type InputError int

const (
	ErrorSkipped InputError = -1
	ErrorInput   InputError = -2
	ErrorSystem  InputError = -3
)

type InputType string

const (
	TypeInt    InputType = "int"
	TypeUInt   InputType = "uint"
	TypeString InputType = "string"
	TypeBool   InputType = "bool"
)

type InputOptions struct {
	BeforeMessage     string
	AfterMessage      string
	SerializedOptions []string
	Type              InputType
	Reader            io.Reader
	data              string
}
