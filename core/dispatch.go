package core

import (
	"fmt"
	"log"

	"github.com/sfreiberg/gotwilio"

	"github.com/reynoldsbd3/smsbot/message"
)


// Waits for messages on the channel and dispatches them to all recipients using
// the given config.Config
func (c *Core) dispatchMessages(messages chan *message.Message) {

	t := gotwilio.NewTwilioClient(c.TwilioSid, c.TwilioToken)

    log.Print("Starting message dispatch loop")
	for msg := range messages {

		log.Printf("Dispatching message from %s", msg.Source)
		for _, to := range c.Recipients {
            if c.Debug {
                log.Printf("--> %s", msg.Text)
            } else {
			    go sendTwilioMessage(t, c.TwilioNumber, to, msg)
            }
		}
	}
}


// Forms and sends a message using the provided Twilio client and handles any
// errors
func sendTwilioMessage(t *gotwilio.Twilio, from, to string, msg *message.Message) {

	log.Printf("Sending message to %s via Twilio", to)

	rawMessage := fmt.Sprintf("%s\n%s", msg.Text, msg.URL)
	_, _, err := t.SendSMS(from, to, rawMessage, "", "")

	if err != nil {
		log.Print(err)
	}
}
