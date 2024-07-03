package entity

import (
	"reflect"
	"testing"
)

func TestNewFileCache(t *testing.T) {

	fc1 := &FileCache{}
	fc1.Cache = make(map[FileID][]string)
	fc2 := NewFileCache()

	fc1.AddNewNote(FileID("f1"), "Note")
	fc2.AddNewNote(FileID("f1"), "Note")

	eq := reflect.DeepEqual(fc1, fc2)

	if !eq {
		t.Errorf("actual %v, expected %v", fc2, fc1)
	}

}

func TestAddNewNote(t *testing.T) {
	fc1 := NewFileCache()
	fc2 := NewFileCache()

	fc1.Cache[FileID("123")] = append(fc1.Cache[FileID("123")], "111")
	fc1.Cache[FileID("123")] = append(fc1.Cache[FileID("123")], "222")
	fc1.Cache[FileID("123")] = append(fc1.Cache[FileID("123")], "333")
	fc1.Cache[FileID("123")] = append(fc1.Cache[FileID("123")], "444")
	fc1.Cache[FileID("123")] = append(fc1.Cache[FileID("123")], "555")
	fc1.Cache[FileID("123")] = append(fc1.Cache[FileID("123")], "666")
	fc1.Cache[FileID("123")] = append(fc1.Cache[FileID("123")], "777")

	fc2.AddNewNote(FileID("123"), "111")
	fc2.AddNewNote(FileID("123"), "222")
	fc2.AddNewNote(FileID("123"), "333")
	fc2.AddNewNote(FileID("123"), "444")
	fc2.AddNewNote(FileID("123"), "555")
	fc2.AddNewNote(FileID("123"), "666")
	fc2.AddNewNote(FileID("123"), "777")

	eq := reflect.DeepEqual(fc1, fc2)

	if !eq {
		t.Errorf("actual %v, expected %v", fc2, fc1)
	}
}

func TestRemoveNotesBy(t *testing.T) {

	fc1 := NewFileCache()
	fc2 := NewFileCache()

	fc1.AddNewNote(FileID("123"), "111")
	fc1.AddNewNote(FileID("123"), "222")
	fc1.AddNewNote(FileID("123"), "333")
	fc1.AddNewNote(FileID("123"), "444")
	fc1.AddNewNote(FileID("123"), "555")
	fc1.AddNewNote(FileID("123"), "666")
	fc1.AddNewNote(FileID("123"), "777")

	fc1.AddNewNote(FileID("222"), "111")
	fc1.AddNewNote(FileID("222"), "222")
	fc1.AddNewNote(FileID("222"), "333")
	fc1.AddNewNote(FileID("222"), "444")
	fc1.AddNewNote(FileID("222"), "555")
	fc1.AddNewNote(FileID("222"), "666")
	fc1.AddNewNote(FileID("222"), "777")

	fc2.AddNewNote(FileID("123"), "111")
	fc2.AddNewNote(FileID("123"), "222")
	fc2.AddNewNote(FileID("123"), "333")
	fc2.AddNewNote(FileID("123"), "444")
	fc2.AddNewNote(FileID("123"), "555")
	fc2.AddNewNote(FileID("123"), "666")
	fc2.AddNewNote(FileID("123"), "777")

	fc2.AddNewNote(FileID("222"), "111")
	fc2.AddNewNote(FileID("222"), "222")
	fc2.AddNewNote(FileID("222"), "333")
	fc2.AddNewNote(FileID("222"), "444")
	fc2.AddNewNote(FileID("222"), "555")
	fc2.AddNewNote(FileID("222"), "666")
	fc2.AddNewNote(FileID("222"), "777")

	fc1.RemoveNotesBy(FileID("123"))
	fc2.RemoveNotesBy(FileID("123"))

	eq := reflect.DeepEqual(fc1, fc2)

	if !eq {
		t.Errorf("actual %v, expected %v", fc2, fc1)
	}

}
