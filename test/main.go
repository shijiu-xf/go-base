package main

import (
	"fmt"
	"github.com/shijiu-xf/go-base/comm/zapsj"
	"github.com/shijiu-xf/go-base/config"
)

func main() {

	var c = config.LogFileConfig{
		Level:     "debug",
		FileName:  "shijiu",
		MaxSize:   500,
		LocalTime: true,
		Compress:  true,
	}

	err := zapsj.InitSJZap(c)
	if err != nil {
		fmt.Println("失败")
		panic(err)
	}
	zapsj.ZapL().Info("ererw")

}
