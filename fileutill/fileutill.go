package fileutill

import "os"

func Exists(path string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
