package main

// This program rewrites and extends rchowe's work
// https://github.com/rchowe/go-irc-bots

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"regexp"
	"time"
	"strconv"
)

// default config
const RECV_BUF_LEN = 1024
const DEFAULT_LISTENER_PORT int = 20080
const DEFAULT_HOST string = "example:6667"
const DEFAULT_NICK string = "wjx-bot"
var   DEFAULT_CHANNELS []string = []string{"#wjx-test"}

var counter = map[string]int {}

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

// handler function type
type CallbackFunc func(data [][]string) (out string, err interface{})

type Handler struct {
	Regexp   *regexp.Regexp
	Callback CallbackFunc
}

type BotClient struct {
	SockReader *bufio.Reader
	SockWriter *bufio.Writer
	ReadChan   chan string
	WriteChan  chan string
	debug      bool
	nick       string
	host       string
	channels   []string
	handler_map map[string]*Handler
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

	b.nick = nick
	b.host = host
	b.channels = channels
	fmt.Fprintf(os.Stderr, "\033[1mUSING\033[0m: %s %s %s\n", host, nick, channels)
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
			b.WriteChan <- "PONG :" + str[7:]
		}

		data := privmsgRegexp.FindAllStringSubmatch(str, 1)
		if data != nil {
			b.processPrivMsg(data)
		}

		data = modeRegexp.FindAllStringSubmatch(str, 1)
		if data != nil {
			b.processModeMsg(data)
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

		str = str[0 : len(str)-2]

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

func (b *BotClient) startListener(port int) {
	listener, err := net.Listen("tcp", ":" + strconv.Itoa(port))
	if err != nil {
		b.warn("Error in listening localhost:" + strconv.Itoa(port))
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "\033[31;1mListener Accept ERROR\033[0m: %s\n", err)
			break
		}
		go b.handleConn(conn)
	}
}

func (b *BotClient) handleConn(conn net.Conn) {
	// One string per connection
	// No need to loop here
	//for {
		buf := make([]byte, RECV_BUF_LEN)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\033[31;1mListener Read ERROR\033[0m: %s, %v\n", err, buf)
			return
		}

		//if string(buf[0:n - 2]) == "exit" {
		//	conn.Close()
		//	break
		//}

		// just post what I received
		for _, channel := range b.channels {
			b.notice(string(buf), channel)
		}

		conn.Close()
	//}
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
	b.ReadChan = make(chan string, 1000)
	b.WriteChan = make(chan string, 1000)
	b.debug = debug
	b.handler_map = make(map[string]*Handler)

	// start goroutine to do this
	go b.startReader()
	go b.startWriter()
	go b.startListener(DEFAULT_LISTENER_PORT)

	b.Connect()
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

func (b *BotClient) processPrivMsg(data [][]string) {
	user_ := User{data[0][1], data[0][2], data[0][3]}
	message := &Message{user_, data[0][4], data[0][5], false}
	if len(message.Content) > 9 && message.Content[0:8] == "\001ACTION " {
		message.Content = message.Content[8 : len(message.Content)-1]
		message.Action = true
	}

	if b.debug {
		//b.notice("Detected : " + message.Content, message.Channel)
	}

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

		// check all handler functions
		for key, handler := range b.handler_map {
			if handler == nil {
				continue;
			}

			data := handler.Regexp.FindAllStringSubmatch(message.Content, 1)
			if data != nil {
				if b.debug {
					fmt.Printf("Handler [%s] matched!\n", key)
				}
				out, err := handler.Callback(data)
				if err == nil {
					b.privmsg(out, message.Channel)
				} else {
					str, err := err.(string)
					if (err != true) {
						b.privmsg(str, message.Channel)
					}
				}
			}
		}
	}
}

func (b *BotClient) processModeMsg(data [][]string) {
	user_ := User{data[0][1], data[0][2], data[0][3]}
	message := &Message{user_, data[0][4], data[0][5], false}
	fmt.Printf("[\033[34m%s\033[0m] \033[4;1m%s\033[0m changed mode to \033[32m%s\033[0m\n", message.Channel, message.User.Nick, message.Content)
}

func (b *BotClient) warn(message string) {
	fmt.Fprintf(os.Stderr, "\033[31;1mERROR\033[0m: %s\n", message)
}

func (b *BotClient) RegistHandler(name, regexp_pattern string, callback CallbackFunc) {
	if b.handler_map[name] != nil {
		b.warn("Already Registered Handler " + name + " ... Skip");
		return;
	}

	handler := new(Handler);
	handler.Regexp = regexp.MustCompile(regexp_pattern)
	handler.Callback = callback

	b.handler_map[name] = handler
}

//
// handler functions
//
func handleTime(data [][]string) (out string, err interface{}) {
	cmd := exec.Command("date")
	cmd_out, cmd_err := cmd.Output()
	return string(cmd_out), cmd_err
}

func handleTimestamp(data [][]string) (out string, err interface{}) {
	cmd := exec.Command("date", "+'%s'")
	cmd_out, cmd_err := cmd.Output()
	return string(cmd_out), cmd_err
}

func handleTimestampConvertion(data [][]string) (out string, err interface{}) {
	timestamp := data[0][1]
	intTimestamp, _ := strconv.ParseInt(timestamp, 10, 64)
	date := time.Unix(intTimestamp, 0)
	return date.Format("2006-01-02 15:04:05"), nil
}

func handleCount(data [][]string) (out string, err interface{}) {
	variable := data[0][1]
	operator := data[0][2]

	if _, ok := counter[variable]; ok == false {
		counter[variable] = 0
	}

	if operator == "++" {
		counter[variable]++
	} else if operator == "--" {
		counter[variable]--
	}
	return variable + " : " + strconv.Itoa(counter[variable]), nil
}

func main() {
	bot := new(BotClient)

	// Load default config or command line argument
	if len(os.Args) < 4 || os.Args[1] == "" || os.Args[2] == "" || os.Args[3] == "" {
		bot.LoadConfig(DEFAULT_HOST, DEFAULT_NICK, DEFAULT_CHANNELS)
	} else {
		bot.LoadConfig(os.Args[1], os.Args[2], os.Args[3:len(os.Args)])
	}

	bot.Init(true)

	//
	// regist handler functions here
	//
	bot.RegistHandler("time", "^time$", handleTime)
	bot.RegistHandler("timestamp", "^timestamp$", handleTimestamp)
	bot.RegistHandler("timestamp_to_date", "^date\\s+(\\d+)$", handleTimestampConvertion)
	bot.RegistHandler("counter", "^(\\S+)(\\+\\+|\\-\\-)$", handleCount)

	bot.Run()
}
