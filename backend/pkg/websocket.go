package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:1024,
	WriteBufferSize:1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
func Upgrade(w http.ResponseWriter,r *http.Request)(*websocket.Conn,error){
	fmt.Println("[+] Host -->",r.Host)
	conn,err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	return conn,nil

}
