package shipper



func ShipEventsTCP(addr string, events chan *rproto.Event) {
    
    conn, err := net.Dial("tcp", addr)
    
}
