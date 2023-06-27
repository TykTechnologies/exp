package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	cfg := NewOptions()

	if err := start(cfg); err != nil {
		log.Fatal(err)
	}
}

type Message struct {
	Time, Package string
	Test          string  `json:"Test,omitempty"`
	Action        string  `json:"Action"`
	Elapsed       float64 `json:",omitempty"`
}

type Result struct {
	RawMessage json.RawMessage
	Message    *Message
}

func start(cfg *options) error {
	f, err := os.Open(cfg.inputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	messages := []Result{}

	for {
		if decoder.More() {
			var m json.RawMessage
			if err := decoder.Decode(&m); err != nil {
				return err
			}

			m2 := &Message{}
			if err := json.Unmarshal(m, m2); err != nil {
				return err
			}

			if m2.Action != "pass" && m2.Action != "fail" {
				messages = append(messages, Result{
					RawMessage: m,
				})
				continue
			}

			messages = append(messages, Result{
				RawMessage: m,
				Message:    m2,
			})
			continue
		}
		break
	}

	if len(messages) == 0 {
		return nil
	}

	total := *(messages[len(messages)-1].Message)
	total.Elapsed = 0.0

	for _, message := range messages {
		msg := message.Message
		msgText := string(message.RawMessage)
		if msg == nil {
			if strings.Contains(msgText, "coverage") {
				continue
			}
			fmt.Println(msgText)
			continue
		}

		if msg.Test != "" {
			fmt.Println(msgText)
		}

		if msg.Test == "" && msg.Action != "pass" {
			total.Action = msg.Action
		}
		total.Elapsed += msg.Elapsed
	}

	b, _ := json.Marshal(total)
	fmt.Println(string(b))

	return nil
}
