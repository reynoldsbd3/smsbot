# Overview

smsbot is a simple application that periodically generates a message from one of
several configured sources and sends that message to a list of recipients using
the Twilio API.

Messages are sent at a random time within regular time windows configured by the
user. For example, smsbot can be configured to send a message between 9am and
5pm every day.

For more details on smsbot's features, see the Configuration section below. 

# Getting Started

1. [Install Go](https://golang.org/doc/install) and configure `$GOPATH`

2. Insall smsbot:

  ```
  go get -u github.com/reynoldsbd3/smsbot
  ```

3. Copy `config-example.json` to the current working directory and rename it to
  `config.json`, replacing its contents with valid
  [Twilio API](https://www.twilio.com/try-twilio) credentials
  
  > Currently, there is no option to specify the path of `config.json`. It must
  > be placed in the current working directory.

4. Run the `smsbot` command

# Features and configuration

smsbot really only does two things: it generates messages and dispatches those
messages. Each message is retrieved from a message source that is randomly
selected from a pre-configured list of sources. Once a message is generated, it
is sent to each configured recipient using the Twilio SMS API.

This process is repeated indefinitely at random times according to smsbot's
timing parameters. For example, the configuration might specify sending messages
daily at a random time between noon and 9 P.M.

## SMS Messages

In order to use smsbot, you must have Twilio API credentials. Access to the SMS
API is free (with strings attached).
[Click here](https://www.twilio.com/try-twilio) to learn more and sign up for
access.

smsbot needs to know your Twilio Account SID and Auth Token as well as the phone
number assigned to you by Twilio. These are configured like so in `config.json`:

```json
{
    ...
    "twilioSid": "XXXXXXXXXXXXXXXX",
    "twilioToken": "XXXXXXXXXXXXXXXX",
    "twilioNumber": "+15551234567",
    ...
}
```

Each time a message is generated, it is sent to each number configured in the
recipients list. Note that if you want only one recipient, the number must
still appear in the JSON list within the `[...]` brackets.

```json
{
    ...
    "recipients": [
        "+15551112222",
        "+15553334444"
    ],
    ...
}
```

## Message Sources

smsbot allows you to configure several message sources. When it's time to
generate a message, a random source is selected and used to generate a message.

Message source configuration looks like this:

```json
{
    ...
    "sources": [
        {
            "type": "SOURCE_TYPE",
            "params": {
                ...
            }
        },
        {
            "type": "SOURCE_TYPE",
            "params": {
                ...
            }
        }
    ]
}
```

There are several types of message source, and each will have a different set of
parameters that may be configured.

### Static Message Source

A static message source is a simple, static message. Each time the source is
selected, it will generate the same message.

```json
{
    ...
    "sources": [
        ...
        {
            "type": "static",
            "params": {
                
                // Actual text of the static message
                "message": "This is the content of the static message",
                
                // Optional URL providing message context
                "url": "http://example.com/"
            }
        },
        ...
    ]
}
```

### Quote Source

A quote source is a message source that uses the
[They Said So API](https://theysaidso.com/api/) to retrieve the quote of the day
in the specified category.

```json
{
    ...
    "sources": [
        ...
        {
            "type": "quote",
            "params": {
                "category": "art"
            }
        },
        ...
    ]
}
```

## Debug

To keep from using the various API's involved, there is a `debug` option that
causes smsbot to simply print out what it would actually be doing.

```json
{
    ...
    "debug": true,
    ...
}
```
