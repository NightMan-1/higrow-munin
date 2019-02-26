package main

import (
    "github.com/kataras/iris"
    "strconv"
)


type sensorStruct struct {
        DeviceId              string `json:"DeviceId"`
        Temperature           int `json:Temperature`
        Humidity              int `json:Humidity`
        Water                 int `json:Water`
        Light                 int `json:Light`
}

var sensorsAll = make(map[string]sensorStruct)

func WriteData(ctx iris.Context) {
    var sensorData sensorStruct

    if err := ctx.ReadJSON(&sensorData); err != nil {
        ctx.StatusCode(iris.StatusBadRequest)
        ctx.WriteString(err.Error())
        return
    }

    if (sensorData.DeviceId != ""){
        sensorsAll[sensorData.DeviceId] = sensorData

        ctx.Writef("Received: %s\n", sensorData)
    }else{
        ctx.WriteString("data error")
    }
}


func GetConfig(ctx iris.Context) {

    ctx.WriteString("multigraph higrow_temp\n")
    ctx.WriteString("graph_title HiGrow Temperature\n")
    ctx.WriteString("graph_vlabel *C\n")
    ctx.WriteString("graph_category HiGrow\n")
    for _, data := range sensorsAll {
        ctx.WriteString(data.DeviceId)
        ctx.WriteString(".label ")
        ctx.WriteString(data.DeviceId)
    }
    ctx.WriteString("\n\n")

    ctx.WriteString("multigraph higrow_hum\n")
    ctx.WriteString("graph_title HiGrow Humidity\n")
    ctx.WriteString("graph_vlabel %\n")
    ctx.WriteString("graph_category HiGrow\n")
    for _, data := range sensorsAll {
        ctx.WriteString(data.DeviceId)
        ctx.WriteString(".label ")
        ctx.WriteString(data.DeviceId)
    }
    ctx.WriteString("\n\n")

    ctx.WriteString("multigraph higrow_water\n")
    ctx.WriteString("graph_title HiGrow Water level\n")
    ctx.WriteString("graph_vlabel level\n")
    ctx.WriteString("graph_category HiGrow\n")
    for _, data := range sensorsAll {
        ctx.WriteString(data.DeviceId)
        ctx.WriteString(".label ")
        ctx.WriteString(data.DeviceId)
    }
    ctx.WriteString("\n\n")

    ctx.WriteString("multigraph higrow_light\n")
    ctx.WriteString("graph_title HiGrow Light level\n")
    ctx.WriteString("graph_vlabel level\n")
    ctx.WriteString("graph_category HiGrow\n")
    for _, data := range sensorsAll {
        ctx.WriteString(data.DeviceId)
        ctx.WriteString(".label ")
        ctx.WriteString(data.DeviceId)
    }
    ctx.WriteString("\n\n")
}


func GetData(ctx iris.Context) {

    ctx.WriteString("multigraph higrow_temp\n")
    for _, data := range sensorsAll {
        ctx.WriteString(data.DeviceId)
        ctx.WriteString(".value ")
        ctx.WriteString(strconv.Itoa(data.Temperature))
    }
    ctx.WriteString("\n\n")

    ctx.WriteString("multigraph higrow_hum\n")
    for _, data := range sensorsAll {
        ctx.WriteString(data.DeviceId)
        ctx.WriteString(".value ")
        ctx.WriteString(strconv.Itoa(data.Humidity))
    }
    ctx.WriteString("\n\n")

    ctx.WriteString("multigraph higrow_water\n")
    for _, data := range sensorsAll {
        ctx.WriteString(data.DeviceId)
        ctx.WriteString(".value ")
        ctx.WriteString(strconv.Itoa(data.Water))
    }
    ctx.WriteString("\n\n")

    ctx.WriteString("multigraph higrow_light\n")
    for _, data := range sensorsAll {
        ctx.WriteString(data.DeviceId)
        ctx.WriteString(".value ")
        ctx.WriteString(strconv.Itoa(data.Light))
    }
    ctx.WriteString("\n\n")
}



func main() {
    app := iris.New()

    app.Post("/write", WriteData)
    app.Get("/get", GetData)
    app.Get("/config", GetConfig)

    app.Run(iris.Addr(":8083"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)
}