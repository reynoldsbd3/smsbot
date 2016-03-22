package message

// A StaticSource produces the same message each time GetMessage is called.
type StaticSource struct {
    
    Message string `json:"message"`
    
    URL string `json:"url"`
}

// GetMessage returns the StaticSource's message
func (ss *StaticSource) GetMessage() (*Message, error) {
    
    return &Message{
        "static message source",
        ss.Message,
        ss.URL,
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
