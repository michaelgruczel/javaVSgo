package main
// Message message
// swagger:model message
type Message struct {
	Id        int       `json:"id"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
}

type Messages []Message
