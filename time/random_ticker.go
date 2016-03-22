package time


import (
    "encoding/json"
    "math/rand"
    "time"
)


// A RandomTicker holds a channel that delivers 'ticks' of a clock at random
// offsets to a regular interval
type RandomTicker struct {
    
    // The channel on which the ticks are delivered
    C chan time.Time `json:"-"`
    
    // Mean interval of ticks
    Duration time.Duration `json:"duration"`
    
    // Minimum time after interval during which a tick can occur 
    Minimum time.Duration `json:"minimum"`
    
    // Maximum time after interval during which a tick can occur
    Maximum time.Duration `json:"maximum"`
}


// UnmarshalJSON interprets human-readable time.Duration representations rather
// than raw nanosecond values (the default JSON encoding for time.Duration).
func (rt *RandomTicker) UnmarshalJSON(data []byte) error {
    
    var s struct{
        Duration string `json:"duration"`
        Minimum string `json:"minimum"`
        Maximum string `json:"maximum"`
    }
    
    err := json.Unmarshal(data, &s)
    if err != nil { return err }
    
    rt.Duration, err = time.ParseDuration(s.Duration)
    if err != nil { return err }
    
    rt.Minimum, err = time.ParseDuration(s.Minimum)
    if err != nil { return err }
    
    rt.Maximum, err = time.ParseDuration(s.Maximum)
    if err != nil { return err }
        
    return nil
}


// Start begins a goroutine that will send the time on the RandomTicker's
// channel at the configured interval
func (rt *RandomTicker) Start() {
    
    rt.C = make(chan time.Time, 1)
    
    go func() {
        
        rand.Seed(time.Now().UTC().UnixNano())
        
        // Sleep until the next occurence of the interval
        nextTime := time.Now().Round(rt.Duration)
        if time.Now().After(nextTime) {
            nextTime = nextTime.Add(rt.Duration)
        }
        time.Sleep(nextTime.Sub(time.Now()))
        
        for t := range time.NewTicker(rt.Duration).C {
            
            time.Sleep(rt.Minimum)
            t = t.Add(rt.Minimum)
            
            // Int63n must have arg > 0
            if rt.Maximum > rt.Minimum {
                r := time.Duration(rand.Int63n(int64(rt.Maximum - rt.Minimum)))
                time.Sleep(r)
                t = t.Add(r)
            }
            
            rt.C <- t
        }
    }()
}
