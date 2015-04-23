package main

import (
	"fmt"
	"github.com/raguay/goAlfred"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// Search Repogitory
func main() {
	execPath, err := exec.LookPath("ghq")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stdout, err := exec.Command(execPath, "list").Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	for i, path := range strings.Split(string(stdout), "\n") {
		paths := strings.Split(path, "/")
		url := "https://" + path
		if len(os.Args) > 1 {
			if m, _ := regexp.MatchString(os.Args[1], paths[len(paths)-1]); !m {
				continue
			}
		}
		goAlfred.AddResult(strconv.Itoa(i), url, paths[len(paths)-1], path, "", "yes", "", "")
	}

	fmt.Print(goAlfred.ToXML())
}
