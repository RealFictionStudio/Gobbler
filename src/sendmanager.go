package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var Debug_dates []time.Time = []time.Time{
	Date(2023, 8, 2),
	Date(2023, 9, 4),
	Date(2023, 10, 31),
	Date(2023, 12, 6),
	Date(2023, 12, 24),
	Date(2023, 12, 31),
	Date(2024, 1, 1),
	Date(2024, 2, 14),
	Date(2024, 5, 2),
}

func CheckIfPosted(hasPosted *bool) {
	file, err := ioutil.ReadFile("sended.log")
	checkNotError(err)

	y, m, d := time.Now().Date()
	var currentDate string = fmt.Sprintf("%v %v %v", y, m, d)
	var dates []string = strings.Split(string(file), "\n")

	for k, v := range dates {
		_ = k
		if currentDate == v {
			*hasPosted = true
			return
		}
	}

	*hasPosted = false
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func CalcDaysRemaining(debug bool, debug_date time.Time) string {
	var t1, t2 time.Time
	if !debug {
		t1 = Date(2024, 5, 6)
		t2 = Date(time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	} else {
		t1 = Date(2024, 5, 6)
		t2 = debug_date
	}
	days := int(t1.Sub(t2).Hours() / 24)

	if days == 0 {
		return "now"
	} else if days < 0 {
		return ""
	} else {
		var timeInfo string = fmt.Sprint(days)

		if days == 1 {
			timeInfo += "dzień!!! :timer:"
		} else {
			timeInfo += "dni :timer:"
		}

		var cd, cm int = t2.Day(), int(t2.Month())
		if cm == 9 && cd == 4 {
			timeInfo += ":checkered_flag:"
		} else if cm == 10 && cd == 31 {
			timeInfo += ":jack_o_lantern:"
		} else if cm == 12 {
			if cd == 6 {
				timeInfo += ":snowman:"
			} else if cd == 24 {
				timeInfo += ":santa:"
			} else if cd == 31 {
				timeInfo += ":firecracker:"
			}
		} else if cm == 1 && cd == 1 {
			timeInfo += ":firecracker:"
		} else if cm == 2 && cd == 14 {
			timeInfo += ":heart:"
		} else if cm == 5 && cd < 6 {
			timeInfo += ":fire:"
		}

		return timeInfo

	}
}

func DebugMessage(mID int) string {

	return CalcDaysRemaining(true, Debug_dates[mID])
}

func regularSend(s *discordgo.Session) {
	if time.Now().Hour() >= 8 {
		var posted bool
		CheckIfPosted(&posted)

		if !posted {
			f, err := os.OpenFile("sended.log", os.O_WRONLY, 0644)
			checkNotError(err)

			y, m, d := time.Now().Date()

			f.WriteString(fmt.Sprintf("\n%v %v %v", y, m, d))
			f.Close()

			var dayLeft string = CalcDaysRemaining(false, Date(0, 0, 0))

			if dayLeft == "now" {
				s.ChannelMessageSend(channel, "**Wielka bitwa się zaczyna** 🗡️")
			} else if dayLeft != "" {
				s.ChannelMessageSend(channel, "**Czas do ostatecznego starcia: "+dayLeft+"**")
			}

		}

	}
}
