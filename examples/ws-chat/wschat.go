package wschat

type MessageType int

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}
