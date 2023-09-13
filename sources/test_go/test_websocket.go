package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"github.com/gorilla/websocket"
// )

// // entrypoint
// var webSocketUpgrader = websocket.Upgrader{
// 	ReadBufferSize: 1024,
// 	WriteBufferSize: 1024,
// }

// func wsEndpoint(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "ws endpoint")

// 	// TODO: change CheckOrigin logic after.
// 	webSocketUpgrader.CheckOrigin = func(r *http.Request) bool {return true}

// 	wsConn, err := webSocketUpgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	log.Println("web socket connected >>>")
// 	err = wsConn.WriteMessage(1, []byte("Hi Client!"))
	
// }

// func setupRoutes() {
// 	http.HandleFunc("/ws", wsEndpoint)
// }

// // service layer

// // repository layer

// func main() {
// 	fmt.Print("start web server >>>")

// 	setupRoutes()
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
