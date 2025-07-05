# Tom Chat

Tom Chat is a friendly CLI chatbot built with Go that leverages streaming responses from the Gemini API. It features a neat terminal interface with support for Markdown-style formatting, command history, and custom keybindings (e.g., `Enter` to submit and `Ctrl+Enter` to insert a new line). Chat history is persisted in a JSON file, so your conversations can be resumed in future sessions.

## Features

- **Streaming Responses:** Messages stream in real-time from the Gemini API.
- **Markdown Formatting:** Bot responses can be rendered with ANSI styling using [glamour](https://github.com/charmbracelet/glamour).
- **Custom Input Handling:** Uses [readline](https://github.com/chzyer/readline) for interactive input with support for:
  - `Enter` to submit your message.
  - `Ctrl+Enter` to insert a newline.
  - Arrow-key navigation and history.
- **Chat History Persistence:** All conversations are stored in `data/chat_history.json` for later retrieval.

## Prerequisites

- [Go](https://golang.org/dl/) 1.18 or later.
- A valid Gemini API key. Store it in a `.env` file as `GEMINI_API_KEY=your_key_here`.

## Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/PriyanshuPz/tom-chat.git
   cd tom-chat
   ```


2. **Install dependencies:**

   ```bash
   go get github.com/joho/godotenv
   go get github.com/chzyer/readline
   go get github.com/charmbracelet/glamour
   go get google.golang.org/genai
   ```

3. **Create the Data Directory:**
   Ensure the `data/` directory exists:

   ```bash
   mkdir -p data
   ```

4. **Set up your environment:**
   Create a `.env` file in the root directory:

   ```dotenv
   GEMINI_API_KEY=your_gemini_api_key_here
   ```

## Usage

Run the chatbot with:

```bash
go run main.go
```

- **Input:**

  - Press `Enter` to submit your message.
  - Press `Ctrl+Enter` to insert a newline (for multi-line messages).

- **Exit:**
  Type `\bye` to quit the application.

## Project Structure

```
tom-chat/
├── go.mod                   # Module file
├── main.go                  # Entry point; contains main loop and integration logic
├── config/
│   └── env.go               # Loads environment variables from .env
├── cmd/
│   ├── assistant.go               # Handles Gemini API streaming
│   ├── prompt.go            # Builds the system prompt including user info
│   ├── model.go             # Data structures (UserProfile, ChatMessage, ChatHistory)
│   └── memory.go            # JSON persistence (load/save chat history)
├── data/
│   └── chat_history.json    # Chat history file (auto-created)
└── README.md                # This file
```

