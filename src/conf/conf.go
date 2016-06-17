package conf

import (
    "bufio"
    "fmt"
    "os"
    "time"
    "encoding/json"

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

func checkErr(err error){
    if err != nil {
        panic(err)
    }
}

func Generate(){
    cfg := Conf{}
    var fId, mail, pass, name, hueser string
    var speed time.Duration

    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Your facebook id:")
    fmt.Scanln(&fId)
    cfg.UserId = fId

    fmt.Println("Your facebook mail:")
    fmt.Scanln(&mail)
    cfg.Mail = mail

    fmt.Println("Your facebook password:")
    pass, err := reader.ReadString('\n')
    checkErr(err)
    cfg.Pass = pass[:len(pass) - 2]

    bridges, err := hue.FindBridges()
    checkErr(err)
    bridge := bridges[0]

    fmt.Println("Please authorize this program, by linking it on your bridge now; you've got 15 seconds.")
    time.Sleep(15000 * time.Millisecond)
    hueser, err = bridge.CreateUser(mail)
    checkErr(err)
    cfg.Hueser = hueser

    fmt.Println("Name of your light:")
    name, err = reader.ReadString('\n')
    checkErr(err)
    cfg.Name = name[:len(name) - 2]

    fmt.Println("Blink speed of the light:")
    fmt.Scanln(&speed)
    cfg.Speed = speed

    byteArr, _ := json.Marshal(cfg)
    file, err := os.Create("config.json")
    checkErr(err)

    _, err =file.Write(byteArr)
    checkErr(err)
}

