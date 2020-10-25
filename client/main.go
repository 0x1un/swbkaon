package main

import (
	"io/ioutil"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

func main() {
	dial, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8888/ws", nil)
	defer func() {
		_ = dial.Close()
	}()
	if err != nil {
		logrus.Errorln(err)
	}
	data, err := ioutil.ReadFile("/home/aumujun/conf.json")
	if err != nil {
		logrus.Fatal(err)
	}
	binMsg := make([]byte, 512+len(data))
	copy(binMsg, []byte("conf.json"))
	copy(binMsg[512:], data)
	err = dial.WriteMessage(websocket.BinaryMessage, binMsg)
	if err != nil {
		logrus.Errorln(err)
	}
}
