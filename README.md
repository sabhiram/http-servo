# slackbot-servo

Expose a RC servo over a slackbot inside a RaspberryPi GPIO + PWM

## Why?

At my current startup, we have an always on video portal, and I have heard "can you turn the camera left" a few too many times. This attempts to use a RaspberryPi3, a tower-pro 9g servo and a slackbot to automate this "problem" away.

## Quickstart

```
go get github.com/sabhiram/slackbot-servo
cd $GOPATH/src/$_
go build .

sudo SLACKBOT_TOKEN=Azalakjds.... slackbot-servo
```

## Commands

Currently the bot supports:

`turn left` : turns the servo "left" (closer to 0) by 18 degrees.
`turn right` : turns the servo "right" (closer to 180) by 18 degrees.
`center`: turns the servo to the center position (90 degrees).
`angle`: returns the current angle to the prompter
`help`: returns a somewhat helpful message to the prompter
