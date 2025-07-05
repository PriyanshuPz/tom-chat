package cmd

type UserProfile struct {
	Name        string `json:"name"`
	Likes       string `json:"likes"`
	Job         string `json:"job"`
	Personality string `json:"personality"`
	FunFact     string `json:"fun_fact"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Message string `json:"message"`
}

type ChatHistory struct {
	User    UserProfile   `json:"user"`
	History []ChatMessage `json:"history"`
}
