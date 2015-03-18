package client

import (
    "riemannd/rproto"
    "riemannd/config"
    "net"
    "log"
    "github.com/golang/protobuf/proto"
    "encoding/binary"
)

var clientChans []chan *rproto.Event


func RunClients(servers []config.ServerConfig, eventChan chan *rproto.Event) {
        
    for _, conf := range servers {
        tmpChan := make(chan *rproto.Event, 10)
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
    addr, _ := net.ResolveTCPAddr("tcp", conf.Host)
    conn, _ := net.DialTCP("tcp", nil, addr)


//    eventBuf := proto.NewBuffer(nil)

    sizeBytes := make([]byte, 4, 4)

//    msgBuf := proto.NewBuffer(nil)
//    msgBytes := make([]byte, 256, 256)
//    msgMsg := &rproto.Msg{}

//    msgEvent := &rproto.Event{}

    for ev := range events {
        log.Println("Event recieved: ", ev)

//        err = eventBuf.Marshal(ev)
        data, err := proto.Marshal(ev)
        if err != nil {
            panic(err)
        }

//        binary.BigEndian.PutUint32(sizeBytes, uint32(len(eventBuf.Bytes())))
        binary.BigEndian.PutUint32(sizeBytes, uint32(len(data)))
        log.Println("Buffer size measured: ", len(data))

//        _, err = conn.Write(sizeBytes)
        log.Println("Buffer size: ", sizeBytes)

        if err != nil {
            panic(err)
        }

        sizeBytes = append(sizeBytes, data...)
        _, err = conn.Write(sizeBytes)

        if err != nil {
            panic(err)
        }

        sizeBytes = sizeBytes[:4]

//        proto.Unmarshal(data, msgEvent)

//        log.Println("Event converted: ", msgEvent)

//        eventBuf.Reset()
//        sizeBytes := msgBytes[:4]
//        conn.Read(sizeBytes)
//        msgBuf.SetBuf(sizeBytes)
//        msgSize, _ := msgBuf.DecodeFixed32()
//        log.Println("Msg size recieved, size: ", msgSize)
//
//
//        if msgSize > uint64(len(msgBytes)) {
//            msgBytes = make([]byte, msgSize, msgSize)
//        }
//
//        conn.Read(msgBytes[:msgSize])
//        msgBuf.SetBuf(msgBytes[:msgSize])
//
//        msgBuf.Unmarshal(msgMsg)
//
//        log.Println(msgMsg)

        

    }    
}
