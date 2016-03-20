package main

import (
    "log"
    "os"
    "os/signal"
    
    "github.com/reynoldsbd3/smsbot/core"
)


func main() {
    
    log.Println("Starting smsbot")
    
    log.Println("Loading configuration")
    c, err := core.LoadCore("config.json")
    if err != nil { log.Fatal(err) }
    
    c.Run()
    
    waitUntilInterrupted()
    
    log.Println("Exiting smsbot")
}


// Blocks goroutine until SIGINT (^C) is received
func waitUntilInterrupted() {
    
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    _ = <-c
}
