package main

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	lines, err := ReadFileByLine("sample.txt")

	if err != nil {
		t.Fatal(err)
	}

	for i, v := range lines {
		fmt.Println("i ", i)
		fmt.Println("v ", v)
	}
}