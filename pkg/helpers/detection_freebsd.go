// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
package helpers

import (
	"bufio"
	"os"
	"strings"
)

const (
	// file path without etc prefix, use host etc in init to get appropriate prefix
	DefaultBSDOsReleasePath = "/os-release"
)

var (
	osReleaseFilePath = DefaultBSDOsReleasePath
)

func GetOS() int {
	return OS_BSD
}

// detect running linux platform/distro
func GetBSDDistro() int {
	if info, err := GetBSDOSInfo(); err == nil {
		// More Specific Tests First
		if identity, ok := info["ID"]; ok {
			switch {
			case identity == "freebsd":
				return BSD_FREE
			case identity == "openbsd":
				return BSD_FREE
			}
		}
	}

	return BSD_UNKNOWN
}

func GetBSDOSInfo() (info map[string]string, err error) {
	osFile := HostEtc(osReleaseFilePath)
	file, err := os.Open(osFile)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()
	info = make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if line := strings.Split(scanner.Text(), "="); len(line) == 2 {
			// strip any surrounding quotation marks
			info[line[0]] = strings.Trim(line[1], "\"")
		}
	}
	err = scanner.Err()
	return
}
