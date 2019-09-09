package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLines() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
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
