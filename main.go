package main

import (
	"fmt"
	"strings"

	"github.com/chzyer/readline"
	"p8labs.tech/tom/cmd"
	"p8labs.tech/tom/config"
)

func main() {
	config.LoadEnv()

	history := cmd.LoadOrInitHistory("data/chat_history.json")

	fmt.Println("Hi", history.User.Name+"! Type '.exit' to quit.")

	for {
		input, err := readInput()
		input = strings.TrimSpace(input)

		if err != nil {
			fmt.Println("❌ Input error:", err)
			break
		}

		if input == ".exit" {
			fmt.Println("👋 Goodbye!")
			break
		}

		cmd.Chat(input, &history)

		cmd.SaveHistory("data/chat_history.json", history)
	}
}

func readInput() (string, error) {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:                 "👉 ",
		HistoryFile:            "/tmp/chatbot_history.tmp",
		InterruptPrompt:        "^C",
		EOFPrompt:              "exit",
		DisableAutoSaveHistory: true,
		FuncFilterInputRune: func(r rune) (rune, bool) {
			// Handle Ctrl+Enter (Ctrl+J) to allow new line
			if r == readline.CharCtrlJ {
				return '\n', true
			}
			return r, true
		},
	})

	if err != nil {
		return "", err
	}
	defer rl.Close()

	var buffer strings.Builder

	for {
		line, err := rl.Readline()
		if err != nil {
			return "", err
		}

		// Submit on regular Enter (empty line or no trailing \)
		if !strings.HasSuffix(line, "\\") {
			buffer.WriteString(line)
			break
		}

		// Remove the trailing backslash and add a real newline
		line = strings.TrimSuffix(line, "\\")
		buffer.WriteString(line + "\n")
		rl.SetPrompt("... ") // prompt for multiline
	}

	rl.SetPrompt("👉 ") // reset prompt
	return buffer.String(), nil
}
