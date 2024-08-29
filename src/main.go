package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/kardianos/service"
	"github.com/s-j-rollins-technologies/hyper-v-monitoring/routes"
	"golang.org/x/sys/windows"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Service panicked: %v", r)
		}
	}()

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
	}
	err = os.Chdir(dir)
	if err != nil {
		log.Fatalf("Failed to change directory: %v", err)
	}

	if err := routes.StartWebServer(); err != nil {
		log.Fatalf("Error: %v", err)
	}

	select {}
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func isInteractive() bool {
	stdinHandle := windows.Handle(os.Stdin.Fd())

	var mode uint32
	err := windows.GetConsoleMode(stdinHandle, &mode)
	return err == nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "HyperVMonitoringAPI",
		DisplayName: "Hyper-V Monitoring API",
		Description: "Service to monitor and expose Hyper-V Monitoring via an API endpoint. Made by Connor - Written in Go",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		action := os.Args[1]
		if action == "install" || action == "uninstall" || action == "start" || action == "stop" || action == "restart" {
			err = service.Control(s, action)
			if err != nil {
				log.Fatalf("Failed to execute action %q: %v", action, err)
			}
		} else {
			log.Fatalf("Invalid action: %q. Valid actions are: %q", action, service.ControlAction)
		}
		return
	}

	if isInteractive() {
		fmt.Println("Running interactively...")
		prg.run()
	} else {
		err = s.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
