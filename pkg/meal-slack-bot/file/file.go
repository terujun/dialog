package file

import "os"

//FileExists will check if the file exists.
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil

}
