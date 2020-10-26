package srmodel

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
	Data   []byte
}

func NewSendReceiveWS() *SendReceiveWS {
	return &SendReceiveWS{}
}

func (s *SendReceiveWS) ReceiveFile(file []byte) *FileHeader {

	return nil
}
