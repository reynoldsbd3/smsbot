package core

import (
    "encoding/json"
    "log"
    "os"
    "time"
    
    "github.com/reynoldsbd3/smsbot/message"
)

// Core represents the core configuration and behavior of smsbot. It coordinates
// all message sources and provides facilities for building a message pipeline.
type Core struct {
    
    // Twilio account SID
    TwilioSid string `json:"twilioSid"`
    
    // Twilio auth token
    TwilioToken string `json:"twilioToken"`
    
    // Twilio SMS-capable number
    TwilioNumber string `json:"twilioNumber"`
    
    // Numbers to send SMS's to
    Recipients []string `json:"recipients"`
    
    // List of sources from which to get messages
    Sources *message.CompositeSource `json:"sources"`
}

// LoadCore loads and returns a new Core using the given path to the
// configuration file
func LoadCore(path string) (c *Core, err error) {
    
    r, err := os.Open(path)
    if err != nil { return nil, err }
    
    dec := json.NewDecoder(r)
    
    c = &Core{}
    err = dec.Decode(c)
    if err != nil { c = nil }
    
    return c, err
}


// Run builds and starts the message pipeline that the core is configured with
// in a new goroutine.
func (c *Core) Run() {
    
    log.Print("Building mesage pipeline")
    m := message.Timer(c.Sources, func() {
        time.Sleep(3 * time.Second)
    })
    
    // Left in to test without actually dispatching messages
    // go func() {
    //     for msg := range m {
    //         log.Printf("Message received: %s - %s", msg.Text, msg.URL)
    //     }
    // }()
    
    go dispatchMessages(m, c)
}
