package main

import (
    "log"
    "os"
    "os/signal"
    "time"
    
    "github.com/reynoldsbd3/smsbot/config"
    "github.com/reynoldsbd3/smsbot/message"
)


func main() {
    
    log.Println("Starting smsbot")
    
    log.Println("Loading configuration")
    c, err := config.NewConfig("config.json")
    if err != nil { log.Fatal(err) }
    
    log.Print("Building message pipeline")
    messages := message.Timer(c.Sources, func() {
        time.Sleep(30 * time.Second)
    })
    
    // Left in to test without actually dispatching messages
    // go func() {
    //     for msg := range messages {
    //         log.Printf("Message received: %s - %s", msg.Text, msg.URL)
    //     }
    // }()
    
    go dispatchMessages(messages, c)
    
    waitUntilInterrupted()
    
    log.Println("Exiting smsbot")
}


// Blocks goroutine until SIGINT (^C) is received
func waitUntilInterrupted() {
    
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    _ = <-c
}
