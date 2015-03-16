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

    


     := [len(rConfig.Servers)]chan *rproto.Event


    for i, addr := range rConfig.Servers {
        tmpChan := make(chan *rproto.Event, 100)
        shipChans[i] = tmpChan
        go shipper.ShipEventsTCP(addr, tmpChan)
    }

    // The eventIn channel is what events are recieved on which is filled
    // by the go routines. The eventDone channel is just closed when shutting
    // down to wait for the go routines 
    eventIn := make(chan *rproto.Event, 100)
    eventDone := make(chan bool)

    //TODO: finish dis..
}
