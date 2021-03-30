package main

import (
	"fmt"
	"os/exec"
	"io/ioutil"

	"github.com/mitchellh/go-ps"
	"github.com/getlantern/systray"
)

const (
	application string = ""
	directory string = ""
)

var (
	launched bool = false
	running bool = true
)

func FindProcessByName(name string) (ps.Process) {
	processes, _ := ps.Processes()

	for _, process := range(processes) {
		if process.Executable() == name {
			return process
		}
	}

	return nil
}

func AwaitProcessOpened(name string) (ps.Process) {
	process := FindProcessByName(name)

	for (process == nil) {
		process = FindProcessByName(name)
	}

	return process
}

func AwaitProcessClosed(name string) {
	process := FindProcessByName(name)

	for (process != nil) {
		process = FindProcessByName(name)
	}

	return
}

func LookForLaunch() {
	if launched {
		return
	}

	process := AwaitProcessOpened(application)

	launched = true
	fmt.Println("process started")

	cmd := exec.Command(directory)
	cmd.Run()

	AwaitProcessClosed(process.Executable())
	fmt.Println("process ended")
	launched = false
}

func OnReady() {
	fmt.Println("started")

	icon, _ := ioutil.ReadFile("Icon.ico")

	systray.SetIcon(icon)
	systray.SetTitle("Title")
	systray.SetTooltip("Tip")

	close := systray.AddMenuItem("Quit", "Close the application")

	go func() {
		<-close.ClickedCh

		fmt.Println("ended")
		systray.Quit()
	}()

	go func() {
		for (running) {
			LookForLaunch()
		}
	}()
}

func OnExit() {
	running = false
}

func main() {
	systray.Run(OnReady, OnExit)
}
