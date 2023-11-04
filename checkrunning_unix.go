//go:build unix
// +build unix

package main

import (
	"os"
	"strconv"
	"strings"
)

func IsRunning(appName string) (bool, error) {
	processFiles, err := os.ReadDir("/proc")
	if err != nil {
		return false, err
	}

	for _, file := range processFiles {
		currentPid, err := strconv.Atoi(file.Name())
		if err == nil && currentPid != os.Getpid() {
			processName, err := os.ReadFile("/proc/" + file.Name() + "/comm")
			if err != nil {
				return false, err
			}

			if strings.Compare(strings.Trim(string(processName), "\n"), appName) == 0 {
				return true, nil
			}
		}
	}

	return false, nil
}
