package file

import (
	"common-go/src/catch"
	"os"
)

func Load(filename string) string {
	return string(catch.Try1(os.ReadFile(filename)))
}
