package exporter

import (
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"io"
	"net/url"
	"time"
)

func DecodeJsonResult(r io.Reader) (TelegramResult, error) {
	var result TelegramResult
	err := json.NewDecoder(r).Decode(&result)
	return result, err
}

func ExportResult(chat ChatResult, w io.Writer, escapeText, b64Text bool) error {
	csvWriter := csv.NewWriter(w)

	header := []string{"from", "date", "text"}
	if err := csvWriter.Write(header); err != nil {
		return err
	}

	for _, message := range chat.Messages {
		// filter out things that aren't a plain text message
		if message.Type != "message" || len(message.Text) == 0 {
			continue
		}

		var text = string(message.Text)
		if escapeText {
			text = url.QueryEscape(text)
		}
		if b64Text {
			text = base64.StdEncoding.EncodeToString([]byte(text))
		}

		date := time.Time(message.Date).Format("2006-01-02 15:04:05")

		row := []string{
			message.From,
			date,
			text,
		}
		if err := csvWriter.Write(row); err != nil {
			return err
		}
	}

	csvWriter.Flush()

	if err := csvWriter.Error(); err != nil {
		return err
	}

	return nil
}
