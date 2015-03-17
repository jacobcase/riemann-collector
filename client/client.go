package client

import (
    "riemannd/rproto"
    "riemannd/config"
    "net"
    "log"
    "github.com/golang/protobuf/proto"
)

var clientChans []chan *rproto.Event


func RunClients(servers []config.ServerConfig, eventChan chan *rproto.Event) {
        
    for _, conf := range servers {
        tmpChan := make(chan *rproto.Event, 200)
        clientChans = append(clientChans, tmpChan)
        go EventClientTCP(conf, tmpChan)        
    }


    go func() {
        // Since the select is non-blocking, this shouldn't need a wait group
        // like...it won't block
        for ev := range eventChan {

            for _, c := range clientChans {
                select {
                    //non-blocking, if a host can't keep up, tough luck
                    case c <-ev:
                        
                    default:
                }
            }         
        }

        for _, c := range clientChans {
            close(c)
        }

    }()
}


func EventClientTCP(conf config.ServerConfig, events chan *rproto.Event) {
    conn, _ := net.Dial("tcp", conf.Host)
    
    //TODO: handle error, keep trying

    for ev := range events {
        log.Println("Event recieved")
        data, err := proto.Marshal(ev)
        if err != nil {
            log.Println(err)
            continue
        }
        _, err = conn.Write(data)
        if err != nil {
            log.Println(err)
        }
    }    
}
