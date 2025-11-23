package debug

import (
	"github.com/davecgh/go-spew/spew"
)

func Dump(v ...interface{}) {
	spew.Dump(v...)
}

func DD(v ...interface{}) {
	panic(spew.Sdump(v...))
}
