package message

import (
    "math/rand"
)

// A CompositeSource is a collection of MessageSources from which a single
// source is picked.
type CompositeSource struct {
    
    // Internal collection of MessageSources
    Sources []Source
}

// GetMessage returns the result of GetMessage from a randomly selected message
// source in the CompositeSource's collection
func (cs *CompositeSource) GetMessage() (m *Message, err error) {
    
    return cs.Sources[rand.Intn(len(cs.Sources))].GetMessage()
}