package filerw

//File reader/writer
//Avoid writing the same piece of code everywhere.

import "fmt"
import "os"
import "strings"

func Write(name string, content string) error{
	return os.WriteFile(name, []byte(content), 0644)
}

func Read(name string) (string, error) {
	bytes, err := os.ReadFile(name)
	if err != nil {
		return "", fmt.Errorf("Error reading file: %w", err)
	}
	return string(bytes), nil
}

func ReadFileToMap(name string) (map[string]string, error) {
	content, err := Read(name)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(content, "\n")
	result := make(map[string]string)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue 
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue 
		}

		key := parts[0]
		value := parts[1]
		result[key] = value
	}

	return result, nil
}
