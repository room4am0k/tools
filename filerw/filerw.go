package filerw

//File reader/writer
//Avoid writing the same piece of code everywhere.

import "fmt"
import "os"

func Write(name string, content string) error{
	return os.WriteFile(nome, []byte(content), 0644)
}

func Read(name string) (string, error) {
	bytes, err := os.ReadFile(name)
	if err != nil {
		return "", fmt.Errorf("Error reading file: %w", err)
	}
	return string(bytes), nil
}
