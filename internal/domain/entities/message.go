package entities

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkoukk/tiktoken-go"
)

type Message struct {
	ID        string
	Role      string
	Content   string
	Tokens    int
	Model     *Model
	CreatedAt time.Time
}

func NewMessage(role, content string, model *Model) (*Message, error) {
	tke, err := tiktoken.EncodingForModel(model.GetModelName())

	if err != nil {
		err = fmt.Errorf("getEncoding: %v", err)
		return nil, err
	}

	token := tke.Encode(content, nil, nil)

	totalTokens := len(token)

	msg := &Message{
		ID:        uuid.New().String(),
		Role:      role,
		Content:   content,
		Tokens:    totalTokens,
		Model:     model,
		CreatedAt: time.Now(),
	}

	if err := msg.Validate(); err != nil {
		return nil, err
	}

	return msg, nil
}

func (m *Message) Validate() error {
	if m.Role != "user" && m.Role != "system" && m.Role != "assistant" {
		return errors.New("invalid role")
	}

	if m.Content == "" {
		return errors.New("content is empty")
	}

	if m.CreatedAt.IsZero() {
		return errors.New("invalid created at")
	}

	return nil
}

func (m *Message) getQtdTokens() int {
	return m.Tokens
}
