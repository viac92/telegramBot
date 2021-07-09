package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"unicode"
	"strings"
	"math/rand"

	"github.com/yanzay/tbot/v2"
)

func randomNumber(n int) int {
	rand.Seed(time.Now().UnixNano());
	return rand.Intn(n)
}

//cammello data una strina la restituisce con lettere alternate minuscole maiuscole
func cammello(s string) string {
	var strOut string
	for _,r := range s {
		if unicode.IsLetter(r) {
			
			if !strings.ContainsAny(string(r), "iI") {

					random := randomNumber(2)
					
					if random == 0 {
					strOut += string(unicode.ToUpper(r))
					} else {
					strOut += string(r)
				}
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
		if strings.ContainsAny(string(rSli[i]), "aeoiuAEOIU") {
			vocale := rSli[i]
			rand := randomNumber(3)
			for j := 0; j < rand; j++ {
				strOut += string(vocale)
			}
		}

		if rSli[i] == '!' {
			
			rand := randomNumber(5)
			for j := 0; j < rand + 5; j++ {

				rand2 := randomNumber(5)
				switch rand2 {
				
				case 0,1,2:
					strOut += "!"
				case 3:
					strOut += "$"
				case 4: 
					strOut += "1"
				}
			}
		}

		if rSli[i] == '?' {
			rand := randomNumber(5)
			for j := 0; j < rand + 5; j++ {

				rand2 := randomNumber(5)
				switch rand2 {
				
				case 0,1,2:
					strOut += "?"
				case 3:
					strOut += "/"
				case 4: 
					strOut += "1"
				}
			}
		}

		if rSli[i] == '.' {
			for j := 0; j < 10; j++ {
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

	var kaffe int

	bot.HandleMessage("kaffe .+", func(m *tbot.Message) {
		text := strings.TrimPrefix(m.Text, "kaffe ")
		nuovaFrase := sostituisciCconK(cammello(moltiplica(text)))
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		time.Sleep(1 * time.Second)
		
		c.SendMessage(m.Chat.ID, nuovaFrase)
		kaffe++
		fmt.Printf("Kaffe bevuti: %d\n", kaffe)
	})

	err := bot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
