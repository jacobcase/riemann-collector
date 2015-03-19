package config

import (
	"errors"
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type EventConfig struct {
	Cmd      string
	Interval float64
	//TODO: timeout
}

type ServerConfig struct {
	Host string
}

type CollectorConfig struct {
	//    Include []string
	Servers []ServerConfig
	Events  []EventConfig
}

func GetConfigPaths() ([]string, error) {

	path := flag.String("c", "./", "Path to yaml config file or directory")

	flag.Parse()

	stat, e := os.Stat(*path)
	if os.IsNotExist(e) {
		return nil, errors.New("Path does not exist")
	}

	if stat.IsDir() {
		//TODO: implement
		return nil, errors.New("Config directories not supported yet")
	}

	//TODO: do what i said it does...
	x := []string{*path}
	return x, nil
}

func BuildCollectorConfig(files []string) (*CollectorConfig, error) {

	//TODO: merge into master
	masterConfig := new(CollectorConfig)

	for _, path := range files {

		data, err := ioutil.ReadFile(path)
		if err != nil {
			//TODO: or skip? meh, config option later
			return nil, err
		}

		err = yaml.Unmarshal(data, masterConfig)
		if err != nil {
			//TODO: error or skip
			return nil, err
		}
	}

	return masterConfig, nil

}
