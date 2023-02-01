package message

import (
	"github.com/segmentio/kafka-go"
)

func CreateMessage(input []string) []kafka.Message {
	msgs := []kafka.Message{}
	for _, v := range input {
		msgs = append(msgs, kafka.Message{
			Value: []byte(v),
		})
	}
	return msgs
}
