package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func checkNotError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var channel string

func init() {
	err := godotenv.Load(".env")
	checkNotError(err)
	channel = os.Getenv("CHANNEL")
}

func main() {
	session, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	checkNotError(err)

	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {

		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "raise" {
			s.ChannelMessageSend(m.ChannelID, "Hello world!")
		} else if m.Content == "whoareyou" {
			s.ChannelMessageSend(m.ChannelID, "I'm **"+s.State.User.Username+"**, that's it :)")
		} else if m.Content == "configdate" {
			fmt.Println("Message author id" + m.Author.ID + " | " + m.Author.Username + "\nChannel id: " + m.ChannelID)
		} else if m.Content == "debug" {
			for i := 0; i < len(Debug_dates); i++ {
				s.ChannelMessageSend(channel, DebugMessage(i))
				fmt.Println("MESSAGE " + string(rune(i+1)) + " SENDED")
			}
		}
	})

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = session.Open()
	checkNotError(err)

	defer session.Close()
	fmt.Println("Gobbler is on the wire!")
	color.Cyan(`
	  _____       _     _     _           
	 / ____|     | |   | |   | |          
   	| |  __  ___ | |__ | |__ | | ___ _ __ 
   	| | |_ |/ _ \| '_ \| '_ \| |/ _ \ '__|
   	| |__| | (_) | |_) | |_) | |  __/ |   
	 \_____|\___/|_.__/|_.__/|_|\___|_|
	 
	`)
	regularSend(session)

	if len(os.Args) > 1 {
		if os.Args[1] == "-f" {
			sc := make(chan os.Signal, 1)
			signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
			<-sc
		}
	}
}
