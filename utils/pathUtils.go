package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var Cwd string

func WindowsCwd() string {
	path, _ := filepath.Abs(".")
	return path
}

func RemoveLastElementInPath(s string) string {
	allParts := strings.Split(s, string(os.PathSeparator))
	remainingParts := allParts[0 : len(allParts)-1]
	res := strings.Join(remainingParts, string(os.PathSeparator))
	return res
}
func MacCwd() string {
	path := os.Args[0]
	return RemoveLastElementInPath(path)
}

func PathTo(s string) string {
	return fmt.Sprintf("%s\\%s", Cwd, s)
}

func PathToMinus(s string, count int) string {
	p := Cwd
	for range count {
		p = RemoveLastElementInPath(p)
	}
	return fmt.Sprintf("%s\\%s", p, s)
}

func init() {
	if runtime.GOOS == "windows" {
		Cwd = WindowsCwd()
	} else {
		Cwd = MacCwd()
		// Should probably handle Linux here at some point
	}
}
