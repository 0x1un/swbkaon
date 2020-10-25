package srmodel

import (
	"github.com/gorilla/websocket"
)

// FileHeader 消息头信息
type FileHeader struct {
	// FileNameChunkSize 文件名称块大小
	FileNameChunkSize int
	// FileNameChunkSize 文件权限块大小
	FilePermChunkSize int
}

type SendReceiveWS struct {
	// Header 文件头信息
	Header FileHeader
	// Ws websocket 连接器
	Ws *websocket.Conn
}

func NewSendReceiveWS() *SendReceiveWS {
	return &SendReceiveWS{}
}

func (s *SendReceiveWS) ReceiveFile() (filename string, data []byte, err error) {
	return
}
