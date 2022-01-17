package main

import (
	"fmt"
	"github.com/scjtqs2/socks5/socks"
	"log"
	"os"
	"os/signal"
)

// 支持socks4和socks5
func main() {
	config := socks.ServerCfg{
		ListenPort: 1080,
		UserName:   "", //校验用户名
		Password:   "", //校验密码
		UDPTimout:  60,
		TCPTimeout: 60,
		UDPAddr:    "127.0.0.1:1080",
		LogLevel:   "error",
	}
	s := socks.NewServer(config)
	err := s.Run()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("socks started! listen port:", config.ListenPort)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signle := <-c
	fmt.Println("quit,Got signal:", signle)
	s.Stop()
	os.Exit(1)
}
