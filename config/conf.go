package config

import (
	"flag"
	"github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

type EventConfig struct {
	Cmd        string
	Interval   float64
	WorkingDir string
	//TODO: timeout
}

type ServerConfig struct {
	Host string
}

type CollectorConfig struct {
	Include []string
	Servers []ServerConfig
	Events  []EventConfig
}

func parsePath(confPath string, order *[]string, loaded map[string]*CollectorConfig) {

	log := GetLogEntry("config").WithFields(logrus.Fields{"path": confPath})

	log.Debugln("Parsing Path")

	if !path.IsAbs(confPath) {
		log.Fatalln("Relative directories are not supported")
	}

	stat, err := os.Stat(confPath)

	if os.IsNotExist(err) {
		log.Fatalln("Path to config file doesn't exist")
	}

	if stat.IsDir() {
		dirPaths, _ := ioutil.ReadDir(confPath)
		for _, p := range dirPaths {
			parsePath(p.Name(), order, loaded)
		}
		return
	}

	if !stat.Mode().IsRegular() {
		log.Fatalln("Attempted to open irregular file")
	}

	if _, ok := loaded[confPath]; !ok {
		log.Warningln("Attempted to open an already opened file, could be circular dependencies. Ignoring")
		return
	}

	data, err := ioutil.ReadFile(confPath)

	if err != nil {
		log.WithFields(logrus.Fields{"error": err}).Fatalln("An error occured opening the config file")
	}

	collectConfig := &CollectorConfig{}

	err = yaml.Unmarshal(data, collectConfig)

	if err != nil {
		log.WithFields(logrus.Fields{"error": err}).Fatalln("An error occured parsing the config file")
	}

	*order = append(*order, confPath)

	loaded[confPath] = collectConfig

	for _, p := range collectConfig.Include {
		parsePath(p, order, loaded)
	}
}

func merge(master *CollectorConfig, toMerge *CollectorConfig) {
	master.Servers = append(master.Servers, toMerge.Servers...)
	master.Events = append(master.Events, toMerge.Events...)

}

func BuildCollectorConfig() *CollectorConfig {

	log := GetLogEntry("config")

	mainPath := flag.String("c", "", "Path to yaml config file or directory")

	flag.Parse()

	if len(*mainPath) == 0 {
		log.Errorln("Path to configuration file not provided")
		flag.Usage()
		os.Exit(1)
	}

	loadOrder := make([]string, 10)
	loadedConfigs := make(map[string]*CollectorConfig)

	parsePath(*mainPath, &loadOrder, loadedConfigs)

	masterConfig := &CollectorConfig{}

	for _, p := range loadOrder {

		merge(masterConfig, loadedConfigs[p])

	}

	return masterConfig

}
