package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Configuration struct {
	Difficulty string
	Mode       string
}

func OpenConfig(path string) (*Configuration, error) {
	fmt.Println("Opening a file ")
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	fmt.Println("Reading the file")
	//return ReadConfig(file), nil
	return nil, nil
}

func ReadConfig(file *os.File) *Configuration {
	config := &Configuration{}

	buffer := bufio.NewScanner(file)
	buffer.Split(bufio.ScanLines)
	var fileTextLines []string
	for buffer.Scan() {
		fileTextLines = append(fileTextLines, buffer.Text())
	}
	return parseConfig(config, fileTextLines)
}

func parseConfig(config *Configuration, fileTextLines []string) *Configuration {
	for _, eachline := range fileTextLines {
		fields := strings.Split(eachline, ":")
		switch strings.TrimSpace(fields[0]) {
		case "difficulty":
			{
				config.Difficulty = strings.TrimSpace(fields[1])
			}
		case "mode":
			{
				config.Mode = strings.TrimSpace(fields[1])
			}
		}
		fmt.Println(fields)
	}
	return config
}
