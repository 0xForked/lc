package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	fileDir, fileExt, fileOutDir = "", "json", "./storage"
	helpFlag                     = []string{"-h", "--h", "--help", "help"}
	fileFlag                     = []string{"-f", "--f", "--file", "file"}
	extensionFlag                = []string{"-t", "--t", "--transform", "transform"}
	outputFlag                   = []string{"-o", "--o", "--output", "output"}
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		log.Println("show home")
		// TODO SHOW APP NAME & DESCRIPTION, ETC.
		return
	}

	if len(args) > 1 {
		lenReadFile := len(args) > 2
		lenWriteFile := len(args) > 4

		if contains(helpFlag, args[1]) {
			log.Println("show help flags")
			// TODO REGISTER NEW FUNCTION THAT SHOW HELP OPTIONS
			return
		}

		for i, arg := range args {
			switch {
			case contains(fileFlag, arg) && lenReadFile:
				fileDir = args[i+1]
			case contains(extensionFlag, arg) && lenWriteFile:
				fileExt = args[i+1]
			case contains(outputFlag, arg) && lenWriteFile:
				fileOutDir = args[i+1]
			}
		}

		if fileDir != "" {
			data := readFileContents(fileDir)
			if data == nil {
				log.Println("cannot proceed log data")
				return
			}

			err := writeFileContent(fileExt, fileOutDir, data)
			if err != nil {
				log.Println("failed to create file")
				return
			}

			log.Printf("new file created at %s/log.%s", fileOutDir, fileExt)
		}
	}
}

func contains(elements []string, element string) bool {
	for _, elem := range elements {
		if strings.Contains(element, elem) {
			return true
		}
	}
	return false
}

func readFileContents(path string) []string {
	filePath := filepath.Join(path)

	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer func(file *os.File) { _ = file.Close() }(file)

	scanner := bufio.NewScanner(file)

	var data []string
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		return nil
	}

	return data
}

func writeFileContent(extension, path string, data []string) error {
	// Open the file for writing
	filename := fmt.Sprintf("%s/log.%s", path, extension)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) { _ = file.Close() }(file)

	if extension == "json" {
		// Marshal the values into a JSON object
		// TODO FORMAT DATA
		jsonData, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}

		// Write the JSON object to a file
		_, err = io.WriteString(file, string(jsonData))
		if err != nil {
			return err
		}
	}

	if extension == "text" {
		for _, value := range data {
			_, err = io.WriteString(file, value+"\n")
			if err != nil {
				return err
			}
		}
	}

	return nil
}
