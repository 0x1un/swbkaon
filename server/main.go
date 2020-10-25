package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logrus.Errorln(err)
			return
		}
		go wsReader(ws)
	})
	logrus.Fatal(http.ListenAndServe(":8888", nil))
}

func wsReader(ws *websocket.Conn) {
	defer func() {
		_ = ws.Close()
	}()
	ws.SetReadLimit(10240)
	if err := ws.SetReadDeadline(time.Now().Add(60 * time.Second)); err != nil {
		logrus.Errorln(err)
	}
	for {
		msgT, msg, err := ws.ReadMessage()
		if err != nil {
			if err, ok := err.(*websocket.CloseError); ok && err.Code == websocket.CloseNormalClosure {
				logrus.Info("normal closed...")
			}
		}
		switch msgT {
		case websocket.BinaryMessage:
			filename := func() string {
				str := ""
				for _, v := range msg[:512] {
					if v == 0 {
						break
					}
					str += string(v)
				}
				return str
			}()
			err := ioutil.WriteFile("../out/"+filename, msg[512:], 0644)
			if err != nil {
				logrus.Errorln(err)
			}
			logrus.Printf("received file: %s", filename)
			return
		case websocket.TextMessage:

		}
	}
}
