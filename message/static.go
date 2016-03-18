package message

// A StaticSource produces the same message each time GetMessage is called.
type StaticSource struct {
    
    message string
    
    url string
}

// GetMessage returns the StaticSource's message
func (ss *StaticSource) GetMessage() (*Message, error) {
    
    return &Message{
        "static message source",
        ss.message,
        ss.url,
    }, nil
}

// NewStaticSource returns a new instance of StaticSource with the given static
// message contents
func NewStaticSource(message string, url string) *StaticSource {
    
    return &StaticSource {
        message,
        url,
    }
}
