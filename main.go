package main

import (
	"encoding/json"
	"log"
)

type Message struct {
	Name string
	Body string
}

func main() {
	b := []byte(`{"Name":"Bob","Body":"Pickle"}`)
	var m Message
	err := json.Unmarshal(b, &m)
	if err != nil {
		log.Println(err)
	}
	log.Println(m)
}
