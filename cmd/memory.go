package cmd

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func LoadOrInitHistory(filename string) ChatHistory {
	history, err := loadHistory(filename)
	if err == nil {
		return history
	}

	// Default profile
	return ChatHistory{
		User: UserProfile{
			Name:        "Priyanshu",
			Likes:       "Coding, coffee, memes",
			Job:         "Software Developer, Student",
			Personality: "Chill, funny, loves tech",
			FunFact:     "Loves building side projects in dev mode",
		},
	}
}

func loadHistory(filename string) (ChatHistory, error) {
	var history ChatHistory
	file, err := os.Open(filename)
	if err != nil {
		return history, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&history)
	return history, err
}

func SaveHistory(filename string, history ChatHistory) {

	// Ensure directory exists
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Println("⚠️ Failed to create directory:", err)
		return
	}
	// Now write the file
	file, err := os.Create(filename)
	if err != nil {
		log.Println("⚠️ Failed to save chat:", err)
		return
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(history); err != nil {
		log.Println("⚠️ Failed to encode JSON:", err)
	}
}
