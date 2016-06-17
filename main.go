package main

import (
    "flag"
    "log"
    "time"
    "os"
    "encoding/json"

    "conf"

    "github.com/1lann/messenger"
    "github.com/Collinux/GoHue"
)

var (
    light hue.Light
    fbState hue.LightState
    cfg conf.Conf
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

    bridge.Login(cfg.Hueser)

    l, err := bridge.GetLightByName(cfg.Name)
    checkErr(err)

    light = l

    fbState = hue.LightState{
        XY: hue.BLUE,
        Bri: 254,
    }
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

func listen(){
    decodeCfg()
    setupLight()

    session := messenger.NewSession();

    err := session.Login(cfg.Mail, cfg.Pass);

    checkErr(err)

    log.Println("Yo, we're waiting for messages!")

    session.OnMessage(onMsg)
    session.Listen()
}

func main(){
    generate := flag.Bool("g", false, "Need help generating a config?")
    flag.Parse()

    if *generate {
        conf.Generate()
    }

    listen()
}