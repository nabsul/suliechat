package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func getMyMessages() {
	path := "messages"
	result := request("GET", path, nil)
	messages := parseMessages(result)
	for _, m := range messages {
		fmt.Printf("[%s]: from %s to %s\n%s\n\n", convertTime(m.TimeSent), m.From, m.To, m.Text)
	}
}

func sendMessage() {
	args := os.Args
	to, message := args[2], strings.Join(args[3:], " ")
	body, err := json.Marshal(map[string]string{"to": to, "text": message})
	checkError(err)
	response := request("POST", "message", strings.NewReader(string(body)))
	fmt.Println(response)
}

func listUsers() {
	result := request("GET", "users", nil)
	users := make([]struct{Username string}, 0)
	err := json.Unmarshal([]byte(result), &users)
	checkError(err)
	fmt.Println("Users:")
	for _, u := range users {
		fmt.Println("  " + u.Username)
	}
}

func getHelp() {
	fmt.Println("SulieChat: A Command-line Messaging Client")
	fmt.Println("")
	fmt.Println("suliechat help\n   This message")
	fmt.Println("suliechat config [server] [username] [password]\n   Create a new configuration")
	fmt.Println("suliechat check\n   Show my messages")
	fmt.Println("suliechat send [user] [message]\n   Send message to user")
	fmt.Println("suliechat users\n   List all users")
}

var commands = map[string] func()() {
	"help": getHelp,
	"config": saveConfig,
	"check": getMyMessages,
	"send": sendMessage,
	"users": listUsers,
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
