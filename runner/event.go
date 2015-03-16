
package runner

import (
    "riemannd/rproto"
    "riemannd/config"
    "os/exec"
    "fmt"
)



//func RunEventForever(chan *rproto.Event, config.EventConfig (error){
//}


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
