package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/yosef-adel/toll-calculator/types"
)

func main() {

	recv := newDataReceiver()
	http.HandleFunc("/ws", recv.handleWS)
	http.ListenAndServe(":30000", nil)
}

type DataReceiver struct {
	msgch chan types.OBUData
	conn  *websocket.Conn
}

func newDataReceiver() *DataReceiver {
	return &DataReceiver{
		msgch: make(chan types.OBUData, 128),
	}
}

func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	dr.conn = conn

	go dr.wsReceiveLoop()
}

func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("NEW OBU connected  and sending data !")
	for {
		var msg types.OBUData
		err := dr.conn.ReadJSON(&msg)
		if err != nil {
			log.Fatal("error reading data from server", err)
			continue
		}
		fmt.Printf("received data from [%d] :: <lat: %.2f, long: %.2f>\n", msg.OBUID, msg.Lat, msg.Long)
		// dr.msgch <- msg
	}
}
