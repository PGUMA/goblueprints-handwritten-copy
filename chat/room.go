package main

import (
	"log"
	"net/http"

	"github.com/PGUMA/go-project-template/trace"
	"github.com/gorilla/websocket"
)

type room struct {
	foward  chan []byte
	join    chan *client
	leave   chan *client
	clients map[*client]bool
	tracer  trace.Tracer
}

func newRoom() *room {
	return &room{
		foward:  make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer:  trace.Off(),
	}
}

func (r *room) run() {
	for {
		// clientsへの同時アクセスを防ぐ仕組み
		select {
		// 参加
		case client := <-r.join:
			r.clients[client] = true
			r.tracer.Trace("join at new chat.")
			// 退出
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("remove from a chat.")
			// 転送
		case msg := <-r.foward:
			r.tracer.Trace("get message!")
			for client := range r.clients {
				select {
				case client.send <- msg:
					r.tracer.Trace("sended message!")
				default:
					delete(r.clients, client)
					close(client.send)
					r.tracer.Trace("send error!!!!!!")
				}
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize: socketBufferSize, WriteBufferSize: messageBufferSize,
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}

	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
