package lsof

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type Connection struct {
	Message string
	Count   int
}

func summarizeConnections(reader io.Reader) []*Connection {
	response := make([]*Connection, 0)

	conn := func(message string) *Connection {
		for _, conn := range response {
			if conn.Message == message {
				return conn
			}
		}
		c := &Connection{
			Message: message,
		}
		response = append(response, c)
		return c
	}

	// Create a scanner to read input line by line
	scanner := bufio.NewScanner(reader)

	// Iterate over each line
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line by space
		parts := strings.Fields(line)

		// Skip header line
		if parts[0] == "COMMAND" {
			continue
		}

		if len(parts) < 10 {
			continue
		}

		// Extract relevant information
		hostPort := parts[8]
		state := parts[9]
		if strings.Contains(hostPort, "->") {
			hostParts := strings.Split(hostPort, "->")
			hostPort = hostParts[1]
		}

		// Construct the summary key
		summaryKey := fmt.Sprintf("%s %s", state, hostPort)

		// Update the connection summary
		linked := conn(summaryKey)
		linked.Count++
	}

	sort.Slice(response, func(i, j int) bool {
		return response[i].Count > response[j].Count
	})

	return response
}

func lsof(cfg *options) error {
	var total int
	report := summarizeConnections(os.Stdin)

	fmt.Println("Open connections:")
	for _, i := range report {
		fmt.Printf("- %d %s\n", i.Count, i.Message)
		total += i.Count
	}
	fmt.Println("Total:", total)

	return nil
}
