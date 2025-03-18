package file

import (
	"os"
)

func GetFileCount(dir string) int {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 5
	}
	count := len(entries)
	if count > 10 {
		return 10
	}
	if count < 3 {
		return 3
	}
	return count
}
