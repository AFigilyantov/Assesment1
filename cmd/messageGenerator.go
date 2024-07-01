// 1. **Каналы для записи сообщений**: Несколько пользователей передают изменения для записи в файлы через каналы.
// 1.1 Пользователей можно сэмулировать обычным циклом

package main

import (
	"sync"
	"time"

	. "asssement1/entities"
)

var stack = []Message{
	Message{Token: "111", FileID: "file_01", Data: "Payload 1 From user with token 111"},
	Message{Token: "111", FileID: "file_02", Data: "Payload 2 From user with token 111"},
	Message{Token: "111", FileID: "file_03", Data: "Payload 3 From user with token 111"},
	Message{Token: "111", FileID: "file_04", Data: "Payload 4 From user with token 111"},
	Message{Token: "222", FileID: "file_01", Data: "Payload 1 From user with token 222"},
	Message{Token: "222", FileID: "file_02", Data: "Payload 2 From user with token 222"},
	Message{Token: "222", FileID: "file_03", Data: "Payload 3 From user with token 222"},
	Message{Token: "222", FileID: "file_04", Data: "Payload 4 From user with token 222"},
	Message{Token: "333", FileID: "file_01", Data: "Payload 1 From user with token 333"},
	Message{Token: "333", FileID: "file_02", Data: "Payload 2 From user with token 333"},
	Message{Token: "333", FileID: "file_03", Data: "Payload 3 From user with token 333"},
	Message{Token: "333", FileID: "file_04", Data: "Payload 4 From user with token 333"},
	Message{Token: "444", FileID: "file_01", Data: "Payload 1 From user with token 444"},
	Message{Token: "444", FileID: "file_02", Data: "Payload 2 From user with token 444"},
	Message{Token: "444", FileID: "file_03", Data: "Payload 3 From user with token 444"},
	Message{Token: "444", FileID: "file_04", Data: "Payload 4 From user with token 444"},
}

type Generator struct {
	Queue []Message
}

func (g *Generator) SendMessage(wg *sync.WaitGroup) chan Message {
	out := make(chan Message)
	go func() {
		defer close(out)
		defer wg.Wait()
		for _, message := range g.Queue {
			wg.Add(1)
			go func(m Message) {
				defer wg.Done()
				time.Sleep(time.Millisecond * 100)
				out <- m
			}(message)
		}
	}()
	return out
}

func (g *Generator) AddMessage(m Message) {
	g.Queue = append(g.Queue, m)
}

func (g *Generator) GetTestMessages() {
	g.Queue = stack
}
