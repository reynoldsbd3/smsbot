package config

import (
    "encoding/json"
    "os"
)

// Config holds configuration data for smsbot. Can be easily loaded from a JSON
// file.
type Config struct {
    
    // Twilio account SID
    TwilioSid string `json:"twilioSid"`
    
    // Twilio auth token
    TwilioToken string `json:"twilioToken"`
    
    // Twilio SMS-capable number
    TwilioNumber string `json:"twilioNumber"`
    
    // Numbers to send SMS's to
    Recipients []string `json:"recipients"`
}

// NewConfig loads and returns a new Config using the given path to the
// configuration file
func NewConfig(path string) (c *Config, err error) {
    
    r, err := os.Open(path)
    if err != nil { return nil, err }
    
    dec := json.NewDecoder(r)
    
    c = &Config{}
    err = dec.Decode(c)
    if err != nil { c = nil }
    
    return c, err
}
