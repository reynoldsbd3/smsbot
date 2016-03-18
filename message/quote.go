package message

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)


// The result of a query to the They Said So API
type result struct {
    
    // Indicates that the query was successful 
    Success struct{
        
        // Number of quotes returned
        Total int
    }
    
    Contents struct{
        
        Quotes []struct{
            
            Quote string
            
            Length string
            
            Author string
            
            Tags []string
            
            Category string
            
            Date string
            
            Title string
            
            Background string
            
            ID string
        }
    }
}


// A QuoteSource produces the quote of the day from theysaidso.com under the
// given calendar
type QuoteSource struct {
    
    // The category that will be used to fetch the quote of the day
    category string
}


// NewQuoteSource returns a new QuoteSource configured to get the quote of the
// day from the given category
func NewQuoteSource(category string) *QuoteSource {
    
    return &QuoteSource{
        category,
    }
}


// GetMessage uses the They Said So API to retrieve the quote of the day for a
// given category
func (qs *QuoteSource) GetMessage() (*Message, error) {
    
    url := fmt.Sprintf("http://quotes.rest/qod.json?category=%s", qs.category)
    
    resp, err := http.Get(url)
    if err != nil {
        log.Print(err)
        return nil, err
    }
    
    defer resp.Body.Close()
    dec := json.NewDecoder(resp.Body)
    
    r := &result{}
    err = dec.Decode(r)
    if err != nil {
        log.Print(err)
        return nil, err
    }
    
    return &Message{
        fmt.Sprintf("They Said So Quote of the Day for %s", qs.category),
        fmt.Sprintf("%s - %s", r.Contents.Quotes[0].Quote, r.Contents.Quotes[0].Author),
        fmt.Sprintf("https://theysaidso.com/qod?category=%s", qs.category),
    }, nil
}
