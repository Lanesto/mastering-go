package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: %s <ip_addr1> [<ip_addr2> [<ip_addr3> ...]]\n", filepath.Base(args[0]))
		return
	}

	IPv4 := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[0-9]?[0-9]){0,3}`)
	for _, address := range args[1:] {
		if !IPv4.MatchString(address) {
			continue
		}

		matchAll := IPv4.FindAllStringSubmatch(address, -1)
		validFields := make([]string, 0, 4)
		invalidFields := make([]string, 0, 4)
		for _, match := range matchAll {
			field, submatch := match[0], match[1]
			if field == submatch {
				validFields = append(validFields, field)
			} else {
				invalidFields = append(invalidFields, field)
			}
		}

		fmt.Printf("checks done for given address %q\n", address)
		if len(validFields) > 0 {
			fmt.Printf("  valid IPv4 fields: %s\n", strings.Join(validFields, " "))
		}
		if len(invalidFields) > 0 {
			fmt.Printf("  invalid IPv4 fields: %s\n", strings.Join(invalidFields, " "))
		}
	}
}
