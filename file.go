package main

import (
	"bufio"
	"os"
)

type LineFile struct {
	Name string
	Content string
}

func ReadFileByLine(path string) ([]LineFile, error) {
	var linesFile []LineFile
	
	file, err := os.Open(path)
    if err != nil {
        return linesFile, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // optionally, resize scanner's capacity for lines over 64K, see next example
    for scanner.Scan() {
        // fmt.Println(scanner.Text())

		content := scanner.Text()

		if content != "" {
			linesFile = append(linesFile, LineFile{
				Content: content,
				Name: content[:20],
			})
		}
    }

    if err := scanner.Err(); err != nil {
        return linesFile, err
    }

	return linesFile, nil
}