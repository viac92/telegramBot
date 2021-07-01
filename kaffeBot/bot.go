package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"unicode"
	"strings"

	"github.com/yanzay/tbot/v2"
)

//cammello data una strina la restituisce con lettere alternate minuscole maiuscole
func cammello(s string) string {
	var strOut string
	for i,r := range s {
		if unicode.IsLetter(r) {
			if i%2 == 0 {
				strOut += string(unicode.ToUpper(r))
			} else {
				strOut += string(r)
			}
		} else {
			strOut += string(r)
		}
	}
	return strOut
}

//sotituisciCconK data una stringa ne restituisce una con le k al posto delle c dure
func sostituisciCconK(s string) string {
	var strOut string
	rSli := []rune(s)

	for i := 0; i < len(rSli); i++ {
		if i < len(rSli) - 1 && (rSli[i] == 'c' || rSli[i] == 'C') {
			if rSli[i+1] != 'i' && rSli[i+1] != 'I' && rSli[i+1] != 'e' && rSli[i+1] != 'E' {
				strOut += "k"
				continue
			}
		} 
		strOut += string(rSli[i])
	}
	return strOut
}

func moltiplica(s string) string {
	var strOut string
	rSli := []rune(s)

	for i := 0; i < len(rSli); i++ {
		if rSli[i] == '!' {
			for j := 0; j < 10; j++ {
				if j%2 == 0 {
					strOut += "!"
				} else {
					strOut += "?"
				}
			}
			continue
		}

		if rSli[i] == '?' {
			for j := 0; j < 10; j++ {
				if j%2 == 0 {
					strOut += "?"
				} else {
					strOut += "!"
				}
			}
			continue
		}

		if rSli[i] == '.' {
			for j := 0; j < 3; j++ {
				strOut += "."
			}
			continue
		}

		strOut += string(rSli[i])
	}	
	return strOut
}




func main() {
	fmt.Println("Sto trasmettendo...")
	bot := tbot.New(os.Getenv("TELEGRAM_TOKEN2"))
	c := bot.Client()

	bot.HandleMessage("kaffe .+", func(m *tbot.Message) {
		text := strings.TrimPrefix(m.Text, "kaffe ")
		nuovaFrase := moltiplica(sostituisciCconK(cammello(text)))
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		time.Sleep(1 * time.Second)
		
		c.SendMessage(m.Chat.ID, nuovaFrase)
	})

	err := bot.Start()
	if err != nil {
		log.Fatal(err)
	}
}