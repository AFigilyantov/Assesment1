package entity

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
