// This is an example program showing the usage of hellabot
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/whyrusleeping/hellabot"
	log "gopkg.in/inconshreveable/log15.v2"
)

var serv = flag.String("server", "irc.freenode.net:6667", "hostname and port for irc server to connect to")
var nick = flag.String("nick", "grobot", "nickname for the bot")
var channel = flag.String("chan", "#test", "channel to join")

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	flag.Parse()

	hijackSession := func(bot *hbot.Bot) {
		bot.HijackSession = false
	}
	sslOptions := func(bot *hbot.Bot) {
		bot.SSL = true
	}
	channels := func(bot *hbot.Bot) {
		bot.Channels = []string{*channel}
	}
	irc, err := hbot.NewBot(*serv, *nick, hijackSession, channels, sslOptions)
	if err != nil {
		panic(err)
	}

	initScoresFromFile()

	irc.AddTrigger(SayInfoMessage)
	irc.AddTrigger(ShrugTrigger)
	irc.AddTrigger(FingerTrigger)
	irc.AddTrigger(LoveTrigger)
	irc.AddTrigger(HelloTrigger)
	irc.AddTrigger(MeTrigger)
	irc.AddTrigger(FooTrigger)
	irc.AddTrigger(OvertimeTrigger)
	irc.AddTrigger(EightBallTrigger)
	irc.AddTrigger(MockTrigger)
	irc.AddTrigger(ByeTrigger)
	irc.AddTrigger(NumIssuesTrigger)
	irc.AddTrigger(YeahTrigger)
	irc.AddTrigger(SuckTrigger)
	irc.AddTrigger(UpdateScoreTrigger)
	irc.AddTrigger(ScoreTrigger)
	irc.AddTrigger(ScoreOverviewTrigger)
	irc.AddTrigger(ScoreAllTrigger)
	irc.AddTrigger(ScoreTopTrigger)
	irc.Logger.SetHandler(log.StdoutHandler)

	// Start up bot (this blocks until we disconnect)
	irc.Run()
	fmt.Println("Bot shutting down.")
}
