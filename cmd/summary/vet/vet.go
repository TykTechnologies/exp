package vet

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type Warning struct {
	Position string `json:"posn"`
	Message  string `json:"message"`
}

func (w Warning) File() string {
	return strings.SplitN(w.Position, ":", 2)[0]
}

type Package map[string][]Warning

func vet(cfg *options) error {
	scanner := bufio.NewScanner(os.Stdin)

	jsonStream := ""

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		jsonStream += line
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	reader := strings.NewReader(jsonStream)
	decoder := json.NewDecoder(reader)

	report := map[string]Package{}
	for {
		row := map[string]Package{}
		if err := decoder.Decode(&row); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		// Each package is unique, merge report (slurp)
		for k, v := range row {
			report[k] = v
		}
	}

	fmt.Println("Warnings by message:")
	byMessage := vetByMessage(report)
	for _, i := range byMessage {
		fmt.Printf("%d %s\n", i.Count, i.Message)
	}
	fmt.Println()

	fmt.Println("Warnings by filename:")
	byFile := vetByFile(report)
	for _, i := range byFile {
		fmt.Printf("%d %s\n", i.Count, i.Message)
	}
	fmt.Println()

	fmt.Println("Total:", vetCount(report))

	return nil
}

type Message struct {
	Message string
	Count   int
}

func vetCount(report map[string]Package) (response int) {
	for _, pkg := range report {
		for _, kind := range pkg {
			response += len(kind)
		}
	}

	return response
}

func vetByFile(report map[string]Package) []*Message {
	response := []*Message{}

	findOrCreate := func(message string) *Message {
		for _, r := range response {
			if r.Message == message {
				return r
			}
		}

		res := &Message{
			Message: message,
		}
		response = append(response, res)
		return res
	}

	for _, pkg := range report {
		for _, kind := range pkg {
			for _, w := range kind {
				m := findOrCreate(w.File())
				m.Count++
			}
		}
	}

	sort.Slice(response, func(i, j int) bool {
		return response[i].Count > response[j].Count
	})

	return response
}

func vetByMessage(report map[string]Package) []*Message {
	response := []*Message{}

	findOrCreate := func(message string) *Message {
		for _, r := range response {
			if r.Message == message {
				return r
			}
		}

		res := &Message{
			Message: message,
		}
		response = append(response, res)
		return res
	}

	for _, pkg := range report {
		for _, kind := range pkg {
			for _, w := range kind {
				m := findOrCreate(w.Message)
				m.Count++
			}
		}
	}

	sort.Slice(response, func(i, j int) bool {
		return response[i].Count > response[j].Count
	})

	return response
}
