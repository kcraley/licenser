package path

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Walk returns
func Walk(paths, exclude, args []string) ([]string, error) {
	var modify []string
	for _, path := range paths {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("Unable to access file %q: %v\n", path, err)
				return err
			}
			if !info.IsDir() && !IsExcluded(path, exclude) {
				// fmt.Printf("Adding file %q to be modified\n", path)
				modify = append(modify, path)
			} else {
				// fmt.Printf("Excluding directory: %q\n", path)
				return nil
			}
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("Error walking the path %v", err)
		}
	}
	return modify, nil
}

// IsExcluded checks if a given path is excluded
func IsExcluded(path string, exclude []string) bool {
	for _, excluded := range exclude {
		// fmt.Printf("Path: %s, Excluded: %s\n", path, excluded)
		if strings.HasPrefix(path, excluded) {
			return true
		}
	}
	return false
}

// Contains will find an entry in an array
func Contains(slice []string, element string) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
