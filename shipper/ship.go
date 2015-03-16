package shipper




func RunShippersForever(servers []ServerConfig, events chan *rproto.Event, done chan bool) {
        
    shipChans := [len(servers)]chan *rproto.Event

    for i, conf := range servers {
        shipChans[i] := make(chan *rproto.Event, 200)
        go ShipEventsTCP(conf, shipChans[i])        
    }

    for {
        select {
            case ev = <-events {
                for c := range shipChans {
                    select {
                        //non-blocking, if a host can't keep up, tough luck
                        case c <- ev:
                        default:
                    }
                }    
                  
            }

            case stop = <- done {
                break
            }
        }    

    }
}

func ShipEventsTCP(conf ServerConfig, events chan *rproto.Event, done chan bool) {
    
    conn, err := net.Dial("tcp", conf.Host)
    
    //TODO: handle error, keep trying

    for {
        
    
    }    

}
