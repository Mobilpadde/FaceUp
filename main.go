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

func checkErr(err error){
    if err != nil {
        log.Fatal(err)
    }
}

func decodeCfg(){
    file, err := os.Open("./config.json")
    checkErr(err)

    decoder := json.NewDecoder(file)

    err = decoder.Decode(&cfg)
    checkErr(err)
}

func setupLight(){
    bridges, err := hue.FindBridges()
    checkErr(err)

    bridge := bridges[0]

    //user, _ := bridge.CreateUser("wowser")
    bridge.Login(cfg.Hueser) // user

    l, err := bridge.GetLightByName(cfg.Name)
    checkErr(err)

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
    if cfg.UserId == msg.FromUserID {
        return
    }

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

func main(){
    session := messenger.NewSession();

    err := session.Login(cfg.Mail, cfg.Pass);

    checkErr(err)

    log.Println("Yo, we're waiting for messages!")

    session.OnMessage(onMsg)
    session.Listen()
}