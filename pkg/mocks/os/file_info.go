package os

import (
	"os"
	sysos "os"
	"time"
)

// FileInfo mocks the os.FileInfo interface
type FileInfo struct {
	isDir   bool
	mode    os.FileMode
	modTime time.Time
	name    string
	size    int64
	sys     interface{}
}

func (m *FileInfo) Returns(toReturn ReturnValue) *FileInfo {
	var ok bool
	switch toReturn.Function {
	case "IsDir":
		if m.isDir, ok = toReturn.Value.(bool); !ok {
			panic("invalid .Value provided to ReturnValue (expected `bool` type)")
		}
	case "Mode":
		if m.mode, ok = toReturn.Value.(sysos.FileMode); !ok {
			panic("invalid .Value provided to ReturnValue (expected `os.FileMode` type)")
		}
	case "ModTime":
		if m.modTime, ok = toReturn.Value.(time.Time); !ok {
			panic("invalid .Value provided to ReturnValue (expected `time.Time` type)")
		}
	case "Name":
		if m.name, ok = toReturn.Value.(string); !ok {
			panic("invalid .Value provided to ReturnValue (expected `string` type)")
		}
	case "Size":
		if m.size, ok = toReturn.Value.(int64); !ok {
			panic("invalid .Value provided to ReturnValue (expected `int64` type)")
		}
	case "Sys":
		m.sys = toReturn.Value
	}
	return m
}

func (m FileInfo) IsDir() bool {
	return m.isDir
}

func (m FileInfo) Name() string {
	return m.name
}

func (m FileInfo) Mode() os.FileMode {
	return m.mode
}

func (m FileInfo) ModTime() time.Time {
	return m.modTime
}

func (m FileInfo) Size() int64 {
	return m.size
}

func (m FileInfo) Sys() interface{} {
	return m.sys
}
