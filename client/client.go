package client

import (
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"riemann-collector/config"
	"riemann-collector/rproto"
)

var clientChans []chan *rproto.Event

func RunClients(servers []config.ServerConfig, eventChan chan *rproto.Event) {

	for _, conf := range servers {
		tmpChan := make(chan *rproto.Event, 10)
		clientChans = append(clientChans, tmpChan)
		go MsgClientTCP(conf, tmpChan)
	}

	go func() {
		// Since the select is non-blocking, this shouldn't need a wait group
		// like...it won't block
		for ev := range eventChan {

			for _, c := range clientChans {
				select {
				//non-blocking, if a host can't keep up, tough luck
				case c <- ev:

				default:
				}
			}
		}

		for _, c := range clientChans {
			close(c)
		}

	}()
}

func MsgClientTCP(conf config.ServerConfig, events chan *rproto.Event) {
	addr, _ := net.ResolveTCPAddr("tcp", conf.Host)
	conn, _ := net.DialTCP("tcp", nil, addr)

	sizeBytes := make([]byte, 4, 4)

	msg := &rproto.Msg{
		Events: make([]*rproto.Event, 1, 1),
	}

	for ev := range events {
		log.Println("Event recieved: ", ev)

		msg.Events[0] = ev

		data, err := proto.Marshal(msg)
		if err != nil {
			panic(err)
		}

		binary.BigEndian.PutUint32(sizeBytes, uint32(len(data)))

		sizeBytes = append(sizeBytes, data...)
		_, err = conn.Write(sizeBytes)
		if err != nil {
			panic(err)
		}

		sizeBytes = sizeBytes[:4]

	}
}
