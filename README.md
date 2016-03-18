# Overview

smsbot is a simple application that periodically generates an SMS message from
one of several configured sources and sends that message to a list of recipients
using the Twilio API.

# Getting Started

1. (Install Go)[https://golang.org/doc/install] and configure $GOPATH

2. Insall smsbot:

  ```
  go get -u github.com/reynoldsbd3/smsbot
  ```

3. Copy config-example.json to the current working directory and rename it to
  config.json, replacing its contents with valid Twilio API credentials

4. Run the smsbot command.

# Limitations

Currently, smsbot will send one of 3 hard-coded messages every 30 seconds.

Stay tuned.
