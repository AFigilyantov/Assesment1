package entity

import "sync"

type FileCache struct {
	Cache map[FileID][]string
}

func NewFileCache() *FileCache {
	var f FileCache
	f.Cache = make(map[FileID][]string)
	return &f
}

func (c *FileCache) AddNewNote(fileId FileID, note string) {

	c.Cache[fileId] = append(c.Cache[fileId], note)

}

func (fc *FileCache) WriteDataTo(messages <-chan Message, users *Users) {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	for mes := range messages {

		_, ok := users.WhiteList[mes.Token] // проверяем валидность токена если нет сообщение отбрасывеется
		if !ok {
			continue
		}
		wg.Add(1)
		go func(mes Message) {
			defer wg.Done()
			mu.Lock()
			fc.AddNewNote(FileID(mes.FileID), mes.Data)
			mu.Unlock()
		}(mes)

	}
	defer wg.Wait()
}

func (fc *FileCache) RemoveNotesBy(fileId FileID) {
	delete(fc.Cache, fileId)
}
