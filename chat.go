package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func getMyMessages() {
	getMessages("")
}

func getSentMessages() {
	getMessages("?to=" + os.Args[2])
}

func sendMessage() {
	args := os.Args
	to, message := args[2], strings.Join(args[3:], " ")
	body, err := json.Marshal(map[string]string{"to": to, "text": message})
	checkError(err)
	response := request("POST", "message", strings.NewReader(string(body)))
	fmt.Println(response)
}

func getHelp() {
	fmt.Println("SulieChat: A Command-line Messaging Client")
	fmt.Println("")
	fmt.Println("suliechat help\n   This message")
	fmt.Println("suliechat config [server] [username] [password]\n   Create a new configuration")
	fmt.Println("suliechat messages\n   Show my messages")
	fmt.Println("suliechat sent [user]\n   Show messages I sent to a user")
	fmt.Println("suliechat send [user] [message]\n   Send message to user")
}

var commands = map[string] func()() {
	"help": getHelp,
	"config": saveConfig,
	"messages": getMyMessages,
	"sent": getSentMessages,
	"send": sendMessage,
}

func main() {
	cmd := "help"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	method, ok := commands[cmd]
	if !ok {
		log.Fatal("Command not found")
	}

	method()
}