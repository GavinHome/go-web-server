package main

import (
	"build-web-application-with-golang/log/logs"
	"errors"
)

func main() {
	logs.Logger.Info("Start server at:%v", "192.168.0.10")
	logs.Logger.Critical("Server err:%v", errors.New("server error"))
}
