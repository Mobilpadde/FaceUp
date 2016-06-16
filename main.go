package main

import (
    "log"
    "time"
    "os"
    "encoding/json"

    "github.com/1lann/messenger"
    "github.com/Collinux/GoHue"
)

type Conf struct{
    UserId string
    Mail string
    Pass string
    Name string
    Speed time.Duration
    Hueser string
}

var (
    light hue.Light
    fbState hue.LightState
    cfg Conf
)

func decodeCfg(){
    file, errOpen := os.Open("./config.json")
    if errOpen != nil {
        log.Fatal(errOpen)
    }

    decoder := json.NewDecoder(file)

    errDecode := decoder.Decode(&cfg)
    if errDecode != nil {
        log.Fatal(errDecode)
    }
}

func setupLight(){
    bridges, _ := hue.FindBridges()
    bridge := bridges[0]

    //user, _ := bridge.CreateUser("wowser")
    bridge.Login(cfg.Hueser) // user

    l, _ := bridge.GetLightByName(cfg.Name)
    light = l

    fbState = hue.LightState{
        XY: hue.BLUE,
        Bri: 254,
    }
}


func init(){
    decodeCfg()
    setupLight()
}

func onMsg(msg *messenger.Message){
    if cfg.UserId != msg.FromUserID {
        oldState := hue.LightState{
            On: light.State.On,
            Bri: light.State.Bri,
            Hue: uint16(light.State.Hue),
        }

        light.SetState(fbState)
        
        light.On()
        time.Sleep(cfg.Speed * time.Millisecond)
        light.Off()

        light.SetState(oldState)
    }
}

func main(){
    session := messenger.NewSession();

    err := session.Login(cfg.Mail, cfg.Pass);

    if err != nil {
        log.Fatal(err)
    } else {
       log.Println("Yo, we're waiting for messages!")
    }

    session.OnMessage(onMsg)
    session.Listen()
}