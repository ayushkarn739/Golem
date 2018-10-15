package FileSearch

import (
	"os"
	"path/filepath"
	"regexp"
)

func Find(filePath string) []string {
	var files []string
	re, _ := regexp.Compile(`([a-zA-Z0-9\s_\\.\-\(\):])+(.avi|.flv|.wmv|.mkv|.mp4)$`)
	err := filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		b := re.MatchString(path)
		if b {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}