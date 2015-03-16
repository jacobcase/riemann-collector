package main

import (
    "log"
//    "github.com/golang/protobuf/proto"
    "fmt"
    "riemannd/runner"
    "riemannd/config"
    "riemannd/rproto"
)



func main() {
    
    
    confPaths, err := config.GetConfigPaths()

    if err != nil {
        log.Fatalln(err)
    }

    rConfig, err := config.BuildRiemanndConfig(confPaths)

    if err != nil {
        log.Fatalln(err)
    }

    fmt.Println(rConfig)

    eventChan := make(chan *rproto.Event, 100)

    err = runner.RunEventOnce(eventChan, rConfig.Events[0])

    if err != nil {
        log.Fatalln(err)
    }

    //TODO: finish dis..
}
