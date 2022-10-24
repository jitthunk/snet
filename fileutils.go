package main

import "os"

func GetFiles(path string) []string {
	files, err := os.ReadDir(path)

	var result []string
	if err == nil {

		for _, f := range files {
			result = append(result, f.Name())
		}
	}

	return result
}
