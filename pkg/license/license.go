package license

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
)

// ValidateLicense verifies that the license header exists for a specific license
func ValidateLicense(files []string, license string) (bool, error) {
	licenseHeader := getLicenseHeaders(license)
	var result bool
	var err error

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			err = fmt.Errorf("Unable to access file: %q", file)
		}
		defer f.Close()

		if containsLicenseHeader(f, licenseHeader) {
			result = true
		} else {
			result = false
		}
	}

	return result, err
}

// getLicenseHeaders returns the license headers for a specific license
func getLicenseHeaders(license string) []string {
	licenses := make(map[string][]string)
	licenses["ASL2"] = ASL2Header

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
