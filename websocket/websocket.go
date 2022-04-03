package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
	// for{
	reader(ws)
	// }
}

func reader(conn *websocket.Conn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("reader panic ", err)
		}
	}()
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func Test() {
	go func() {
		fmt.Println("enter web sock test")
		fmt.Println("Hello World")
		setupRoutes()
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
}
