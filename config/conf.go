package config

import (
	"errors"
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
        "github.com/Sirupsen/logrus"
        "path" 
)

type EventConfig struct {
	Cmd         string
	Interval    float64
        WorkingDir  string
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


func parseConfig(filePath string) (*CollectorConfig) {
        
        log := GetLogEntry("config")

	stat, e := os.Stat(*filePath)

        if os.IsNotExist(e) {
		log.WithFields(logrus.Fields{"path": filePath}).Errorln("Path to config file doesn't exist")
	}

	data, err := ioutil.ReadFile(filePath)

        if err != nil {
            log.WithFields(logrus.Fields{"error": err, "path": filePath}).Errorln("An error occured reading the config file")
        }

        collectConfig := &CollectorConfig{}

        err = yaml.Unmarshal(data, collectConfig)
        if err != nil {
            log.WithFields(logrus.Fields{"error": err, "path": filePath}).Errorln("An error occured parsing the config file")
        }

        return collectConfig

}


//scanned must be sorted
func getConfigPaths(configPaths []string, workDir string, scanned map[string]CollectorConfig){
    
    fullPaths := make([]string, 0, 0)

    for _, p := range configPaths {
        if path.IsAbs(p) {
            fullPaths = append(fullPaths, p)
        } else {
            fullPaths = append(fullPaths, path.Join(workDir, p))
        }
    }

    fullFilePaths := make([]string, 0, 0)
    for _, p := range fullPaths {
    
        info, err := os.Stat(p)
        if err != nil {


    }


    for _, p := range fullPaths {

    
        if val, ok := scanned[p]; !ok{
            scanned[p] = parseConfig(p)
        } else {
            continue
        }
        

        
    }
}

func BuildCollectorConfig() (*CollectorConfig) {

        log := GetLogEntry("config")

	mainPath := flag.String("c", nil, "Path to yaml config file or directory")

	flag.Parse()

        if mainPath == nil {
            log.Errorln("Path to configuration file not provided")
            flag.Usage()
            os.Exit(1)
        }

	//TODO: merge into master
	masterConfig := parseConfig(mainPath)

        scanned := make([]string, 0, 10)

        unscanned := make([]string, 0, 10)

        //Since the input must be a file, and it would fail by now if it wasn't, start with
        //the working directory being its dir
        workDir := path.Dir(mainPath)

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
