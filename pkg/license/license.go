package license

import (
	"bufio"
	"io"
	"os"
	"reflect"
)

// Result is a single instance of a validation result
type Result struct {
	FileName  string
	Validated bool
	Error     error
}

// Results is a slice of validation result
var Results []Result

// ValidateLicense verifies that the license header exists for a specific license
func ValidateLicense(files []string, license string) []Result {
	licenseHeader := getLicenseHeaders(license)

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

		if containsLicenseHeader(f, licenseHeader) {
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

// getLicenseHeaders returns the license headers for a specific license
func getLicenseHeaders(license string) []string {
	licenses := make(map[string][]string)
	licenses["ASL2"] = ASL2Header
	licenses["MIT"] = MITHeader

	return licenses[license]
}

//containsLicenseHeader checks if the license header exists
func containsLicenseHeader(file io.Reader, license []string) bool {
	var found []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		found = append(found, scanner.Text())
	}

	if len(found) < len(license) {
		return false
	}

	if !reflect.DeepEqual(found[:len(license)], license) {
		return false
	}

	return true
}
