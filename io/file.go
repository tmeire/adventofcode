package io

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// ReadLinesFromFile reads all lines of a file into a slice of strings
func ReadLinesFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}
	return lines, nil
}

// ReadLinesFromFile reads all lines of a file into a slice of strings
func ReadByteLinesFromFile(filename string) ([][]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return bytes.Split(b, []byte{'\n'}), nil
}
