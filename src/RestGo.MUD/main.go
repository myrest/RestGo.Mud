package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"rest.com.tw/tinymud/src/RestGo.MUD.Core/Config"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core/Telnet"

	//註冊命令用
	_ "rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands/PPLCommand"

	//初始化物件
	_ "rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/Room/RoomHelper"
)

func main() {
	port, _ := strconv.Atoi(os.Getenv("MUDPORT"))
	if port < 1 {
		port = Config.ServiceConfig.ListenOnPortNumber
	}

	server, err := Telnet.Listen(port)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Start the server in a goroutine
	go server.Start()

	// Wait for a shutdown signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	server.Shutdown()
}
