package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
	
	"github.com/yanzay/tbot/v2"
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
	fmt.Printf("Sto trasmettendo...")
	var quote string
	bot := tbot.New(os.Getenv("TELEGRAM_TOKEN"))
	c := bot.Client()

	var help int 
	var citazione int
	var manifesto int
	var opere int
	var didattica int
	
	// il comando /help restituisce la lista dei comandi disponibili
	bot.HandleMessage("/help", func(m *tbot.Message) {
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		time.Sleep(1 * time.Second)
		helpfile := readFile("help.txt")
		c.SendMessage(m.Chat.ID, helpfile)
		help++
		t := time.Now()
		fmt.Printf("\nHelp: %d --- %v", help, t)
	})
	
	//il comando /citazione restitutisce una citazione casuale di Marx
	bot.HandleMessage("/citazione", func(m *tbot.Message) {
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		time.Sleep(1 * time.Second)
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(23)
		rS := strconv.Itoa(r)

		quote = readFile("./Citazioni/"+rS)

		c.SendMessage(m.Chat.ID, quote)
		citazione++
		t := time.Now();
		fmt.Printf("\nCitazioni: %d --- %v", citazione, t)
	})

	// il comando /manifesto restituisce il link al manifesto del pc su marxists.org
	bot.HandleMessage("/manifesto", func(m *tbot.Message) {
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		time.Sleep(1 * time.Second)
		c.SendMessage(m.Chat.ID, "https://www.marxists.org/italiano/marx-engels/1848/manifesto/index.htm")

		manifesto++
		t := time.Now()
		fmt.Printf("\nManifesto, %d --- %v", manifesto, t)
	})

	// il comando /opere restituisce un link alle opere marxiste.
	bot.HandleMessage("/opere", func(m *tbot.Message) {
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		time.Sleep(1 * time.Second)
		c.SendMessage(m.Chat.ID, "https://www.marxists.org/italiano/marx-engels/index.htm")

		opere++
		t := time.Now()
		fmt.Printf("\nOpere: %d --- %v", opere, t);
	})

	// il  comando /iniziamo fornisce un link alla documentazione di base sul pensiero Marxista
	bot.HandleMessage("/didattica", func(m *tbot.Message) {
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		time.Sleep(1 * time.Second)
		c.SendMessage(m.Chat.ID, "https://www.marxists.org/italiano/sezione/studenti/index.htm")

		didattica++
		t := time.Now()
		fmt.Printf("\nDidattica: %d --- %v", didattica, t)
	})

	err := bot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
