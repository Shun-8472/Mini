package chat

type Request struct {
	UserInput string `json:"user_input"`
}

type Response struct {
	Message []byte `json:"message"`
}
