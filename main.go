package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	osSignalCh := make(chan os.Signal, 1)
	signal.Notify(osSignalCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer signal.Stop(osSignalCh)

	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Hello, %dth World!\n", i)
			time.Sleep(time.Second * 1)
		}
		done <- true
	}()

	for {
		select {
		case killSignal := <-osSignalCh:
			switch killSignal {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				log.Println("killed")
				for i := 0; i < 10; i++ {
					fmt.Printf("Goodbye, %dth World!\n", i)
					time.Sleep(time.Second * 1)
				}
				return
			}
		case <-done:
			log.Println("done")
			return
		}
	}
}
