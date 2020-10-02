package router

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsEndpoint(msgChan <-chan int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// TODO: Allow origin
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
		}

		// helpful log statement to show connections
		log.Println("Client Connected")
		err = ws.WriteMessage(1, []byte("Hi Client!"))
		if err != nil {
			log.Println(err)
		}

		writer(ws, msgChan)
	}
}

func writer(conn *websocket.Conn, msgChan <-chan int) {
	defer conn.Close()
	for {
		msgInt := <-msgChan
		msgString := strconv.Itoa(msgInt)
		log.Println("writer: ", msgString)

		if err := conn.WriteMessage(websocket.TextMessage, []byte(msgString)); err != nil {
			log.Println("conn.WriteMessage: ", err)
			return
		}

	}
}
