package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"
	"time"
)

type User struct {
	Nick string
	User string
	Host string
}

type Message struct {
	User    User
	Channel string
	Content string
	Action  bool
}

type BotClient struct {
	SockReader *bufio.Reader
	SockWriter *bufio.Writer
	ReadChan   chan string
	WriteChan  chan string
	debug      bool
	nick string
	host string
	channels []string
}

func (b *BotClient) Usage() {
	fmt.Fprintf(os.Stderr, "\033[1mUSAGE\033[0m: go run %s [server[:port]] [nick] [channel [channel...]]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func (b *BotClient) LoadConfig(host, nick string, channels []string) {
	// Add the nick channel if it's not in there already
	// IMO too much work for this one thing
/*	
	found := false
	for i := range channels {
		if channels[i] == nick {
			found = true
			break
		}
	}

	if found == false {
		old, channels := channels, make([]string, len(channels) + 1)
		for i := range old {
			channels[i] = old[i]
		}

		channels[len(channels) - 1] = nick
	}
*/
	if host == "" || nick == "" {
		b.Usage()
	}

	b.nick = nick
	b.host = host
	b.channels = channels
	fmt.Printf("nick: %v, host: %v, channels: %v\n", b.nick, b.host, b.channels)
}

func (b *BotClient) Connect() {
	// Connect
	b.WriteChan <- "NICK " + b.nick
	b.WriteChan <- "USER " + b.nick + " * * :wjx-test-bot"
	for c := range b.channels {
		b.WriteChan <- "JOIN " + b.channels[c]
	}

	// To make foonetic admins happy
	b.WriteChan <- "MODE +Bix " + b.nick
}

func (b *BotClient) Run() {

	// Regexes
	// Add the matcher here
	privmsgRegexp := regexp.MustCompile("^:(.+?)!(.+?)@(.+?)\\sPRIVMSG\\s(.+?)\\s:(.+)$")
	modeRegexp := regexp.MustCompile("^:(.+?)!(.+?)@(.+?)\\sMODE\\s(.+?)\\s(.+)$")

	for {
		str := <-b.ReadChan
		if str[0:6] == "PING :" {
			if b.debug {
				fmt.Printf("\033[34mSERVER PING\033[0m\n")
			}
			b.WriteChan <- "PONG :" + str[7:len(str)-1]
		}

		data := privmsgRegexp.FindAllStringSubmatch(str, 1)
		if data != nil {
			b.process_privmsg(data);
		}

		data = modeRegexp.FindAllStringSubmatch(str, 1)
		if data != nil {
			b.process_modemsg(data);
		}

		// and add the handler function here
	}
}

func (b *BotClient) startReader() {
	for {
		str, err := b.SockReader.ReadString(byte('\n'))
		if err != nil {
			fmt.Fprintf(os.Stderr, "\033[31;1mERROR\033[0m: %s\n", err)
			break
		}

		str = str[0 : len(str) - 2]

		if b.debug {
			fmt.Printf("<- \033[34;1m%s\033[0m\n", str)
		}
		b.ReadChan <- str
	}
}

func (b *BotClient) startWriter() {
	for {
		str := <-b.WriteChan

		if b.debug {
			fmt.Printf("-> \033[35;1m%s\033[0m\n", str)
		}

		_, err := b.SockWriter.WriteString(str + "\r\n")
		if err != nil {
			fmt.Fprintf(os.Stderr, "\033[31;1mERROR\033[0m: %s\n", err)
			break
		}

		b.SockWriter.Flush()
	}
}

func (b *BotClient) Init(debug bool) {
	// Connect to the server
	addr, err := net.ResolveTCPAddr("tcp", b.host)
	if err != nil {
		b.warn("Unable to resolve TCP Address " + b.host)
	}

	// Dial the socket
	socket, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		b.warn("Unable to dial socket")
	}

	b.SockReader = bufio.NewReader(socket)
	b.SockWriter = bufio.NewWriter(socket)
	b.ReadChan   = make(chan string, 1000)
	b.WriteChan  = make(chan string, 1000)
	b.debug      = debug

	// start goroutine to do this
	go b.startReader()
	go b.startWriter()
}

func (b *BotClient) privmsg(str string, channel string) {
	b.WriteChan <- "PRIVMSG " + channel + " :" + str
	if channel != b.nick {
		fmt.Printf("[\033[34m%s\033[0m] \033[32;4;1m%s\033[0m: %s\n", channel, b.nick, str)
	}
}

func (b *BotClient) notice(str string, channel string) {
	b.WriteChan <- "NOTICE " + channel + " :" + str
	if channel != b.nick {
		fmt.Printf("[\033[34m%s\033[0m] \033[32;4;1m%s\033[0m: %s\n", channel, b.nick, str)
	}
}

func (b *BotClient) process_privmsg(data [][]string) {
	user_ := User{data[0][1], data[0][2], data[0][3]}
	message := &Message{user_, data[0][4], data[0][5], false}
	if len(message.Content) > 9 && message.Content[0:8] == "\001ACTION " {
		message.Content = message.Content[8 : len(message.Content)-1]
		message.Action = true
	}

	b.notice(message.Content, message.Channel)
	if message.Action {
		fmt.Printf("[\033[34m%s\033[0m] ** \033[4;1m%s\033[0m \033[1m%s\033[0m **\n", message.Channel, message.User.Nick, message.Content)
	} else if message.Content == "\001VERSION\001" {
		b.WriteChan <- b.nick + " v1.0"
	} else {
		fmt.Printf("[\033[34m%s\033[0m] \033[4;1m%s\033[0m: %s\n", message.Channel, message.User.Nick, message.Content)

		// Check if we're being addressed
		if len(message.Content) > len(b.nick)+2 && (message.Content[0:len(b.nick)+2] == b.nick+", " || message.Content[0:len(b.nick)+2] == b.nick+": ") {
			if message.Content[len(b.nick)+2:len(message.Content)] == "hello" {
				b.privmsg("Hello, "+message.User.Nick, message.Channel)
			}

			if message.Content[len(b.nick)+2:len(message.Content)] == "!quit" {
				b.WriteChan <- "QUIT :Bye!"
				time.Sleep(2)
				os.Exit(1)
			}
		}

		if message.Content == "botsnack" {
			b.privmsg("<3", message.Channel)
		}
	}
}

func (b *BotClient) process_modemsg(data [][]string) {
	user_ := User{data[0][1], data[0][2], data[0][3]}
	message := &Message{user_, data[0][4], data[0][5], false}
	fmt.Printf("[\033[34m%s\033[0m] \033[4;1m%s\033[0m changed mode to \033[32m%s\033[0m\n", message.Channel, message.User.Nick, message.Content)
}


func (b *BotClient) warn(message string) {
	fmt.Fprintf(os.Stderr, "\033[31;1mERROR\033[0m: %s\n", message)
}

func main() {
	bot := new(BotClient)
	bot.LoadConfig(os.Args[1], os.Args[2], os.Args[3:len(os.Args)])
	bot.Init(true)
	bot.Connect()
	bot.Run()
}
