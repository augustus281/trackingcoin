package dto

type Message struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Message string `json:"message"`
	Email   string `json:"email"`
}

type Notification struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Message string   `json:"message"`
}
