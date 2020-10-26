package main

import (
	"io/ioutil"
	"swbkaon/srmodel"

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
	data, err := ioutil.ReadFile("../.gitignore")
	if err != nil {
		logrus.Fatal(err)
	}
	err = dial.WriteMessage(websocket.BinaryMessage, srmodel.WriteFile(
		srmodel.File{
		FileName: "conf.json",
	}, data))
	if err != nil {
		logrus.Errorln(err)
	}
}
