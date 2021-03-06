package main

import (
    "time"
    "github.com/gogf/gf/g/os/glog"
    "github.com/gogf/gf/g/os/gtime"
)

func main() {
    gtime.SetTimeout(3*time.Second, func() {
        glog.SetDebug(false)
    })
    for {
        glog.Debug(gtime.Datetime())
        time.Sleep(time.Second)
    }
}


