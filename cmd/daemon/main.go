package main

import (
	"fmt"
	"github.com/kardianos/service"
	"github.com/scjtqs2/socks5/socks"
	"log"
	"os"
	"os/signal"
	"time"
)

type program struct {
	Server socks.Server
}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	for {
		p.Server = socks.NewServer(socks.ServerCfg{
			ListenPort: 1080,
			UserName:   "",
			Password:   "",
			UDPTimout:  60,
			TCPTimeout: 60,
			UDPAddr:    "127.0.0.1:1080",
			LogLevel:   "error",
		})
		err := p.Server.Run()
		if err != nil {
			log.Fatalln(err)
		}
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		signle := <-c
		fmt.Println("quit,Got signal:", signle)
		os.Exit(1)
	}
}

func (p *program) Stop(s service.Service) error {
	if p.Server != nil {
		p.Server.Stop()
	}
	fmt.Println("stoping")
	<-time.After(time.Second * 2)
	return nil
}

/**
* MAIN函数，程序入口
 */

func main() {
	svcConfig := &service.Config{
		Name:        "socks_Proxy", //服务显示名称
		DisplayName: "socks代理服务",   //服务名称
		Description: "socks代理服务",   //服务描述
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			s.Install()
			log.Println("服务安装成功 \r\n")
			return
		}

		if os.Args[1] == "uninstall" {
			s.Uninstall()
			log.Println("服务卸载成功 \r\n")
			return
		}
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = s.Run()
	if err != nil {
		log.Printf("ERROR: %s \r\n", err.Error())
	}
}
