package main

import (
    //"fmt"

    "gobot.io/x/gobot"
    "gobot.io/x/gobot/platforms/leap"

    // デバッグ用
    "github.com/davecgh/go-spew/spew"
)

func main() {
    leapMotionAdaptor := leap.NewAdaptor("127.0.0.1:6437")
    l := leap.NewDriver(leapMotionAdaptor)
    work := func() {
        l.On(l.Event("message"), func(data interface{}) {
            spew.Dump(data.(leap.Frame).Pointables)
        })
    }
    robot := gobot.NewRobot("leapBot",
        []gobot.Connection{leapMotionAdaptor},
        []gobot.Device{l},
        work,
    )
    robot.Start()
}
