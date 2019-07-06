package main

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info("websocket connected.")

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Error(err)
			return
			//goto ERR
		}
		log.Info("msgType: ", msgType)
		log.Infof("receive message: %s ", msg)

		log.Infof("send message: %s ", msg)
		if err = conn.WriteMessage(msgType, msg); err != nil {
			log.Error(err)
			return
			//goto ERR
		}
	}
}

func main() {
	log.Info("man...")
	http.HandleFunc("/ws", wsHandler)

	http.ListenAndServe(":8080", nil)
}
