package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func readLines() ([]string, error) {
	allInput, err := ioutil.ReadAll(os.Stdin)
	return strings.Split(string(allInput), "\n"), err
}

func outputLines(lines []string, outputFilePath string) error {
	out := os.Stdout
	var err error
	if outputFilePath != "" {
		out, err = os.OpenFile(outputFilePath, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return err
		}
		defer out.Close()
	}
	for _, line := range lines {
		_, err = fmt.Fprintln(out, line)
		if err != nil {
			return err
		}
	}
	return nil
}
