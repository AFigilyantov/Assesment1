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

func (fc *FileCache) WriteDataTo(wg *sync.WaitGroup, mu *sync.Mutex, messages <-chan Message, users *Users) {

	for mes := range messages {
		_, ok := users.WhiteList[mes.Token]
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

}

func (fc *FileCache) RemoveNotesBy(fileId FileID) {
	delete(fc.Cache, fileId)
}
