// 1. **Каналы для записи сообщений**: Несколько пользователей передают изменения для записи в файлы через каналы.
// 1.1 Пользователей можно сэмулировать обычным циклом

package main

import (
	"sync"
	"time"

	en "asssement1.ru/entities"
)

var stack = []en.Message{
	{Token: "111", FileID: "file_01", Data: "Payload 1 From user with token 111"},
	{Token: "111", FileID: "file_02", Data: "Payload 2 From user with token 111"},
	{Token: "111", FileID: "file_03", Data: "Payload 3 From user with token 111"},
	{Token: "111", FileID: "file_04", Data: "Payload 4 From user with token 111"},
	{Token: "222", FileID: "file_01", Data: "Payload 1 From user with token 222"},
	{Token: "222", FileID: "file_02", Data: "Payload 2 From user with token 222"},
	{Token: "222", FileID: "file_03", Data: "Payload 3 From user with token 222"},
	{Token: "222", FileID: "file_04", Data: "Payload 4 From user with token 222"},
	{Token: "333", FileID: "file_01", Data: "Payload 1 From user with token 333"},
	{Token: "333", FileID: "file_02", Data: "Payload 2 From user with token 333"},
	{Token: "333", FileID: "file_03", Data: "Payload 3 From user with token 333"},
	{Token: "333", FileID: "file_04", Data: "Payload 4 From user with token 333"},
	{Token: "444", FileID: "file_01", Data: "Payload 1 From user with token 444"},
	{Token: "444", FileID: "file_02", Data: "Payload 2 From user with token 444"},
	{Token: "444", FileID: "file_03", Data: "Payload 3 From user with token 444"},
	{Token: "444", FileID: "file_04", Data: "Payload 4 From user with token 444"},
}

type Generator struct {
	Queue []en.Message
}

func (g *Generator) SendMessage() <-chan en.Message {
	wg := &sync.WaitGroup{}
	out := make(chan en.Message, 4)
	go func() {
		defer close(out)
		defer wg.Wait()
		for _, message := range g.Queue {
			wg.Add(1)
			go func(m en.Message) {
				defer wg.Done()
				out <- m
				time.Sleep(time.Second * 1)
			}(message)
		}
	}()
	return out
}

func (g *Generator) AddMessage(m en.Message) {
	g.Queue = append(g.Queue, m)
}

func (g *Generator) GetTestMessages() {
	g.Queue = stack
}
