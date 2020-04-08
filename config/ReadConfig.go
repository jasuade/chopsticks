package config

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

//Configuration of the game
type Configuration struct {
	Players    int
	Difficulty string
	Mode       string
}

//OpenConfig opens a read-only file with the game configuration
func OpenConfig(path string) (*Configuration, error) {
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	log.Println("Reading configuration file . . .")
	return ReadConfig(file), nil
}

//ReadConfig reads configuration of the game from a file
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

//parseConfig includes the read configuration into the Configuration
func parseConfig(config *Configuration, fileTextLines []string) *Configuration {
	for _, eachline := range fileTextLines {
		fields := strings.Split(eachline, ":")
		switch strings.TrimSpace(fields[0]) {
		case "players":
			{
				config.Players, _ = strconv.Atoi(strings.TrimSpace(fields[1]))
			}
		case "difficulty":
			{
				config.Difficulty = strings.TrimSpace(fields[1])
			}
		case "mode":
			{
				config.Mode = strings.TrimSpace(fields[1])
			}
		}
	}
	return config
}
