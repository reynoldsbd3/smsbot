package message


import (
    "encoding/json"
    "fmt"
    "math/rand"
)


// A CompositeSource is a collection of MessageSources from which a single
// source is picked.
type CompositeSource []Source


// GetMessage returns the result of GetMessage from a randomly selected message
// source in the CompositeSource's collection
func (cs *CompositeSource) GetMessage() (m *Message, err error) {
    
    return (*cs)[rand.Intn(len(*cs))].GetMessage()
}


// UnmarshalJSON instantiates the list of concrete message sources comprising a
// CompositeSource
func (cs *CompositeSource) UnmarshalJSON(data []byte) error {
    
    var sources []struct{
        Type string
        Params *json.RawMessage
    }
    
    err := json.Unmarshal(data, &sources)
    if err != nil { return err }
    
    for _, rs := range sources {
        
        var s Source
        
        switch rs.Type {
        case "quote":
            s = &QuoteSource{}
        case "static":
            s = &StaticSource{}
        default:
            return fmt.Errorf("invalid type: %s", rs.Type)
        }
        
        err = json.Unmarshal(*rs.Params, s)
        if err != nil { return err }
        
        *cs = append(*cs, s)
    }
    
    return nil
}
