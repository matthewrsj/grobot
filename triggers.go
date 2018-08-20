package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"

	"github.com/matthewrsj/grobot/ghub"
	"github.com/whyrusleeping/hellabot"
)

func randomString(ss []string) string {
	return ss[rand.Intn(len(ss)-1)]
}

func randSleep() {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
}

func ircReply(irc *hbot.Bot, m *hbot.Message, msg string) {
	randSleep()
	irc.Reply(m, msg)
}

// This trigger replies Hello when you say hello
var SayInfoMessage = hbot.Trigger{
	func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && m.Content == "!info"
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		ircReply(irc, m, "!info !shrug !finger !love !yeah !{hello|hi} !foo !overtime !8ball")
		ircReply(irc, m, "!mock !{goodbye|later|bye} !numissues")
		return false
	},
}

var ShrugTrigger = hbot.Trigger{
	func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, "!shrug")
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		ircReply(irc, m, `¯\_(ツ)_/¯ -`+m.From)
		return false
	},
}

var FingerTrigger = hbot.Trigger{
	func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, "!finger")
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		ircReply(irc, m, `╭∩╮(Ο_Ο)╭∩╮ -`+m.From)
		return false
	},
}

var LoveTrigger = hbot.Trigger{
	func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, "!love")
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		ircReply(irc, m, `(♥_♥) -`+m.From)
		return false
	},
}

func isYeah(m string) bool {
	m = strings.ToLower(m)
	return strings.HasPrefix(m, *nick+": yeah") ||
		strings.HasPrefix(m, *nick+": you rock") ||
		strings.HasPrefix(m, "!yeah")
}

var YeahTrigger = hbot.Trigger{
	func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && isYeah(m.Content)
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		ircReply(irc, m, "( •_•)")
		ircReply(irc, m, "( •_•)>⌐■-■")
		ircReply(irc, m, "(⌐■_■)")
		return false
	},
}

var hellos = []string{
	"hey %s!",
	"hey there %s!",
	"yesssss %s?",
	"watchu want %s?",
	"how *you* doin %s?",
	"hihihihihihihihihihihihihihi %s!",
	"mmmmmmmmmm %s",
	"huh? what %s?",
	"*yawn* yo whassup %s?",
	"enough flimflam, get to the point %s",
	"no, I wasn't talking about you behind your back %s, oh I mean hello...",
}

var gabiHellos = []string{
	"...",
	"... yeah?",
	"no",
	"yes gabi wabi fliby flaby?",
	"wtf why you keep talkin' to me?",
	"was I talking to you?",
	"anyways, what was everyone *else* talking about?",
	"*yaaaaaaaaawn*",
	"just kidding yo, you ROCK!",
	"YAAAAS",
}

var tudorHellos = []string{
	"hi turdor lolololol",
	"Marcuaeeueuu... is that French?",
	"torta?",
	"dude chill",
	"mmmmmmmm... gym",
	"dude, you're lookin' small",
	"hi",
	"stop poking me",
}

var mrsjHellos = []string{
	"hey there big guy",
	";)",
	"wassup wassup",
	"your wish is my command",
}

func isHello(message string) bool {
	message = strings.ToLower(message)
	return strings.HasPrefix(message, "!hello") ||
		strings.HasPrefix(message, *nick+": hello") ||
		strings.HasPrefix(message, *nick+": hi") ||
		strings.HasPrefix(message, "!hi")
}

var HelloTrigger = hbot.Trigger{
	func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && isHello(m.Content)
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		switch m.From {
		case "mrsj":
			ircReply(irc, m, randomString(mrsjHellos))
		case "gnbeyer":
			ircReply(irc, m, randomString(gabiHellos))
		case "tmarcu":
			ircReply(irc, m, randomString(tudorHellos))
		default:
			ircReply(irc, m, fmt.Sprintf(randomString(hellos), m.From))
		}
		return false
	},
}

var MeTrigger = hbot.Trigger{
	func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" &&
			strings.HasPrefix(m.Content, *nick+":") &&
			!isHello(m.Content) &&
			!isBye(m.Content) &&
			!isYeah(m.Content)
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		ircReply(irc, m, "I like it when you call me "+*nick+" ;)")
		return false
	},
}

var FooTrigger = hbot.Trigger{
	func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, "!foo")
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		ircReply(irc, m, "bar")
		return false
	},
}

var OvertimeTrigger = hbot.Trigger{
	func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, "!overtime")
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		ircReply(irc, m, fmt.Sprintf("You get $%d of overtime pay!", rand.Intn(1000)))
		return false
	},
}

var eightBallResps = []string{
	// yes
	"It is certain.",
	"It is decidedly so.",
	"Withoug a doubt.",
	"Yes - definitely.",
	"You may rely on it.",
	"As I see it, yes.",
	"Most likely.",
	"Outlook good.",
	"Yes.",
	"Signs point to yes.",
	// maybe
	"Reply hazy, try again.",
	"Ask again later.",
	"Better not tell you now.",
	"Cannot predict now.",
	"Concentrate and ask again.",
	// no
	"Don't count on it.",
	"My reply is no.",
	"My sources say no.",
	"Outlook is not good.",
	"Very doubtful.",
}

var EightBallTrigger = hbot.Trigger{
	func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, "!8ball")
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		ircReply(irc, m, randomString(eightBallResps))
		return false
	},
}

func replaceAtIdx(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

var MockTrigger = hbot.Trigger{
	func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, "!mock")
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		reply := strings.TrimPrefix(m.Content, "!mock")
		reply = strings.ToLower(strings.TrimSpace(reply))
		for i, r := range reply {
			if rand.Intn(1000)%2 == 0 {
				reply = replaceAtIdx(reply, unicode.ToUpper(r), i)
			}
		}
		ircReply(irc, m, reply)
		return false
	},
}

var byes = []string{
	"byeeee",
	"cya",
	"okokokokokbye",
	"getcho ass outta here",
	"Goodbye",
	"latazzz",
	"latezzz",
	"laterrr",
	"in a while crocodile",
	"cya later alligator",
	"ok bye",
	"y go tho?",
	"but we were having so much fun... :(",
	"whatevz",
}

func isBye(msg string) bool {
	msg = strings.ToLower(msg)
	return strings.HasPrefix(msg, *nick+": bye") ||
		strings.HasPrefix(msg, *nick+": goodbye") ||
		strings.HasPrefix(msg, "!goodbye") ||
		strings.HasPrefix(msg, "!later") ||
		strings.HasPrefix(msg, "!bye")
}

var ByeTrigger = hbot.Trigger{
	func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && isBye(m.Content)
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		ircReply(irc, m, randomString(byes))
		return false
	},
}

var NumIssuesTrigger = hbot.Trigger{
	func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, "!numissues ")
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		msg := strings.TrimPrefix(m.Content, "!numissues ")
		msg = strings.TrimSpace(msg)
		open, closed, err := ghub.GetNumOpenClosedIssues("clearlinux", msg)
		if err != nil {
			ircReply(irc, m, "sorry! I got an error: "+err.Error())
		} else {
			ircReply(irc, m, fmt.Sprintf("that repo has %d open issues and %d closed issues", open, closed))
		}
		return false
	},
}
