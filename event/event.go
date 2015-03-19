
package event

import (
    "riemann-collector/rproto"
    "riemann-collector/config"
    "os/exec"
    "fmt"
    "time"
    "sync"
    "log"
)

var wg sync.WaitGroup

var EventChan = make(chan *rproto.Event, 10)

var quitChan = make(chan bool)

func StopEvents() {
    log.Println("Ordering event runners to stop")
    close(quitChan)
    log.Println("Waiting for event runners")
    wg.Wait()
    quitChan = make(chan bool)
}


func RunEvents(configs []config.EventConfig) {
    
    for _, ev := range configs { 
        wg.Add(1)
        go RunEventForever(ev)
    }
}

func RunEventForever(conf config.EventConfig) {
    defer wg.Done()

    dur := time.Duration(conf.Interval) * time.Second

    //TODO: Log here?

    // I would be interested to find out how the ticker behaves,
    // like when a command takes a long time, does that mean the 
    // next command fires immediately? 
    ticker := time.Tick(dur)


    mainloop:
        for { 
         
            select {
                case <-quitChan:
                    log.Println("Recieved quit")
                    break mainloop
                

                case <-ticker:
                    out, err := exec.Command("sh", "-c", conf.Cmd).Output()
                    if err != nil {
                        //TODO: log error
                        continue
                    }
                    
                    ev := rproto.JsonToEvent(out)

                    //TODO: should i wrap this in select and have it non-blocking?
                    EventChan <- ev

                
            }
        }
    log.Println("Exiting event runner")

}


func RunEventOnce(conf config.EventConfig) (error){
    
    out, err := exec.Command("sh", "-c", conf.Cmd).Output()
    if err != nil {
        //TODO: best error, it's 3:30 am IDK?
        panic(err)
    }

    fmt.Println("Output from command: ", string(out))

    ev := rproto.JsonToEvent(out)

    EventChan <- ev

    return nil
}
