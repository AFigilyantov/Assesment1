/*
Реализация многопользовательской системы чтения/записи с кешированием

## Условия задачи

Компоненты системы:

1. **Каналы для записи сообщений**: Несколько пользователей передают изменения для записи в файлы через каналы.
1.1 Пользователей можно сэмулировать обычным циклом
2. **Кеширование сообщений**: Сообщения из каналов кешируются.
3. **Воркер**: Воркер проходит по кешу и каждые секунду записывает изменения в соответствующие файлы.
4. **Валидация токенов**: Если сообщение приходит от пользователя с неправильным токеном, оно не должно быть записано.
5. **Graceful Shutdown**. При остановке приложения все закешированные данные должны записаться в файл

## Требования

1. **Структура данных для сообщений**:
   - Каждое сообщение должно содержать идентификатор пользователя, токен, идентификатор файла и данные для записи.
   - Пример структуры для сообщения:
     ```go
     type Message struct {
       Token     string
       FileID    string
       Data      string
     }
     ```

2. **Кеширование сообщений**:
   - Реализуйте обработчик, который будет вычитывать данные из каналов и помещать их в кеш, если токен пользователя валидный.
   - Используйте map для хранения сообщений в кеше. Ключом служит идентификатор файла, а значением список сообщений.

3. **Валидация токенов**:
   - Создайте функцию или метод для проверки токенов пользователей. Храните допустимые токены(white-list) в специальной структуре

4. **Воркер для записи данных в файлы**:
   - Воркер должен периодически (раз в секунду и задается с конфига) проходить по кешу и записывать все изменения в соответствующие файлы.
   - После записи кеш для этих файлов должен быть очищен.


## Cхема

![](./schema.png "Schema")



## Cценарии (можно и тестики по ним написать)

### Сценарий 1: Успешная запись
1. Пользователь отправляет сообщение с правильным токеном в канал записи.
2. Сообщение кешируется.
3. Воркер через заданный интервал времени (секунда) берет сообщение из кеша.
4. Воркер записывает сообщение в целевой файл.
5. Кеш очищается для этого файла.

### Сценарий 2: Неверный токен
1. Пользователь отправляет сообщение с неправильным токеном в канал записи.
2. Сообщение проверяется на валидность токена.
3. Сообщение не кешируется и отбрасывается.

### Сценарий 3: Остановка приложения (Graceful Shutdown)
1. Приложение получает сигнал остановки.
2. Воркер проходит по кешу и записывает все оставшиеся сообщения в соответствующие файлы.
3. Приложение завершает работу.

### Сценарий 4: Высокая нагрузка
1. Пользователи массово отправляют сообщения в каналы записи.
2. Обработчики записывают сообщения в кеш.
3. Воркер масштабируется

### Cценарий 5: Файл с одновременной записью
1. Корректная запись данных в один и тот же файл, когда несколько пользователей одновременно отправляют сообщения для него.
2. Синхронизация, чтобы избежать конфликтов и потери данных.

### Cценарий 6: Сбор работы воркера
1. Если воркер сталкивается с ошибкой при записи данных в файл (недоступность файла, отказ диска, etc), система должна предпринять меры по обработке этой ситуации.
2. Ретраи, получается
*/

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	. "asssement1.ru/entities"
)

func main() {

	var whiteList = NewUsers()

	whiteList.AddNewUser("111")
	whiteList.AddNewUser("222")
	whiteList.AddNewUser("333")
	whiteList.AddNewUser("444")

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	g := Generator{}     // создается генератор
	g.GetTestMessages()  // заполняем етстовые данные
	fc := NewFileCache() // создаем cache сообщений

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM) // gracefull Shutdown
	defer stop()

	messageFromOutSide := g.SendMessage(wg) // пишем в канал сообщения из генератора

	fc.WriteDataTo(wg, mu, messageFromOutSide, whiteList)
	go WriteDataFrom(ctx, wg, fc)

	<-ctx.Done()
	wg.Wait()

}

func WriteDataFrom(ctx context.Context, wg *sync.WaitGroup, fc *FileCache) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ParseFileCache(wg, fc)
		case <-ctx.Done():
			fmt.Println("INTERRUPTED")
			return

		}
	}
}

func writeText(td TemporaryData) <-chan TemporaryData {

	retryChan := make(chan TemporaryData)
	defer close(retryChan)

	fileName := string(td.FileID) + ".txt"

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		retryChan <- td // здесь должен быть re try
		return retryChan

	}

	file.WriteString(fmt.Sprintf("%s\n", td.Payload))
	defer file.Close()
	return retryChan
}

func ParseFileCache(wg *sync.WaitGroup, fc *FileCache) {
	go func() {
		defer wg.Wait()
		for fileId, data := range fc.Cache {
			fc.RemoveNotesBy(fileId)
			wg.Add(1)
			go func(fileId FileID, data []string) {
				defer wg.Done()
				for _, d := range data {
					writeText(TemporaryData{FileID: fileId, Payload: d}) // хздесь можно привернуть интерфейс
				}

			}(fileId, data)

		}

	}()
}

// type Writer interface {
// 	Write(fc FileCache)
// }

// type WriteToFile struct {
// }

// func (wtf *WriteToFile) Write(wg *sync.WaitGroup, mu *sync.RWMutex, fc FileCache) {

// }
