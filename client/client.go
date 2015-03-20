package client

import (
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"net"
        "github.com/Sirupsen/logrus"
	"riemann-collector/config"
	"riemann-collector/rproto"
        "time"
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

        var conn *net.TCPConn

        log := config.GetLogEntry("TCPClient")

        addr, err := net.ResolveTCPAddr("tcp", conf.Host)
        if err != nil {
            log.WithFields(logrus.Fields{"error": err, "host": conf.Host}).Fatalln("A fatal error occured resolving host")
        }

        for {
	    conn, err = net.DialTCP("tcp", nil, addr)
            if err == nil {
                break
            }
            log.WithFields(logrus.Fields{"error": err, "addr": addr}).Errorln("An error occured connecting to host, retrying in 1 second")
            time.Sleep(time.Second)
        }
        
	sizeBytes := make([]byte, 4, 4)

	msg := &rproto.Msg{
		Events: make([]*rproto.Event, 1, 1),
	}

        buf := proto.NewBuffer(nil)

	for ev := range events {
		msg.Events[0] = ev

                err := buf.Marshal(msg)
		if err != nil {
			log.WithFields(logrus.Fields{"error": err, "msg": msg}).Errorln("An error occured marshalling a message, skipping")
		}


		binary.BigEndian.PutUint32(sizeBytes, uint32(len(buf.Bytes())))

		sizeBytes = append(sizeBytes, buf.Bytes()...)
		_, err = conn.Write(sizeBytes)
		if err != nil {
                        //TODO: what do?
			panic(err)
		}

		sizeBytes = sizeBytes[:4]

	}
}
