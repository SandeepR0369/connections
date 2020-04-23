package main

import (
        "fmt"
        "slack"
)

var TAGME = "<@Sandeep Reddy>"
var SLACK_CHANNEL = "test-alerts"
func main (){
        slack.PostMessage(slack.Message{Channel: SLACK_CHANNEL, Message: fmt.Sprintf("TEST %s", TAGME)})
	//slack.PostMessage(slack.Message{Channel: SLACK_CHANNEL, Message: fmt.Sprintf("TEST @Sandeep Reddy")})
}
