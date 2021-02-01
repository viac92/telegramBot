// make a marxist bot for all comrades
package main

import (
	"log"
	"os"
	"time"
	"math/rand"

	"github.com/yanzay/tbot"
)



func main() {
	var quote string
	bot := tbot.New(os.Getenv("TELEGRAM_TOKEN"))
	c := bot.Client()
	
	bot.HandleMessage("/priviet", func(m *tbot.Message) {
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		time.Sleep(1 * time.Second)
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(11)
		
		switch r {
		case 0:
			quote = "The oppressed are allowed once every few years to decide which particular representatives of the oppressing class are to represent and repress them."
		case 1:
			quote = "The philosophers have only interpreted the world, in various ways. The point, however, is to change it."
		case 2: 
			quote = "Surround yourself with people who make you happy.\nPeople who make you laugh, who help you when you’re in need.\nPeople who genuinely care.\nThey are the ones worth keeping in your life.\nEveryone else is just passing through."	
		case 3:
			quote = "The last capitalist we hang shall be the one who sold us the rope."
		case 4:
			quote = "Let the ruling classes tremble at a Communistic revolution.\nThe proletarians have nothing to lose but their chains. They have a world to win."
		case 5:
			quote = "Workingmen of all countries unite!"
		case 6:
			quote = "Reason has always existed, but not always in a reasonable form."
		case 7:
			quote = "Go on, get out! Last words are for fools who haven't said enough!"
		case 8:
			quote = "To be radical is to grasp things by the root."
		case 9:
			quote = "The proletarians have nothing to loose but their chains.\nThey have a world to win."
		case 10:
			quote = "In proportion therefore, as the repulsiveness of the work increases, the wage decreases."
		}
		c.SendMessage(m.Chat.ID, quote)
	})
	err := bot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
