package core

import (
    "encoding/json"
    "log"
    "os"
    
    "github.com/reynoldsbd3/smsbot/message"
    "github.com/reynoldsbd3/smsbot/time"
)

// Core represents the core configuration and behavior of smsbot. It coordinates
// all message sources and provides facilities for building a message pipeline.
type Core struct {
    
    // Sets server to debug mode, which prints what the server will do rather
    // than actually taking action and potentially consuming any API's
    Debug bool `json:"debug"`
    
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
    
    // Describes interval at which messages will be produced and dispatched
    Time *time.RandomTicker `json:"time"`
}


// GetMessage retrieves a message from the configured message sources
func (c *Core) GetMessage() (*message.Message, error) {
    
    if c.Debug {
        return &message.Message{
            Source: "debug system",
            Text: "debug message",
            URL: "example.com",
        }, nil
    }
    
    return c.Sources.GetMessage()
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
    m := make(chan *message.Message, 1)
    
    go func() {
        c.Time.Start()
        
        for t := range c.Time.C {
            
            log.Print("Got message after interval ", t)
            msg, err := c.GetMessage()
            if err != nil {
                log.Print(err)
            } else {
                m <- msg
            }
        }
    }()
    
    go c.dispatchMessages(m)
}
