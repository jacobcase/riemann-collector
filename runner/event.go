
package runner

import (
    "riemannd/rproto"
    "riemannd/config"
    "os/exec"
    "fmt"
    "log"
)


func RunEvents(configs []config.EventConfig, eventChan chan *rproto.Event, doneChan chan bool) {
    
    for _, ev := range configs {        
        go RunEventForever(ev, eventChan, doneChan)
    }


}

func RunEventForever(conf config.EventConfig, eventChan chan *rproto.Event, done chan bool) (error){

    dur := time.Duration(conf.Interval * time.Second)

    //TODO: Log here?

    // I would be interested to find out how the ticker behaves,
    // like when a command takes a long time, does that mean the 
    // next command fires immediately? 
    ticker := time.Ticke(dur)

    for { 
     
        select {
            case stop = <-done:
                break

            case tick = <-ticker {
                out, err := exec.Command("sh", "-c", conf.Cmd).Output()
                if err != nil {
                    //TODO: log error
                    continue
                }
                
                ev := rproto.JsonToEvent(out)

                //TODO: should i wrap this in select and have it non-blocking?
                eventChan <- ev

            }
        }
    }

}


func RunEventOnce(eventChan chan *rproto.Event, conf config.EventConfig) (error){
    
    out, err := exec.Command("sh", "-c", conf.Cmd).Output()
    if err != nil {
        //TODO: best error, it's 3:30 am IDK?
        panic(err)
    }

    fmt.Println("Output from command: ", string(out))

    ev := rproto.JsonToEvent(out)

    eventChan <- ev

    return nil
}
