package utils

import (
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"strings"
)

// ListFilesByExtension returns all json files form specified path
func ListFilesByExtension(dataPath, dataExt string) []string {
	files, err := ioutil.ReadDir(dataPath)
	if err != nil {
		log.Fatal(err)
	}

	var result []string

	for _, f := range files {
		basename := f.Name()
		if path.Ext(basename) == dataExt {
			basename = strings.TrimSuffix(basename, filepath.Ext(basename))
			result = append(result, basename)
		}
	}
	return result
}
