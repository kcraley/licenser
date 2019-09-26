package license

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"time"
)

// Result is a single instance of a validation result
type Result struct {
	FileName  string
	Validated bool
	Error     error
}

// Results is a slice of validation result
var Results []Result

// Validate verifies that the license header exists for a specific license
func Validate(files []string, author, license string) []Result {
	licenseHeader := getHeader(license)

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			Results = append(Results, Result{
				FileName:  file,
				Validated: false,
				Error:     err,
			})
		}
		defer f.Close()

		if containsHeader(f, author, licenseHeader) {
			Results = append(Results, Result{
				FileName:  file,
				Validated: true,
				Error:     nil,
			})
		} else {
			Results = append(Results, Result{
				FileName:  file,
				Validated: false,
				Error:     nil,
			})
		}
	}

	return Results
}

// getHeader returns the license headers for a specific license
func getHeader(license string) []string {
	licenses := make(map[string][]string)
	licenses["ASL2"] = ASL2Header
	licenses["MIT"] = MITHeader

	return licenses[license]
}

// containsHeader checks if the license header exists
func containsHeader(file io.Reader, author string, license []string) bool {
	var found []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		found = append(found, scanner.Text())
	}

	if len(found) < len(license) {
		return false
	}
	for key, line := range license {
		if strings.Contains(line, "%04d %s") {
			license[key] = replaceYearAuthor(line, author)
		}
	}

	if !reflect.DeepEqual(found[:len(license)], license) {
		return false
	}

	return true
}

// replaceAuthor will replace a string with the author
func replaceYearAuthor(s, author string) string {
	return fmt.Sprintf(s, time.Now().Year(), author)
}
