package srmodel

import (
	"bytes"
	"encoding/binary"
	"reflect"
)

const (
	FileHeaderChunkSize = 128
	FileNameChunkSize   = 16
)

type RawData []byte
type EncData RawData

type SendReceive struct {
	// Data 文件源数据
	Data []byte
}

type File struct {
	FileName string
	FileData []byte
}

func NewSendReceive(file *File) *SendReceive {
	return &SendReceive{
		Data: make([]byte, len(file.FileName) + len(file.FileData)),
	}
}

func WriteFile(fileHeader File, rawData []byte) EncData {
	encodedData := make(EncData, FileHeaderChunkSize + fileHeader.lengthOfFile() + len(rawData))
	copy(encodedData[:FileHeaderChunkSize], []byte(fileHeader.FileName))
	copy(encodedData[FileHeaderChunkSize:], rawData)
	return encodedData
}

func ReadFile(rdata RawData) *File {
	fileHeader := &File{}
	header := rdata[:FileHeaderChunkSize]
	fileHeader.FileName = readChar(header[:FileNameChunkSize])
	fileHeader.FileData = readRawData(rdata[FileHeaderChunkSize:])
	return fileHeader
}

func (f File)lengthOfFile() int{
	var i = 0
	var totalLength = 0
	for value := reflect.ValueOf(f); i < value.NumField(); i++ {
		totalLength+=value.Field(i).Len()
	}
	return totalLength
}

func BytesToInt(b []byte) int {
	bytesBuf := bytes.NewBuffer(b)
	var x int
	err := binary.Read(bytesBuf, binary.BigEndian, &x)
	if err != nil {
		return -1
	}
	return x
}

func IntToBytes(n int) []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, uint32(n))
	return bs
}

func readChar(b []byte) string {
	x := ""
	for _, v := range b {
		if v != 0 {
			x += string(v)
			continue
		}
	}
	return x
}

func readRawData(b []byte) []byte {
	r := make([]byte, 0)
	for _, v := range b {
		if v != 0 {
			r = append(r, v)
			continue
		}
	}
	return r
}