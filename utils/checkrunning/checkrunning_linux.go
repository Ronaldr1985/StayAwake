//go:build linux

package checkrunning

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func IsRunning(app_name string) (bool, error) {
	files, err := os.ReadDir("/proc")
	if err != nil {
		log.Fatal("Could not open directory /proc")
	}

	for _, file := range files {
		current_pid, err := strconv.Atoi(file.Name())
		if err == nil && current_pid != os.Getpid() {
			process_files, err := os.ReadDir("/proc/" + file.Name())
			if err != nil {
				return false, err
			}
			for _, process_file := range process_files {
				if strings.Contains(process_file.Name(), "comm") {
					process_name, err := ioutil.ReadFile("/proc/" + file.Name() + "/" + process_file.Name())
					if err != nil {
						return false, err
					} else {
						if strings.Compare(strings.Trim(string(process_name), "\n"), app_name) == 0 {
							return true, nil
						}
					}
				}
			}
		}
	}
	return false, nil
}
