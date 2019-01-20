package main

import (
	"encoding/json"
	"log"
	"sort"
)

type message struct {
	To, From, Text, TimeRead, TimeSent string
}

func parseMessages(body string) []message {
	messages := make([]message, 0)
	err := json.Unmarshal([]byte(body), &messages)
	checkError(err)
	if err != nil {
		log.Fatal(err)
	}
	sortMessages(messages)
	return messages
}

func sortMessages(messages []message) {
	sort.Sort(byDate(messages))
}

type byDate []message

func (s byDate) Len() int {
	return len(s)
}

func (s byDate) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byDate) Less(i, j int) bool{
	return s[i].TimeSent < s[j].TimeSent
}
