package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

const importPath = "github.com/gaillard/go-online-linear-regression/v1"

func main() {
	outBytes, err := exec.Command("go", "test", "-cover", importPath).CombinedOutput()
	out := string(outBytes)
	if err != nil {
		fmt.Print(out)
		os.Exit(1)
	}

	percentage := regexp.MustCompile("coverage: (.*)%").FindStringSubmatch(out)[1]

	if percentage != "100.0" {
		fmt.Println("code coverage must be 100%")
		os.Exit(1)
	}

	outBytes, err = exec.Command("go", "vet", importPath).CombinedOutput()
	if err != nil {
		fmt.Print(string(outBytes))
		os.Exit(1)
	}
}
