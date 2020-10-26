package srmodel

import (
	"testing"
)

func TestReceiveFile(t *testing.T) {
	fName := "text.txx"
	raw := "hello, world, text txx file..."
	data := WriteFile(File{
		FileName: fName,
		FileData: nil,
	}, []byte(raw))
	file := ReadFile(RawData(data))
	if !(string(file.FileData) == raw && file.FileName == fName) {
		t.Fatal("failed!")
	}
}
