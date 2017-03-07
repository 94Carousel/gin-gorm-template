package helps

import (
	"os"
	"strings"
)

// CommandPath return executable command
func CommandPath(cmd string) string {
	var cmdPath string
	envPath := os.Getenv("PATH")
	pathArray := strings.Split(envPath, ":")
	for _, path := range pathArray {
		p := path + "/" + cmd
		if info, err := os.Stat(p); err == nil {
			if (info.Mode() & 0111) > 0 {
				cmdPath = p
				break
			}
		}
	}
	return cmdPath
}
