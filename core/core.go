package core

import (
    "encoding/json"
    "os"
    
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
