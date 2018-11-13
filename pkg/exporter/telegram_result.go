package exporter

import (
	"encoding/json"
	"time"
)

type TelegramResult struct {
	Result struct {
		Chats []ChatResult `json:"list,omitempty"`
	} `json:"chats,omitempty"`
}

func (t TelegramResult) FindChat(chat string) *ChatResult {
	for _, cr := range t.Result.Chats {
		if cr.Name == chat {
			return &cr
		}
	}
	return nil
}

type ChatResult struct {
	Name     string          `json:"name,omitempty"`
	Messages []MessageResult `json:"messages,omitempty"`
}

type MessageResult struct {
	Type string     `json:"type,omitempty"`
	Date TimeResult `json:"date,omitempty"`
	From string     `json:"from,omitempty"`
	Text TextResult `json:"text,omitempty"`
}

type TextResult string

func (t *TextResult) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		return nil
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*t = TextResult(s)
	return nil
}

type TimeResult time.Time

func (t *TimeResult) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	format := "2006-01-02T15:04:05"
	parsedTime, err := time.Parse(format, s)
	if err != nil {
		return err
	}
	*t = TimeResult(parsedTime.Local())
	return nil
}
