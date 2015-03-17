package main

import (
    "log"
//    "github.com/golang/protobuf/proto"
    "riemannd/event"
    "riemannd/config"
    "riemannd/client"
    "os/signal"
    "os"
    "syscall"
)



func main() {
    
    sigs := make(chan os.Signal, 1)

    confPaths, err := config.GetConfigPaths()

    if err != nil {
        log.Fatalln(err)
    }

    rConfig, err := config.BuildRiemanndConfig(confPaths)

    if err != nil {
        log.Fatalln(err)
    }


    // The eventIn channel is what events are recieved on which is filled
    // by the go routines. The eventDone channel is just closed when shutting
    // down to wait for the go routines 

    event.RunEvents(rConfig.Events)
    client.RunClients(rConfig.Servers, event.EventChan)

    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    log.Println("Running: waiting for kill signal")
    for range sigs {
        log.Println("Signal recieved")
        event.StopEvents()
        close(event.EventChan)
        break
    }
}
