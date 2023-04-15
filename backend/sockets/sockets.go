package sockets

import (
	"encoding/json"
	"log"

	"github.com/achintya-7/sigma-chat/models"
	socketio "github.com/googollee/go-socket.io"
)

func Setup() *socketio.Server {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		log.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "join", func(s socketio.Conn, msg string) {
		log.Println("join : ", msg)
		s.Join(msg)
	})

	server.OnEvent("/", "send", func(s socketio.Conn, msg string) {
		var req models.Send
		err := json.Unmarshal([]byte(msg), &req)
		if err != nil {
			log.Println(err)
			return
		}

		ok := server.BroadcastToRoom("/", req.Room, "message", req.Msg); if !ok {
			log.Println("failed to broadcast, err : ", err)
		}
	})

	return server
}
