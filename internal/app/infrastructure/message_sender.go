package infrastructure

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/service"
	"github.com/gorilla/websocket"
)

const writeWait = 10 * time.Millisecond

type Client struct {
	conn   *websocket.Conn
	cancel chan struct{}
	ch     chan interface{}
	err    chan error
}

func (c *Client) run() {
	for {
		select {
		case <-c.cancel:
			return
		case msg := <-c.ch:
			switch msg := msg.(type) {
			case []byte:
				err := c.conn.WriteMessage(websocket.BinaryMessage, msg)
				if err != nil {
					c.err <- err
					return
				}
			case string:
				err := c.conn.WriteMessage(websocket.TextMessage, []byte(msg))
				fmt.Println(msg)
				if err != nil {
					c.err <- err
					return
				}
			default:
				c.err <- errors.New("unknown message type")
			}
		}
	}
}

type MsgSender struct {
	mutex   *sync.RWMutex
	clients map[string]*Client
}

func NewMsgSender() service.IMessageSender {
	return &MsgSender{
		mutex:   new(sync.RWMutex),
		clients: make(map[string]*Client),
	}
}

// Send implements service.MessageSender.
func (s *MsgSender) Send(ctx context.Context, to string, data interface{}) error {
	s.mutex.RLock()
	client, ok := s.clients[to]
	s.mutex.RUnlock()
	if !ok {
		return errors.New("client not found")
	}

	select {
	case client.ch <- data:
		return nil
	case <-time.After(writeWait):
		return errors.New("websocket write timeout")
	}
}

func (s *MsgSender) Broadcast(ctx context.Context, ids []string, data interface{}) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, id := range ids {
		client, ok := s.clients[id]
		if !ok {
			continue
		}

		select {
		case client.ch <- data:
		case <-time.After(writeWait):
			return errors.New("websocket write timeout")
		}

	}
	return nil
}

func (s *MsgSender) Register(userID string, conn *websocket.Conn, err chan error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	client := &Client{
		conn:   conn,
		cancel: make(chan struct{}),
		ch:     make(chan interface{}, 100),
		err:    err,
	}
	go client.run()

	s.clients[userID] = client
}

func (s *MsgSender) Unregister(userID string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	client, ok := s.clients[userID]
	if !ok {
		return
	}

	close(client.cancel)
	delete(s.clients, userID)
}
