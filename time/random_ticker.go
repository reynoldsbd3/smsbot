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
    Duration time.Duration
    
    // Minimum time after interval during which a tick can occur 
    Minimum time.Duration
    
    // Maximum time after interval during which a tick can occur
    Maximum time.Duration
}


// Gets the random offset to the internal ticker at which to actually send a
// tick
func (rt *RandomTicker) getSleepTime() time.Duration {
    
    t := rt.Minimum
    
    if rt.Minimum < rt.Maximum {
        t += time.Duration(rand.Int63n(int64(rt.Maximum - rt.Minimum)))
    }
    
    return t
}


// Sleep until the next time that is a multiple of rt.Duration (for example,
// sleep until the next hour or day)
func (rt *RandomTicker) sleepUntilNext() {
        
    // Sleep until the next occurence of the duration
    nextTime := time.Now().Round(rt.Duration)
    if time.Now().After(nextTime) {
        nextTime = nextTime.Add(rt.Duration)
    }
    time.Sleep(nextTime.Sub(time.Now()))
}


// Start begins a goroutine that will send the time on the RandomTicker's
// channel at the configured interval
func (rt *RandomTicker) Start() {
    
    rt.C = make(chan time.Time, 1)
    
    go func() {
        
        // Start at a predictable time
        rt.sleepUntilNext()
        
        // Start with a tick so we don't have to wait one full interval before
        // things start to happen
        go func() {
            time.Sleep(rt.getSleepTime())
            rt.C <- time.Now()
        }()
        
        for range time.NewTicker(rt.Duration).C {
            go func() {
                time.Sleep(rt.getSleepTime())
                rt.C <- time.Now()
            }()
        }
    }()
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
