package main

import (
	"strconv"
	"bufio"
	"log"
	"os"
	"time"
	"math/rand"

	"github.com/yanzay/tbot"
)

// radFile legge un file e restituisce una stringa contenente il contenuto
func readFile(fileName string) string {
	var stringOut string

	f,_ := os.Open(fileName)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
			stringOut += scanner.Text() + "\n"

	}

	return stringOut
}

func main() {
	var quote string
	bot := tbot.New(os.Getenv("TELEGRAM_TOKEN"))
	c := bot.Client()
	
	// il comando /help restituisce la lista dei comandi disponibili
	bot.HandleMessage("/help", func(m *tbot.Message) {
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		time.Sleep(1 * time.Second)
		helpfile := readFile("help.txt")
		c.SendMessage(m.Chat.ID, helpfile)
	})
	
	//il comando /citazione restitutisce una citazione casuale di Marx
	bot.HandleMessage("/citazione", func(m *tbot.Message) {
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		time.Sleep(1 * time.Second)
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(22)
		rS := strconv.Itoa(r)

		quote = readFile("./Citazioni/"+rS)

		c.SendMessage(m.Chat.ID, quote)
	})

	// il comando /manifesto restituisce il link al manifesto del pc su marxists.org
	bot.HandleMessage("/manifesto", func(m *tbot.Message) {
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		time.Sleep(1 * time.Second)
		c.SendMessage(m.Chat.ID, "https://www.marxists.org/italiano/marx-engels/1848/manifesto/index.htm")
	})

	// il  comando /iniziamo fornisce un link alla documentazione di base sul pensiero Marxista
	bot.HandleMessage("/didattica", func(m *tbot.Message) {
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		time.Sleep(1 * time.Second)
		c.SendMessage(m.Chat.ID, "https://www.marxists.org/italiano/sezione/studenti/index.htm")
	})

	err := bot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
