package message

import (
    "log"
)


// A Message encapsulates a short string and some metadata that can be
// dispatched by smsbot.
type Message struct {
    
    // Identifies the source of this message
    Source string
    
    // The actual contents of the message
    Text string
    
    // A relevant URL that can be followed by the user to learn more about the
    // message
    URL string
}


// Source defines the interface of a message source, which can produce a
// message on demand.
type Source interface {
    
    // Gets an instance of Message from this source
    GetMessage() (*Message, error)
}


// Timer transforms a Source into a channel that produces messages at some
// interval determined by the provided interval function, which is called
// between each insertion to create a delay.
func Timer(source Source, interval func()) chan *Message {
    
    log.Print("Creating timed message channel")
    
    c := make(chan *Message, 1)
    
    go func() {
        for {
            msg, err := source.GetMessage()
            
            if err != nil {
                log.Print(err)
            } else {
                c <- msg
            }
            
            interval()
            log.Print("Timer elapsed, creating message")
        }
    }()
    
    return c
}
