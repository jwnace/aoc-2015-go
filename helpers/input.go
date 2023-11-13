package helpers

import (
	"os"
	"path"
	"runtime"
)

func ReadInput() string {
	_, file, _, _ := runtime.Caller(1)
	dir := path.Dir(file)
	input, err := os.ReadFile(dir + "/input.txt")

	if err != nil {
		return ""
	}

	return string(input)
}
