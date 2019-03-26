package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	spinner.New(spinner.CharSets[9], 100*time.Millisecond).Start()
	for {
	}
}

func signalHandler() {
	exitChan := make(chan int, 1)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan)

	go func() {
		for {
			s := <-sigChan
			switch s {
			case syscall.SIGINT:
				fmt.Println("SIGINT")
				exitChan <- 0
			case syscall.SIGTERM:
				fmt.Println("SIGTERM")
				exitChan <- 0
			case syscall.SIGKILL:
				fmt.Println("SIGKILL is uncaught!")
			case syscall.SIGQUIT:
				fmt.Println("SIGQUIT")
				exitChan <- 0
			case syscall.SIGABRT:
				fmt.Println("SIGABRT")
				exitChan <- 0
			case syscall.SIGSTOP:
				fmt.Println("SIGSTOP")
				exitChan <- 0
			case syscall.SIGTSTP:
				fmt.Println("SIGTSTP")
				exitChan <- 0
			}
		}
	}()
	exitCode := <-exitChan
	fmt.Println("exit status: ", exitCode)
}
