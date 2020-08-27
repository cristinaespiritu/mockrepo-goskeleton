// Copyright 2020 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// WriteToFile - Writes the content to file
func WriteToFile(filePath string, fileName string, content string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := os.Mkdir(filePath, os.ModePerm)

		if err != nil {
			log.Fatal(err)
		}
	}

	finalFilePath := fmt.Sprintf(filePath + filePath)
	f, err := os.OpenFile(finalFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	defer f.Close()
	fmt.Fprintln(f, content)
}

// ReadFile - read file using file util and return byte array
func ReadFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	return ioutil.ReadAll(file)
}

func RootDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	return dir
}
