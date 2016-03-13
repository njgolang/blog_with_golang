package main

import (
	"fmt"
	"path/filepath"
	"os"
	"strings"
)


func getFilelist(path string) []string{
	var mdList []string

	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if (f == nil) {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if (strings.HasSuffix(path, ".md")) {
			mdList = append(mdList, path)
		}
		
		return nil
	})

	if err != nil {
		fmt.Printf("err")
	}
	return mdList
}
