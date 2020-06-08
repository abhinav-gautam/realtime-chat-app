package main

import (
	"fmt"
	"github.com/abhinav-gautam/realtime-chat-app/pkg"
	"log"
	"net/http"
)


func serveWs(pool *websocket.Pool,w http.ResponseWriter,r *http.Request){
	fmt.Println("[+] Websocket endpoint hit")
	conn,err := websocket.Upgrade(w,r)
	if err != nil {
		fmt.Fprintf(w,"%+V\n",err)
		return
	}
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register<-client
	client.Read()
}

func setupRoutes(){
	pool := websocket.NewPool()
	go pool.Start()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool,w,r)
	})
}
func main() {
	fmt.Println("[+] Server started")
	setupRoutes()
	log.Fatalln(http.ListenAndServe(":8080",nil))
}

