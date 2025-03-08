package chat

import chat "mini/internal/processor/version/procedure/chat"

type Procedure struct {
}

func NewChatProcedure() chat.Procedure {
	return &Procedure{}
}
