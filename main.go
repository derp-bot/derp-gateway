package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	if len(Token) == 0 {
		Token = os.Getenv("DISCORD_TOKEN")
	}

	if len(Token) == 0 {
		fmt.Println("You're missing a token; specify it with -t.")
		return
	}

	// Create a new Discord session using the provided bot token.
	client, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	client.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	client.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = client.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	client.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if message.Author.ID == session.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if message.Content == "ping" {
		session.ChannelMessageSend(message.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if message.Content == "pong" {
		session.ChannelMessageSend(message.ChannelID, "Ping!")
	}
}
